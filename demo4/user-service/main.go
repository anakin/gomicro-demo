package main

import (
	"demo4/user-service/dbops"
	pb "demo4/user-service/proto/user"
	"log"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker/nats"
	"github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"
	"github.com/micro/go-micro/registry/consul"
	rl "github.com/juju/ratelimit"
)

func main() {
	var consulAddr string
	reg := consul.NewRegistry()

	r:=rl.NewBucketWithRate(1,1)
	broker := nats.NewBroker()
	repo := &UserRepository{}
	srv := micro.NewService(
		micro.Registry(reg),
		micro.Broker(broker),
		micro.Name("chope.co.srv.user"),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(r,false)),
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
	pub := micro.NewPublisher("chope.co.pubsub.user", srv.Client())
	pb.RegisterUserServiceHandler(srv.Server(), &service{repo: repo, pub: pub})
	err := srv.Run()

	if err != nil {

		log.Fatal("run user service error", err)

	}
}
