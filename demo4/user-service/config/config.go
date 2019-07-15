package config

import (
	"fmt"
	"sync"

	"github.com/micro/go-micro/config/source/consul"

	"github.com/micro/go-micro/config/source/file"

	"github.com/micro/go-micro/config"
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

//JaegerConfig opentracing endpoint
type JaegerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

//Cfg config struct
type Cfg struct {
	lock     sync.Mutex
	conf     config.Config
	Db       AllDBConfig    `json:"db"`
	Redis    RedisConfig    `json:"redis"`
	Memcache MemcacheConfig `json:"memcache"`
	Jaeger   JaegerConfig   `json:"jaeger"`
}

var (
	G_cfg *Cfg
	err   error
)

//consul
func InitWithConsul(address string) {
	consulSource := consul.NewSource(consul.WithAddress(address))
	conf := config.NewConfig()
	// Load file source
	err := conf.Load(consulSource)
	if err != nil {
		fmt.Println("load err:", err)
	}
	c := Cfg{}
	err = conf.Get("micro", "config", "database", "user").Scan(&c)
	if err != nil {
		fmt.Println("get error:", err)
	}
	G_cfg = &c
}

//InitWithFile,init the config from file
func InitWithFile(filepath string) {
	conf := config.NewConfig()
	fileSource := file.NewSource(
		file.WithPath(filepath),
	)
	// Load file source
	err := conf.Load(fileSource)
	if err != nil {
		fmt.Println("load err:", err)
	}
	c := Cfg{}
	err = conf.Scan(&c)
	if err != nil {
		fmt.Println("get error:", err)
	}
	G_cfg = &c
}
