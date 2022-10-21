package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	lift "github.com/liftbridge-io/go-liftbridge/v2"
	"google.golang.org/grpc"

	"github.com/gopperin/sme-mini/eventstore/config"
	"github.com/gopperin/sme-mini/eventstore/persist"
	"github.com/gopperin/sme-mini/types/proto"
)

type server struct {
	LiftClient lift.Client
}

// CreateEvent RPC creates a new Event into EventStore
func (s *server) CreateEvent(ctx context.Context, in *proto.Event) (*proto.Response, error) {
	// Persist events as immutable logs into CockroachDB
	err := persist.GMariadb.CreateEvent(*in)
	if err != nil {
		return nil, err
	}
	// Publish event on NATS Streaming Server
	go publishEvent(s.LiftClient, in)
	return &proto.Response{IsSuccess: true}, nil
}

// publishEvent publishes an event via NATS Streaming server
func publishEvent(client lift.Client, event *proto.Event) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := client.Publish(ctx, event.Stream,
		[]byte(event.EventData),
		lift.Key([]byte(strconv.FormatInt(event.EventId, 10))),
		lift.ToPartition(config.Lift.Partition),
		lift.AckPolicyAll(),
	); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("pub", event.Stream, event.AggregateId)
}

// preCreateStream
func preCreateStream(client lift.Client, subjects map[string]interface{}) error {

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
	fmt.Println("====== mariadb init")
	defer persist.Close()

	liftClient, err := lift.Connect(config.Lift.Addrs)
	if err != nil {
		fmt.Println("*** lift error : ", err.Error())
		return
	}
	defer liftClient.Close()

	preCreateStream(liftClient, config.Lift.Subjects)

	lis, err := net.Listen("tcp", config.Server.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("====== grpc listen:", config.Server.GrpcPort)

	// Creates a new gRPC server
	s := grpc.NewServer()
	proto.RegisterEventStoreServer(s, &server{liftClient})
	s.Serve(lis)
}
