package handler

import (
	"demo4/middleware"
	"net/http"

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
	res, err := u.userS.Get(c, req)
	middleware.Trace(c, req, res, err)
	c.JSON(http.StatusOK, res)
}

func (u *UserServiceHandler) Create(c *gin.Context) {

	req := &usersrv.User{Name: "anakin", Password: "123456", Company: "chope", Email: "anakin.sun@chope.co"}
	res, err := u.userS.Create(c, req)
	middleware.Trace(c, req, res, err)
	c.JSON(http.StatusOK, res)
}

func (u *UserServiceHandler) Auth(c *gin.Context) {
	req := &usersrv.User{
		Email:    "anakin.sun@chope.co",
		Password: "123456",
	}
	res, err := u.userS.Auth(c, req)
	middleware.Trace(c, req, res, err)
	c.JSON(http.StatusOK, res)
}
