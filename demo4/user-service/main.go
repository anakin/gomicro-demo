package main

import (
	"demo4/middleware"
	"demo4/user-service/config"
	"demo4/user-service/dbops"
	pb "demo4/user-service/proto/user"
	"log"

	rl "github.com/juju/ratelimit"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker/nats"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"
)

func main() {
	var consulAddr string

	//服务发现使用consul
	reg := consul.NewRegistry()

	//限流
	r := rl.NewBucketWithRate(1, 1)

	//异步的pub/sub使用nats
	broker := nats.NewBroker()
	breaker := middleware.NewHytrixWrapper()
	srv := micro.NewService(
		micro.Registry(reg),
		micro.Broker(broker),
		micro.Name("chope.co.srv.user"),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(r, false), middleware.LogHandlerWrapper),
		micro.Flags(cli.StringFlag{
			Name:   "consul_address",
			Usage:  "consul address for K/V",
			EnvVar: "CONSUL_ADDRESS",
			Value:  "127.0.0.1:8500",
		}),
		micro.WrapClient(breaker, middleware.LogClientWrapper),
		//从命令行获取consul服务的地址
		micro.Action(func(ctx *cli.Context) {
			consulAddr = ctx.String("consul_address")
		}),
	)
	srv.Init()

	//from file
	config.InitWithFile(".env.json")

	//from consul
	//config.InitWithConsul(consulAddr)

	//初始化数据库
	dbops.Init()

	//注册broker
	pub := micro.NewPublisher("chope.co.pubsub.user", srv.Client())

	repo := &UserRepository{}

	//注册服务
	handler := NewService(srv.Client(), repo, pub)
	pb.RegisterUserServiceHandler(srv.Server(), handler)
	err := srv.Run()

	if err != nil {
		log.Fatal("run user service error", err)
	}
}
