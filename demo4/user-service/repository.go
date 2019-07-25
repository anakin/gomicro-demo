package main

import (
	"demo4/user-service/dbops"
	pb "demo4/user-service/proto/user"
)

// Repository interface
type Repository interface {
	Get(id int32) (*pb.User, error)
	Create(info *pb.User) (*pb.User, error)
	GetByEmail(email string) (*pb.User, error)
}

//UserRepository struct
type UserRepository struct {
}

func (u *UserRepository) Create(uinfo *pb.User) (*pb.User, error) {
	dbops.Init()
	defer dbops.DBConn.Close()
	uu := &dbops.User{
		Name:     uinfo.Name,
		Company:  uinfo.Company,
		Email:    uinfo.Email,
		Password: uinfo.Password,
	}
	id, err := dbops.Create(uu)
	if err != nil {
		return nil, err
	}
	user := &pb.User{
		Id: id,
	}
	return user, nil
}

//Get get user info
func (u *UserRepository) Get(id int32) (*pb.User, error) {
	dbops.Init()
	defer dbops.DBConn.Close()
	return dbops.GetUserById(id)

}
func (u *UserRepository) GetByEmail(email string) (*pb.User, error) {
	dbops.Init()
	defer dbops.DBConn.Close()
	return dbops.GetUserByEmail(email)
}
