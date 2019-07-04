package main

import (
	"context"
	"log"
	"os"

	pb "github.com/anakin/gomicro-demo/demo3/user-service/proto/user"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
)

func main() {
	reg := consul.NewRegistry()
	srv := micro.NewService(
		micro.Name("shippy.service-user.cli"),
		micro.Registry(reg),
	)
	srv.Init()
	client := pb.NewUserService("shippy.service.user", srv.Client())
	name := "anakin sun"
	email := "anakinsun@gmail.com"
	password := "test123"
	company := "chope"
	log.Println(name, email, password)
	r, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("could not create :%v", err)
	}
	log.Printf("created:%v", r.User.Email)
	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("could not list users:%v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}
	authRes, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("auth failed:%v", err)
	}
	log.Printf("token is :%v", authRes.Token)
	os.Exit(0)
}
