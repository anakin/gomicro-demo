package main

import (
	"demo4/api/handler"
	"fmt"
	"log"

	"demo4/api/config"
	"demo4/tracer"

	"github.com/micro/go-micro"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
	"github.com/opentracing/opentracing-go"
)

const ServiceName = "chope.co.api.user"

func main() {

	//init config
	config.InitWithFile(".env.json")
	url := fmt.Sprintf("%s:%d", config.G_cfg.Jaeger.Host, config.G_cfg.Jaeger.Port)
	//opentracing
	t, io, err := tracer.NewTracer(ServiceName, url)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//registry
	reg := consul.NewRegistry()
	srv := web.NewService(
		web.Name(ServiceName),
		web.Registry(reg),
	)

	_ = srv.Init()
	service := micro.NewService(
		micro.WrapClient(ocplugin.NewClientWrapper(t)),
	)
	h := handler.New(service.Client())
	router := gin.Default()
	r := router.Group("/user")
	r.GET("/info", h.Info)
	srv.Handle("/", router)
	err = srv.Run()
	if err != nil {
		fmt.Println(err)
	}
}
