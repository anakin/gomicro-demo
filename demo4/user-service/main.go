package main

import (
	"log"

	"demo4/user-service/dbops"
	"github.com/micro/go-micro/broker/nats"
	pb "demo4/user-service/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
)

func main() {
	var consulAddr string
	reg := consul.NewRegistry()

	broker:=nats.NewBroker()
	repo := &UserRepository{}
	srv := micro.NewService(
		micro.Registry(reg),
		micro.Broker(broker),
		micro.Name("chope.co.srv.user"),
		micro.Flags(cli.StringFlag{
			Name:   "consul_address",
			Usage:  "consul address for K/V",
			EnvVar: "CONSUL_ADDRESS",
			Value:  "127.0.0.1:8500",
		}),

		micro.Action(func(ctx *cli.Context) {
			consulAddr = ctx.String("consul_address")
		}),
	)
	srv.Init()
	// fmt.Println("consul addres:", consulAddr)
	dbops.Init(consulAddr)
	pub:=micro.NewPublisher("chope.co.pubsub.user",srv.Client())
	pb.RegisterUserServiceHandler(srv.Server(), &service{repo: repo,pub:pub})
	err := srv.Run()
	if err != nil {
		log.Fatal("run user service error",err)
	}
}
