package middleware

import (
	"context"
	"io"
	"time"

	"github.com/micro/go-micro/metadata"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// const endpointURL = "localhost:6831"

func NewTracer(servicename, url string) (opentracing.Tracer, io.Closer, error) {
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
	sender, err := jaeger.NewUDPTransport(url, 0) // set Jaeger report revice address
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

func Trace(ctx context.Context, req, res interface{}, err error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	var sp opentracing.Span
	wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	// create new span and bind with context
	sp = opentracing.StartSpan("Get", opentracing.ChildOf(wireContext))
	// record request
	sp.SetTag("request:", req)
	if err != nil {
		sp.SetTag("err", err)
	} else {
		sp.SetTag("response:", res)
	}
	sp.Finish()
}
