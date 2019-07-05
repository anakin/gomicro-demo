package dbops

import (
	pb "demo4/user-service/proto/user"
)

func GetUserById(id int32) (*pb.User, error) {
	user := &pb.User{
		Id: id,
	}
	dbConn.First(&user)
	return user, nil
}
