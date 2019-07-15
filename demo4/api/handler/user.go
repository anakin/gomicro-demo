package handler

import (
	"net/http"

	usersrv "demo4/user-service/proto/user"

	"github.com/micro/go-micro/metadata"
	opentracing "github.com/opentracing/opentracing-go"

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
	span, ctx := opentracing.StartSpanFromContext(c, "call")
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	defer span.Finish()
	opentracing.GlobalTracer().Inject(span.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md))
	ctx = opentracing.ContextWithSpan(ctx, span)
	ctx = metadata.NewContext(ctx, md)
	req := &usersrv.User{Id: 1}
	span.SetTag("req", req)
	res, err := u.userS.Get(ctx, req)
	if err != nil {
		span.SetTag("err", err)
		return
	}
	span.SetTag("resp", res)
	c.JSON(http.StatusOK, res)
}
