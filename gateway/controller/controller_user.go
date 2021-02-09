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

	_obj, err := persist.GMariadb.GetUserByUID(uid)
	if err != nil {
		return _obj, errors.New("获取用户信息错误")
	}

	return _obj, nil
}

// CreateUserRPC calls the CreateEvent RPC
func CreateUserRPC(cmd pb.GudpUserCreateCommand) error {

	_client := pb.NewEventStoreClient(mygrpc.ClientConn)
	_jsonStr, _ := json.Marshal(cmd)

	event := &pb.Event{
		EventId:       cmd.Uid,
		EventType:     myconstant.EventUserCreate,
		AggregateId:   cmd.Uid,
		AggregateType: myconstant.AggregateUser,
		EventData:     string(_jsonStr),
		Channel:       myconstant.AggregateUser,
		Stream:        myconstant.EventUserCreate,
	}

	_resp, err := _client.CreateEvent(context.Background(), event)
	if err != nil {
		return errors.Wrap(err, "Error from RPC server")
	}
	if _resp.IsSuccess {
		return nil
	}
	return errors.Wrap(err, "Error from RPC server")
}
