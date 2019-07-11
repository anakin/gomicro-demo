package redisops

import "github.com/gomodule/redigo/redis"

var (
	redisConn redis.Conn
	err       error
)

func init() {
	if redisConn, err = redis.Dial("tcp", ":6779"); err != nil {
		panic(err.Error())
	}
}
