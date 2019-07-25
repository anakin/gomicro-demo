package main

import (
	"demo4/lib/tracer"
	"demo4/middleware"
	pb "demo4/restaurant-service/proto/restaurant"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
)

const ServiceName = "chope.co.srv.restaurant"

func init() {
	middleware.InitWithFile(".env.json")
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000",
		FullTimestamp:   true,
	})
}

func main() {
	//opentracing
	t, io, err := tracer.NewTracer(ServiceName)
	if err != nil {
		logrus.Error(err)
	}
	defer io.Close()

	reg := consul.NewRegistry()
	srv := micro.NewService(
		micro.Name(ServiceName),
		micro.RegisterInterval(time.Second*10),
		micro.RegisterTTL(time.Second*30),
		micro.Registry(reg),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(t), middleware.LogHandlerWrapper),
		//micro.WrapClient(ocplugin.NewClientWrapper(t), middleware.LogClientWrapper),
		micro.WrapClient(middleware.LogClientWrapper),
	)
	srv.Init()
	repo := &BookRepository{}
	_ = pb.RegisterBookServiceHandler(srv.Server(), &service{repo: repo})
	err = srv.Run()
	if err != nil {
		logrus.Fatal(err)
	}
}
