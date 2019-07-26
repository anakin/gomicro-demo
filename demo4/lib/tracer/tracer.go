package tracer

import (
	"context"
	"demo4/lib/config"
	"fmt"
	"io"
	"time"

	"github.com/micro/go-micro/metadata"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// NewTracer 创建一个jaeger Tracer
func NewTracer(servicename string) (opentracing.Tracer, io.Closer, error) {
	cfg := jaegercfg.Configuration{
		ServiceName: servicename,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	addr := fmt.Sprintf("%s:%d", config.G_cfg.Jaeger.Host, config.G_cfg.Jaeger.Port)
	sender, err := jaeger.NewUDPTransport(addr, 0)
	if err != nil {
		return nil, nil, err
	}

	reporter := jaeger.NewRemoteReporter(sender)
	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Reporter(reporter),
	)

	return tracer, closer, err
}

func Trace(ctx context.Context, method string, req, res interface{}, err error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	var sp opentracing.Span
	wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	// create new span and bind with context
	sp = opentracing.StartSpan(method, opentracing.ChildOf(wireContext))
	// record request
	sp.SetTag("request:", req)
	if err != nil {
		sp.SetTag("err", err)
	} else {
		sp.SetTag("response:", res)
	}
	sp.Finish()
}
