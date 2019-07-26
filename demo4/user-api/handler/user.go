package handler

import (
	"demo4/lib/tracer"
	libtracer "demo4/lib/wrapper/tracer"
	"net/http"

	"github.com/sirupsen/logrus"

	usersrv "demo4/user-service/proto/user"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
)

const UserServiceName = "chope.co.srv.user"

type UserServiceHandler struct {
	userS usersrv.UserService
}

func New(client client.Client) *UserServiceHandler {
	return &UserServiceHandler{
		userS: usersrv.NewUserService(UserServiceName, client),
	}
}

func (u *UserServiceHandler) Info(c *gin.Context) {

	req := &usersrv.User{Id: 1}
	ctx, ok := libtracer.ContextWithSpan(c)
	if !ok {
		logrus.Error("context error")
	}
	res, err := u.userS.Get(ctx, req)
	tracer.Trace(ctx, "Info", req, res, err)
	if err != nil {
		logrus.Error(err)
	}
	c.JSON(http.StatusOK, res)
}

func (u *UserServiceHandler) Create(c *gin.Context) {

	req := &usersrv.User{Name: "anakin", Password: "123456", Company: "chope", Email: "anakin.sun@chope.co"}
	ctx, ok := libtracer.ContextWithSpan(c)
	if !ok {
		logrus.Error("context error")
	}
	res, err := u.userS.Create(ctx, req)
	if err != nil {
		logrus.Error(err)
	}
	tracer.Trace(ctx, "Create", req, res, err)
	c.JSON(http.StatusOK, res)
}

func (u *UserServiceHandler) Auth(c *gin.Context) {
	req := &usersrv.User{
		Email:    "anakin.sun@chope.co",
		Password: "123456",
	}
	ctx, ok := libtracer.ContextWithSpan(c)
	if !ok {
		logrus.Error("context error")
	}
	res, err := u.userS.Auth(ctx, req)
	tracer.Trace(ctx, "Auth", req, res, err)
	if err != nil {
		logrus.Error(err)
	}
	c.JSON(http.StatusOK, res)
}
