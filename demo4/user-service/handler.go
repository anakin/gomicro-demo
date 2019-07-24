package main

import (
	"context"
	rest "demo4/restaurant-service/proto/restaurant"
	"demo4/tracer"
	pb "demo4/user-service/proto/user"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

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

func (srv *service) Create(ctx context.Context, in *pb.User, res *pb.Response) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	in.Password = string(hashedPass)
	r, err := srv.repo.Create(in)
	if err != nil {
		return err
	}
	res.User = r
	tracer.Trace(ctx, in, res, err)
	return nil
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	//dbops.Init()

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
	tracer.Trace(ctx, req, res, err)
	//发布broker消息
	go srv.pub.Publish(ctx, ev)
	return nil
}
