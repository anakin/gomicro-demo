package handler

import (
	"net/http"

	usersrv "github.com/anakin/gomicro-demo/demo4/user-service/proto/user"
	"github.com/gin-gonic/gin"
	"github.com/go-log/log"
	"github.com/micro/go-micro/client"
)

type UserServiceHandler struct {
	userS usersrv.UserService
}

func New(client client.Client) *UserServiceHandler {
	return &UserServiceHandler{
		userS: usersrv.NewUserService("chope.co.srv.user", client),
	}
}

func (u *UserServiceHandler) Info(c *gin.Context) {
	log.Log("received /user/info request")
	res, err := u.userS.Get(c, &usersrv.User{Id: 1})
	if err != nil {
		log.Log(err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}
