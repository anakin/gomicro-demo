package redisops

func Get(key string) (interface{}, error) {
	return redisConn.Do("GET", key)
}

func Set(key, value string) (interface{}, error) {
	return redisConn.Do("SET", key, value)
}

func LPush(q, value string) (interface{}, error) {
	return redisConn.Do("LPUSH", q, value)
}

func RPop(q string) (interface{}, error) {
	return redisConn.Do("RPOP", q)
}
