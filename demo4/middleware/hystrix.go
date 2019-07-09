package middleware

import (
	"context"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
)

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		return c.Client.Call(ctx, req, rsp, opts...)
	}, nil)
}

func NewHytrixWrapper() client.Wrapper {
	hystrix.DefaultTimeout = 120
	// how long to open circuit breaker again
	hystrix.DefaultSleepWindow = 200
	// percent of bad response
	hystrix.DefaultErrorPercentThreshold = 10
	// how much request can be accessed in persecond
	hystrix.DefaultMaxConcurrent = 2

	hystrix.DefaultVolumeThreshold = 1

	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}
