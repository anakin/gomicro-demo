package main

import (
	"context"
	rest "demo4/restaurant-service/proto/restaurant"
	"demo4/user-service/dbops"
	pb "demo4/user-service/proto/user"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/micro/go-micro/metadata"

	micro "github.com/micro/go-micro"
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
	//dbops.Migrate()
	fmt.Println("receiveid /user/get request", req)
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	var sp opentracing.Span
	wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	// create new span and bind with context
	sp = opentracing.StartSpan("Get", opentracing.ChildOf(wireContext))
	// record request
	sp.SetTag("req", req)
	defer func() {
		sp.SetTag("res", res)
		sp.Finish()
	}()
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

	rs, err := srv.restS.Book(context.Background(), &rest.Request{Id: "2345"})
	if err != nil {
		fmt.Println("req restaurant error,", err)
		return err
	}
	fmt.Println("got rest resp:", rs)

	//发布broker消息
	go srv.pub.Publish(context.Background(), ev)
	return nil
}
