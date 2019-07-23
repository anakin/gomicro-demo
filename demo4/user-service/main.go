package main

import (
	"demo4/middleware"
	"demo4/tracer"
	"demo4/user-service/config"
	pb "demo4/user-service/proto/user"
	"fmt"
	"log"

	"github.com/opentracing/opentracing-go"

	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"

	rl "github.com/juju/ratelimit"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker/nats"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"
)

const ServiceName = "chope.co.srv.user"
const BrokerServiceName = "chope.co.pubsub.user"

func main() {
	var consulAddr string

	//from file
	config.InitWithFile(".env.json")

	url := fmt.Sprintf("%s:%d", config.G_cfg.Jaeger.Host, config.G_cfg.Jaeger.Port)
	//服务发现使用consul
	reg := consul.NewRegistry()

	//限流
	r := rl.NewBucketWithRate(1, 1)

	//opentracing
	t, io, err := tracer.NewTracer(ServiceName, url)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//异步的pub/sub使用nats
	broker := nats.NewBroker()
	breaker := middleware.NewHytrixWrapper()
	srv := micro.NewService(
		micro.Registry(reg),
		micro.Broker(broker),
		micro.Name(ServiceName),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(r, false), middleware.LogHandlerWrapper, ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(ocplugin.NewClientWrapper(opentracing.GlobalTracer())),
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

	//from consul
	//config.InitWithConsul(consulAddr)

	//注册broker
	pub := micro.NewPublisher(BrokerServiceName, srv.Client())

	repo := &UserRepository{}

	//注册服务
	handler := NewService(srv.Client(), repo, pub)
	_ = pb.RegisterUserServiceHandler(srv.Server(), handler)
	err = srv.Run()

	if err != nil {
		log.Fatal("run user service error", err)
	}
}
