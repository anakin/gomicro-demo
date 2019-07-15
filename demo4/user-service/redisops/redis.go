package redisops

import (
	"github.com/gomodule/redigo/redis"
)

//Get get value from redis
func Get(key string) (interface{}, error) {
	return redisConn.Do("GET", key)
}

//Set set value into redis
func Set(key, value string) (interface{}, error) {
	return redis.String(redisConn.Do("SET", key, value))
}

//LPush insert in q
func LPush(q, value string) (interface{}, error) {
	return redisConn.Do("LPUSH", q, value)
}

//RPop pop from q
func RPop(q string) (interface{}, error) {
	return redis.String(redisConn.Do("RPOP", q))
}
