package main

import (
	"context"
	"demo4/lib/tracer"
	"demo4/user-api/handler"
	"net/http"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/sirupsen/logrus"

	hystrixplugin "github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	libtracer "demo4/lib/wrapper/tracer"
	"demo4/middleware"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/web"
)

const ServiceName = "chope.co.api.user"

func init() {
	middleware.InitWithFile(".env.json")
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000",
		FullTimestamp:   true,
	})
	logrus.SetOutput(os.Stderr)
}

func main() {

	libtracer.SetSamplingFrequency(50)
	t, io, err := tracer.NewTracer(ServiceName)
	if err != nil {
		logrus.Fatal(err)
	}
	defer io.Close()

	opentracing.SetGlobalTracer(t)

	//registry
	reg := consul.NewRegistry()
	srv := web.NewService(
		web.Name(ServiceName),
		web.MicroService(grpc.NewService()),
		web.Registry(reg),
		web.RegisterTTL(time.Second*15),
		web.RegisterInterval(time.Second*10),
	)

	_ = srv.Init()

	hystrix.DefaultTimeout = 5000

	sClient := hystrixplugin.NewClientWrapper()(srv.Options().Service.Client())
	_ = sClient.Init(
		client.Retries(3),
		client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (bool, error) {
			logrus.Info(req.Method(), retryCount, " client retry")
			return true, nil
		}),
	)

	h := handler.New(sClient)
	router := gin.Default()
	r := router.Group("/user")
	r.Use(libtracer.GinWrapper)
	r.GET("/info", h.Info)
	r.POST("/create", h.Create)
	r.POST("/auth", h.Auth)
	srv.Handle("/", router)
	err = srv.Run()
	if err != nil {
		logrus.Fatal(err)
	}
	//PrometheusBoot()
}
func PrometheusBoot() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":8085", nil)
		if err != nil {
			logrus.Error(err)
		}
	}()
}
