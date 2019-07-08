package main

import (
	"context"
	pb "demo4/user-service/proto/user"
	"fmt"
	"time"

	micro "github.com/micro/go-micro"
)

type service struct {
	repo Repository
	pub  micro.Publisher
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	fmt.Println("receiveid /user/get request", req)
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	ev := &pb.Event{
		Id:        "111",
		Timestamp: time.Now().Unix(),
		Message:   fmt.Sprintf("user message,%s", user),
	}
	fmt.Println(ev)
	 err=srv.pub.Publish(context.Background(), ev)
	if err != nil {
		fmt.Println("pub error",err)
	}
	return nil
}
