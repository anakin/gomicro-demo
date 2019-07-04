package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
)

func main() {
	reg := consul.NewRegistry()
	srv := web.NewService(
		web.Name("anakin.sun.api.user"),
		web.Registry(reg),
	)
	srv.Init()
	router := gin.Default()
	r := router.Group("/user")
	r.GET("/test", testhandler)
	srv.Handle("/", router)
	srv.Run()
}

func testhandler(c *gin.Context) {
	fmt.Println("received user/test request")
}
