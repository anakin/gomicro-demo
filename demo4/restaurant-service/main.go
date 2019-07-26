package main

import (
	"demo4/lib/config"
	"demo4/lib/logger"
	"demo4/lib/tracer"
	pb "demo4/restaurant-service/proto/restaurant"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/sirupsen/logrus"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
)

const ServiceName = "chope.co.srv.restaurant"

func init() {
	config.InitWithFile(".env.json")
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000",
		FullTimestamp:   true,
	})
	logrus.SetOutput(os.Stderr)
}

func main() {
	//opentracing
	t, io, err := tracer.NewTracer(ServiceName)
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
		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer()), logger.LogHandlerWrapper),
		//micro.WrapClient(ocplugin.NewClientWrapper(t), middleware.LogClientWrapper),
	)
	srv.Init()
	repo := &BookRepository{}
	_ = pb.RegisterBookServiceHandler(srv.Server(), &service{repo: repo})
	err = srv.Run()
	if err != nil {
		logrus.Fatal(err)
	}
}
