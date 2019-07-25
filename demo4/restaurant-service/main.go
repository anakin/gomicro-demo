package main

import (
	"demo4/middleware"
	pb "demo4/restaurant-service/proto/restaurant"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

const ServiceName = "chope.co.srv.restaurant"

func main() {
	//opentracing

	t, io, err := middleware.NewTracer(ServiceName)
	if err != nil {
		logrus.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	reg := consul.NewRegistry()
	srv := micro.NewService(
		micro.Name(ServiceName),
		micro.RegisterInterval(time.Second*10),
		micro.RegisterTTL(time.Second*30),
		micro.Registry(reg),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(t), middleware.LogHandlerWrapper),
		micro.WrapClient(ocplugin.NewClientWrapper(t), middleware.LogClientWrapper),
	)
	srv.Init()
	repo := &BookRepository{}
	_ = pb.RegisterBookServiceHandler(srv.Server(), &service{repo: repo})
	err = srv.Run()
	if err != nil {
		logrus.Fatal(err)
	}
}
