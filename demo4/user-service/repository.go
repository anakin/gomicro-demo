package main

import (
	"demo4/user-service/dbops"
	pb "demo4/user-service/proto/user"

	"github.com/sirupsen/logrus"
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
	res, err := dbops.GetUserById(id)
	//logrus.Info("get user:", res)
	err1 := dbops.DBConn.Close()
	if err1 != nil {
		logrus.Error("close db error:", err1)
	}
	return res, err

}
func (u *UserRepository) GetByEmail(email string) (*pb.User, error) {
	dbops.Init()
	defer dbops.DBConn.Close()
	return dbops.GetUserByEmail(email)
}
