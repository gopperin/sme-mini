package controller

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	mygrpc "gateway/grpc"
	"gateway/persist"
	"gateway/snowflake"
	myconstant "types/constant"
	mystore "types/mariadb"
	"types/pb"
)

// CreateUser CreateUser
func CreateUser(cmd pb.GudpUserCreateCommand) (int64, error) {

	cmd.Uid = snowflake.Olaf.ID64()

	err := CreateUserRPC(cmd)
	if err != nil {
		return -1, err
	}
	return cmd.Uid, nil
}

// GetUserByUID GetUserByUID
func GetUserByUID(uid int64) (mystore.GudpUserBase, error) {

	obj, err := persist.GMariadb.GetUserByUID(uid)
	if err != nil {
		return obj, errors.New("获取用户信息错误")
	}

	return obj, nil
}

// CreateUserRPC calls the CreateEvent RPC
func CreateUserRPC(cmd pb.GudpUserCreateCommand) error {

	client := pb.NewEventStoreClient(mygrpc.ClientConn)
	jsonStr, _ := json.Marshal(cmd)

	event := &pb.Event{
		EventId:       cmd.Uid,
		EventType:     myconstant.EventUserCreate,
		AggregateId:   cmd.Uid,
		AggregateType: myconstant.AggregateUser,
		EventData:     string(jsonStr),
		Channel:       myconstant.AggregateUser,
		Stream:        myconstant.EventUserCreate,
	}

	resp, err := client.CreateEvent(context.Background(), event)
	if err != nil {
		return errors.Wrap(err, "Error from RPC server")
	}
	if resp.IsSuccess {
		return nil
	}
	return errors.Wrap(err, "Error from RPC server")
}
