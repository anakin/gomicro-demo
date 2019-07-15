package redisops

import (
	"demo4/user-service/config"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

var (
	redisConn redis.Conn
	err       error
)

func init() {
	cfg := config.G_cfg
	host := cfg.Redis.Host
	port := strconv.Itoa(cfg.Redis.Port)
	if redisConn, err = redis.Dial("tcp", host+":"+port); err != nil {
		panic(err.Error())
	}
}
