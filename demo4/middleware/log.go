package middleware

import (
	"context"
	"os"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	log "github.com/sirupsen/logrus"
)

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	begin := time.Now()
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Error("failed to open file")
	}
	logMsg := log.WithFields(log.Fields{
		"ctx":     md,
		"service": req.Service(),
		"method":  req.Method(),
	})
	logMsg.Info("calling service")
	err = l.Client.Call(ctx, req, rsp)
	if err != nil {
		logMsg = logMsg.WithFields(log.Fields{
			"error": err,
		})
	}
	logMsg.WithFields(log.Fields{
		"duration:": int64(float64(time.Since(begin))/float64(time.Microsecond) + 0.5),
	}).Info("called service")
	return err
}
func LogClientWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}

func LogHandlerWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		md, _ := metadata.FromContext(ctx)
		log.WithFields(log.Fields{
			"ctx":    md,
			"method": req.Method(),
		}).Info("serving request")
		err := fn(ctx, req, rsp)
		return err
	}
}
