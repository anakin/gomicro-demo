package config

import (
	"fmt"
	"sync"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/consul"
)

//DBConfig Database config
type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

//AllDBConfig include master and slave
type AllDBConfig struct {
	Master DBConfig `json:"master"`
	Slave  DBConfig `json:"slave"`
}

//RedisConfig Redis config info
type RedisConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

//MemcacheConfig memcache config
type MemcacheConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

//Cfg config struct
type Cfg struct {
	lock     sync.Mutex
	conf     config.Config
	Db       AllDBConfig    `json:"db"`
	Redis    RedisConfig    `json:"redis"`
	Memcache MemcacheConfig `json:"memcache`
}

//InitConfig ,init the config
func InitConfig(address string) (*Cfg, error) {
	consulSource := consul.NewSource(consul.WithAddress(address))
	conf := config.NewConfig()

	// Load file source
	err := conf.Load(consulSource)
	if err != nil {
		fmt.Println("load err:", err)
	}
	var c Cfg
	err = conf.Get("micro", "config", "database", "user").Scan(&c)
	if err != nil {
		fmt.Println("get error:", err)
	}
	return &c, nil
}
