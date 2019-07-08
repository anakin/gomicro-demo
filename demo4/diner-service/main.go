package main

import (
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/broker/nats"
	"github.com/micro/go-micro/registry/consul"
)

func main() {
	reg := consul.NewRegistry()
	broker := nats.NewBroker()
	srv := micro.NewService(
		micro.Name("chope.co.srv.diner"),
		micro.Broker(broker),
		micro.Registry(reg),
	)
	srv.Init()
	micro.RegisterSubscriber("chope.co.pubsub.user", srv.Server(),new(Sub))
	srv.Run()
}
