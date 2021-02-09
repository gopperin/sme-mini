package main

import (
	"context"
	"encoding/json"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	lift "github.com/liftbridge-io/go-liftbridge/v2"

	"store-user/config"
	"store-user/controller"
	"store-user/persist"
	pb "types/pb"
)

// createStream
func createStream(client lift.Client, subjects map[string]interface{}) error {

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

	_liftClient, err := lift.Connect(config.Lift.Addrs)
	if err != nil {
		fmt.Println("*** lift error : ", err.Error())
		return
	}
	defer _liftClient.Close()

	createStream(_liftClient, config.Lift.Subjects)

	ctx := context.Background()
	if err := _liftClient.Subscribe(ctx, "gudp.user.create", func(msg *lift.Message, err error) {
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("gudp.user.create:", msg.Timestamp(), msg.Offset(), string(msg.Key()), string(msg.Value()))
		var cmd pb.GudpUserCreateCommand
		err = json.Unmarshal(msg.Value(), &cmd)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("CreateUser", cmd)
		controller.CreateUser(cmd)

	}, lift.StartAtEarliestReceived(), lift.Partition(0)); err != nil {
		panic(err)
	}

	<-ctx.Done()

}
