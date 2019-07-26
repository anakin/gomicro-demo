package main

import (
	"demo4/lib/tracer"
	"demo4/lib/wrapper/breaker/hystrix"
	"demo4/lib/wrapper/metrics/prometheus"
	"demo4/lib/wrapper/tracer/opentracing/stdhttp"
	"demo4/middleware"

	"github.com/micro/go-micro"

	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func init() {

	_ = plugin.Register(cors.NewPlugin())

	_ = plugin.Register(plugin.NewPlugin(
		plugin.WithName("tracer"),
		plugin.WithHandler(
			stdhttp.TracerWrapper,
		),
	))
	_ = plugin.Register(plugin.NewPlugin(
		plugin.WithName("breaker"),
		plugin.WithHandler(
			hystrix.BreakerWrapper,
		),
	))
	_ = plugin.Register(plugin.NewPlugin(
		plugin.WithName("metrics"),
		plugin.WithHandler(
			prometheus.MetricsWrapper,
		),
	))
	middleware.InitWithFile(".env.json")
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000",
		FullTimestamp:   true,
	})
}

const name = "API gateway"

func main() {
	stdhttp.SetSamplingFrequency(50)
	t, io, err := tracer.NewTracer(name)
	if err != nil {
		logrus.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//hystrixStreamHandler := hx.NewStreamHandler()
	//hystrixStreamHandler.Start()
	//go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)

	cmd.Init(
		micro.Name("chope.co.api"),
	)
}
