package main

import (
	"context"
	pb "demo4/user-service/proto/user"
	"fmt"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/util/log"
)

type Sub struct {
}
//func (s *Sub) Handle(ctx context.Context, msg *pb.Event) error {
//	log.Log("Handler Received message: ", msg.Message)
//	return nil
//}

func (s *Sub) Process(ctx context.Context, ev *pb.Event) error{
	md, _ := metadata.FromContext(ctx)
	fmt.Println("[diner] Received event  with metadata ", ev, md)
	log.Logf("[diner] Received event %+v with metadata %+v\n", ev, md)
	return nil
}
