package main

import (
	"context"
	"fmt"

	pb "github.com/anakin/gomicro-demo/demo4/user-service/proto/user"
)

type service struct {
	repo Repository
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	fmt.Println("receiveid /user/get request", req)
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}
