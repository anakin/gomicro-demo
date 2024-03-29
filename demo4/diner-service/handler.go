package main

import (
	"context"
	"demo4/lib/tracer"
	pb "demo4/user-service/proto/user"

	"github.com/sirupsen/logrus"

	"github.com/micro/go-micro/metadata"
)

type Sub struct {
}

//func (s *Sub) Handle(ctx context.Context, msg *pb.Event) error {
//	log.Log("Handler Received message: ", msg.Message)
//	return nil
//}

func (s *Sub) Process(ctx context.Context, ev *pb.Event) (err error) {
	defer tracer.Trace(ctx, "Process", ev, nil, err)
	md, _ := metadata.FromContext(ctx)
	logrus.Infof("[diner] Received event %+v with metadata %+v\n", ev, md)
	return nil
}
