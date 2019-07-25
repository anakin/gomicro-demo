package main

import (
	"demo4/middleware"
	pb "demo4/restaurant-service/proto/restaurant"
	"log"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

const ServiceName = "chope.co.srv.restaurant"

func main() {
	//opentracing
	//TODO from config file
	url := "jaeger:6831"
	t, io, err := middleware.NewTracer(ServiceName, url)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	reg := consul.NewRegistry()
	srv := micro.NewService(
		micro.Name(ServiceName),
		micro.RegisterInterval(time.Second*10),
		micro.RegisterTTL(time.Second*30),
		micro.Registry(reg),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(ocplugin.NewClientWrapper(opentracing.GlobalTracer())),
	)
	srv.Init()
	repo := &BookRepository{}
	pb.RegisterBookServiceHandler(srv.Server(), &service{repo: repo})
	srv.Run()
}
