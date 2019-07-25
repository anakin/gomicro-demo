package main

import (
	"demo4/lib/tracer"
	"demo4/middleware"

	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker/nats"
	"github.com/micro/go-micro/registry/consul"
	"github.com/sirupsen/logrus"
)

const ServiceName = "chope.co.srv.diner"
const PubSubServiceName = "chope.co.pubsub.user"

func init() {
	middleware.InitWithFile(".env.json")
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000",
		FullTimestamp:   true,
	})
}

func main() {
	//opentracing
	tracer, closer, err := tracer.NewTracer(ServiceName)
	if err != nil {
		logrus.Error(err)
	}
	defer closer.Close()
	//registry
	reg := consul.NewRegistry()

	//broker
	broker := nats.NewBroker()
	srv := micro.NewService(
		micro.Name(ServiceName),
		micro.Broker(broker),
		micro.Registry(reg),
		micro.WrapSubscriber(ocplugin.NewSubscriberWrapper(tracer)),
	)
	srv.Init()
	_ = micro.RegisterSubscriber(PubSubServiceName, srv.Server(), new(Sub))
	err = srv.Run()
	if err != nil {
		logrus.Fatal(err)
	}
}
