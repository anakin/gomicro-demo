package main

import (
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
)

func main() {
	reg := consul.NewRegistry()
	srv := micro.NewService(
		micro.Name("chope.co.srv.diner"),
		micro.Registry(reg),
	)
	srv.Init()
	srv.Run()
}
