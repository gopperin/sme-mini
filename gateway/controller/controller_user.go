package controller

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	mygrpc "github.com/gopperin/sme-mini/gateway/grpc"
	"github.com/gopperin/sme-mini/gateway/persist"
	"github.com/gopperin/sme-mini/gateway/snowflake"
	myconstant "github.com/gopperin/sme-mini/types/constant"
	mystore "github.com/gopperin/sme-mini/types/mariadb"
	"github.com/gopperin/sme-mini/types/proto"
)

// CreateUser CreateUser
func CreateUser(cmd mystore.GudpUserBase) (int64, error) {

	cmd.UID = snowflake.Olaf.ID64()

	err := CreateUserRPC(cmd)
	if err != nil {
		return -1, err
	}
	return cmd.UID, nil
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
func CreateUserRPC(cmd mystore.GudpUserBase) error {

	client := proto.NewEventStoreClient(mygrpc.ClientConn)
	jsonStr, _ := json.Marshal(cmd)

	event := &proto.Event{
		EventId:       cmd.UID,
		EventType:     myconstant.EventUserCreate,
		AggregateId:   cmd.UID,
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
