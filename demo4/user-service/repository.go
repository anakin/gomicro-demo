package main

import (
	"demo4/user-service/dbops"
	pb "demo4/user-service/proto/user"
)

// Repository interface
type Repository interface {
	Get(id int32) (*pb.User, error)
	Create(info *pb.User) (*pb.User, error)
}

//UserRepository struct
type UserRepository struct {
}

func (u *UserRepository) Create(uinfo *pb.User) (*pb.User, error) {
	dbops.Init()
	defer dbops.DBConn.Close()

	id := dbops.Create(uinfo)
	user := &pb.User{
		Id: id,
	}
	return user, nil
	//return &pb.User{
	//	Id:      123,
	//	Name:    uinfo.Name,
	//	Company: uinfo.Company,
	//	Email:   uinfo.Email,
	//}, nil
}

//Get get user info
func (u *UserRepository) Get(id int32) (*pb.User, error) {
	dbops.Init()
	defer dbops.DBConn.Close()
	res, err := dbops.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
