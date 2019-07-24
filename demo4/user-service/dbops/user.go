package dbops

import (
	pb "demo4/user-service/proto/user"
)

func GetUserById(id int32) (*pb.User, error) {
	user := &pb.User{
		Id: id,
	}
	DBConn.First(&user)
	return user, nil
}

func Create(user *pb.User) int32 {
	DBConn.Save(user)
	return user.Id
}
