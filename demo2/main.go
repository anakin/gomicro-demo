package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb "demo2/proto/user"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus"
)

type User struct {
}

func (u *User) Get(ctx context.Context, req *pb.Request, res *pb.Reponse) error {
	fmt.Println("received user/get request")
	user := &pb.Uinfo{
		Id:   "1",
		Name: "test",
	}
	res.User = user
	return nil
}

func main() {
	reg := consul.NewRegistry()
	srv := micro.NewService(
		micro.Name("anakin.sun.api.user"),
		micro.Registry(reg),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)
	srv.Init()
	_ = pb.RegisterUserHandler(srv.Server(), new(User))

	err := srv.Run()
	if err != nil {
		fmt.Println(err)
	}
}
func PrometheusBoot() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":8085", nil)
		if err != nil {
			log.Println(err)
		}
	}()
}
