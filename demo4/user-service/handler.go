package main

import (
	"context"
	rest "demo4/restaurant-service/proto/restaurant"
	"demo4/tracer"
	"demo4/user-service/dbops"
	pb "demo4/user-service/proto/user"
	"fmt"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

type service struct {
	repo  Repository
	restS rest.BookService
	pub   micro.Publisher
}

//NewService factory
func NewService(client client.Client, repo Repository, pub micro.Publisher) *service {
	return &service{
		repo:  repo,
		restS: rest.NewBookService("chope.co.srv.restaurant", client),
		pub:   pub,
	}
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	dbops.Init()

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

	rs, err := srv.restS.Book(ctx, &rest.Request{Id: "2345"})
	if err != nil {
		fmt.Println("req restaurant error,", err)
		return err
	}
	fmt.Println("got rest resp:", rs)

	tracer.Trace(ctx, req, res)
	//发布broker消息
	go srv.pub.Publish(context.Background(), ev)
	return nil
}
