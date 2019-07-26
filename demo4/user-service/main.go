package main

import (
	"demo4/lib/tracer"
	"demo4/middleware"
	pb "demo4/user-service/proto/user"

	"github.com/opentracing/opentracing-go"

	"github.com/sirupsen/logrus"

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

func init() {
	middleware.InitWithFile(".env.json")
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000",
		FullTimestamp:   true,
	})
}

func main() {
	var consulAddr string

	//config from file

	//opentracing
	t, c, err := tracer.NewTracer(ServiceName)
	if err != nil {
		logrus.Error(err)
	}
	defer c.Close()
	opentracing.SetGlobalTracer(t)

	//限流
	r := rl.NewBucketWithRate(1000, 1000)

	//registry
	reg := consul.NewRegistry()

	//异步的pub/sub使用nats
	broker := nats.NewBroker()

	//breaker
	breaker := middleware.NewHytrixWrapper()
	srv := micro.NewService(
		micro.Registry(reg),
		micro.Broker(broker),
		micro.Name(ServiceName),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(r, false), middleware.LogHandlerWrapper, ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
		//micro.WrapClient(ocplugin.NewClientWrapper(tracer), middleware.LogClientWrapper),
		micro.WrapClient(middleware.LogClientWrapper),
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
		logrus.Fatal("run user service error", err)
	}
}
