package main

import (
	pb "demo4/restaurant-service/proto/restaurant"
"time"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
)

func main() {
	reg := consul.NewRegistry()
	srv := micro.NewService(
		micro.Name("chope.co.srv.restaurant"),
		micro.RegisterInterval(time.Second*10),
		micro.RegisterTTL(time.Second*30),
		micro.Registry(reg),
	)
	srv.Init()
	repo := &BookRepository{}
	pb.RegisterBookServiceHandler(srv.Server(), &service{repo: repo})
	srv.Run()
}
