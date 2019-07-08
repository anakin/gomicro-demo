package main

import (
	"demo4/api/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
)

func main() {
	reg := consul.NewRegistry()
	srv := web.NewService(
		web.Name("chope.co.api.user"),
		web.Registry(reg),
	)
	srv.Init()
	h := handler.New(srv.Options().Service.Client())
	router := gin.Default()
	r := router.Group("/user")
	r.GET("/info", h.Info)
	srv.Handle("/", router)
	err := srv.Run()
	if err != nil {
		fmt.Println(err)
	}
}
