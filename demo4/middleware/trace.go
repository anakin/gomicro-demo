package middleware

import (
	"demo4/user-service/config"
	"fmt"
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// const endpointURL = "localhost:6831"

func NewTracer(servicename string) (opentracing.Tracer, io.Closer, error) {
	cfg := config.G_cfg
	endpointURL := fmt.Sprintf("%s:%d", cfg.Jaeger.Host, cfg.Jaeger.Port)
	jCfg := jaegercfg.Configuration{
		ServiceName: servicename, // tracer name
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	sender, err := jaeger.NewUDPTransport(endpointURL, 0) // set Jaeger report revice address
	if err != nil {
		return nil, nil, err
	}
	reporter := jaeger.NewRemoteReporter(sender) // create Jaeger reporter
	// Initialize Opentracing tracer with Jaeger Reporter
	tracer, closer, err := jCfg.NewTracer(
		jaegercfg.Reporter(reporter),
	)
	return tracer, closer, err
}
