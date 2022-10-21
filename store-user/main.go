package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	lift "github.com/liftbridge-io/go-liftbridge/v2"

	"github.com/gopperin/sme-mini/store-user/bolt"
	"github.com/gopperin/sme-mini/store-user/config"
	"github.com/gopperin/sme-mini/store-user/controller"
	"github.com/gopperin/sme-mini/store-user/persist"
	mystore "github.com/gopperin/sme-mini/types/mariadb"
	pb "github.com/gopperin/sme-mini/types/proto"
)

// EventUserCreate EventUserCreate
const EventUserCreate = "event.user.create"

// preCreateStream 预先创建stream
func preCreateStream(client lift.Client, subjects map[string]interface{}) map[string]interface{} {

	// 遍历配置的subject
	for _subject, _streams := range subjects {

		// 遍历每个subject下的stream
		for _, _stream := range _streams.([]interface{}) {
			err := client.CreateStream(
				context.Background(), _subject, _stream.(string),
				lift.MaxReplication(),
				lift.Partitions(12),
			)

			if err == nil {
				fmt.Println("created stream", _stream.(string))
				continue
			}

			if err != nil && err != lift.ErrStreamExists {
				fmt.Println("created stream", err.Error())
				continue
			}

			fmt.Println("stream exist", _stream.(string))
		}

	}

	return nil
}

func main() {

	err := persist.GMariadb.Init()
	if err != nil {
		fmt.Println("*** mariadb error : ", err.Error())
		return
	}
	fmt.Println("====== mariadb init ======")

	liftClient, err := lift.Connect(config.Lift.Addrs)
	if err != nil {
		fmt.Println("*** lift error : ", err.Error())
		return
	}
	defer liftClient.Close()

	preCreateStream(liftClient, config.Lift.Subjects)

	db, _ := bolt.OpenDatabase("my.db")
	defer db.Close()

	ctx := context.Background()

	var offset int64
	strOffset := db.GetData("lift_offset", EventUserCreate)
	if len(strOffset) > 0 {
		offset, _ = strconv.ParseInt(strOffset, 10, 64)
	}
	fmt.Println("start event:", EventUserCreate, "with offset:", offset)

	if err := liftClient.Subscribe(ctx, EventUserCreate, func(msg *lift.Message, err error) {
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(EventUserCreate, msg.Timestamp(), msg.Offset(), string(msg.Key()), string(msg.Value()))

		// 保存offset
		db.PutData("lift_offset", EventUserCreate, strconv.FormatInt(msg.Offset()+1, 10))

		var cmd pb.GudpUserCreateCommand
		err = json.Unmarshal(msg.Value(), &cmd)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var obj mystore.GudpUser
		obj.UID = cmd.Uid
		obj.NickName = cmd.NickName
		obj.Mobile = cmd.Mobile
		obj.Email = cmd.Email
		obj.Pwd = cmd.Pwd
		obj.SecretKey = cmd.SecretKey

		controller.CreateUser(obj)

	}, lift.StartAtOffset(offset), lift.Partition(config.Lift.Partition)); err != nil {
		panic(err)
	}

	<-ctx.Done()

}
