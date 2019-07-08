package main

import (
	"context"
	pb "demo4/user-service/proto/user"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/util/log"
)

type Sub struct {
}

func (s *Sub) Process(ctx context.Context, ev *pb.Event) {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[diner] Received event %+v with metadata %+v\n", ev, md)
}
