package main

import (
	"log"

	"demo3/middleware"
	pb "demo3/user-service/proto/user"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

func main() {
	reg := consul.NewRegistry()
	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("could not connect to DB:%v", err)
	}
	//db.AutoMigrate(&pb.User{})
	repo := &UserRepository{db}
	tokenService := &TokenService{repo}
	t, _, err := middleware.NewTracer("user-service")
	if err != nil {
		log.Fatal("tracer error", err)
	}
	opentracing.InitGlobalTracer(t)
	srv := micro.NewService(
		micro.Name("shippy.service.user"),
		micro.Registry(reg),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())), // add tracing plugin in to middleware
	)
	srv.Init()
	_ = pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
