package main

import (
	"context"
	"fmt"

	pb "github.com/anakin/gomicro-demo/demo2/proto/user"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
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
	)
	srv.Init()
	pb.RegisterUserHandler(srv.Server(), new(User))

	err := srv.Run()
	if err != nil {
		fmt.Println(err)
	}
}
