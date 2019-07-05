package main

import (
	"demo4/user-service/dbops"
	pb "demo4/user-service/proto/user"
)

type Repository interface {
	Get(id int32) (*pb.User, error)
}

type UserRepository struct {
}

func (u *UserRepository) Get(id int32) (*pb.User, error) {
	res, err := dbops.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
