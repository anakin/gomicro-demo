package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/micro/cmd"
)

func main() {
	reg := consul.NewRegistry()
	cmd.Init(
		micro.Name("chope.co.api.gateway"),
		micro.Registry(reg),
	)
}
