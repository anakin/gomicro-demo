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
	token TokenService
}

//NewService factory
func NewService(client client.Client, repo Repository, pub micro.Publisher) *service {
	return &service{
		repo:  repo,
		restS: rest.NewBookService("chope.co.srv.restaurant", client),
		pub:   pub,
		token: TokenService{},
	}
}

func (srv *service) Create(ctx context.Context, in *pb.User, res *pb.Response) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		tracer.Trace(ctx, req, res, err)
		return err
	}
	in.Password = string(hashedPass)
	r, err := srv.repo.Create(in)
	if err != nil {
		tracer.Trace(ctx, req, res, err)
		return err
	}
	res.User = r
	tracer.Trace(ctx, in, res, err)
	return nil
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		tracer.Trace(ctx, req, res, err)
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
		tracer.Trace(ctx, req, res, err)
		return err
	}
	fmt.Println("got rest resp:", rs)
	tracer.Trace(ctx, req, res, err)
	//发布broker消息
	go srv.pub.Publish(ctx, ev)
	return nil
}

func (srv *service) Auth(ctx context.Context, in *pb.User, out *pb.Token) error {
	user, err := srv.repo.GetByEmail(in.Email)
	if err != nil {
		tracer.Trace(ctx, in, out, err)
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password))
	if err != nil {
		tracer.Trace(ctx, in, out, err)
		return err
	}
	token, err := srv.token.Encode(user)
	if err != nil {
		tracer.Trace(ctx, in, out, err)
		return err
	}
	out.Token = token
	tracer.Trace(ctx, in, out, err)
	return nil
}
