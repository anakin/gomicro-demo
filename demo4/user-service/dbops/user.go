package dbops

import (
	pb "demo4/user-service/proto/user"
)

type User struct {
	Id       int32  `gorm:"AUTO_INCREMENT,PRIMARY_KEY"`
	Name     string `gorm:"size:255"`
	Company  string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
	Password string `gorm:"size:128"`
}

func GetUserById(id int32) (*pb.User, error) {
	user := User{}
	DBConn.First(&user, id)
	u := &pb.User{
		Name:    user.Name,
		Company: user.Company,
		Email:   user.Email,
	}
	return u, nil
}

func Create(user *User) int32 {
	DBConn.Save(user)
	return user.Id
}

func GetUserByEmail(email string) (*pb.User, error) {
	user := &pb.User{}
	DBConn.Debug().Where(&User{Email: email}).First(&user)
	u := &pb.User{
		Name:     user.Name,
		Email:    user.Email,
		Company:  user.Company,
		Password: user.Password,
	}
	return u, nil
}
