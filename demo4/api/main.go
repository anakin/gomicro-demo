package main

import (
	"demo4/api/handler"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/micro/go-plugins/wrapper/monitoring/prometheus"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/micro/go-micro"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"

	"demo4/middleware"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
)

const ServiceName = "chope.co.api.user"

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000",
		FullTimestamp:   true,
	})
}
func main() {
	middleware.InitWithFile(".env.json")
	//opentracing
	tracer, closer, err := middleware.NewTracer(ServiceName)
	if err != nil {
		logrus.Error(err)
	}
	defer closer.Close()

	//registry
	reg := consul.NewRegistry()
	srv := web.NewService(
		web.Name(ServiceName),
		web.Registry(reg),
	)

	_ = srv.Init()
	service := micro.NewService(
		micro.WrapClient(ocplugin.NewClientWrapper(tracer), middleware.LogClientWrapper),
		micro.WrapHandler(prometheus.NewHandlerWrapper(), middleware.LogHandlerWrapper, ocplugin.NewHandlerWrapper(tracer)),
	)
	h := handler.New(service.Client())
	router := gin.Default()
	r := router.Group("/user")
	r.GET("/info", h.Info)
	r.POST("/create", h.Create)
	r.POST("/auth", h.Auth)
	srv.Handle("/", router)
	err = srv.Run()
	if err != nil {
		logrus.Fatal(err)
	}
	PrometheusBoot()
}
func PrometheusBoot() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":8085", nil)
		if err != nil {
			logrus.Println(err)
		}
	}()
}
