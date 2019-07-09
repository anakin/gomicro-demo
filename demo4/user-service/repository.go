package main

import (
	"demo4/user-service/dbops"
	pb "demo4/user-service/proto/user"
)

// Repository interface
type Repository interface {
	Get(id int32) (*pb.User, error)
}

//UserRepository struct
type UserRepository struct {
}

//Get get user info
func (u *UserRepository) Get(id int32) (*pb.User, error) {
	res, err := dbops.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
