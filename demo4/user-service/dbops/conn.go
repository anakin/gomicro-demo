package dbops

import (
	"fmt"
	"log"

	"demo4/user-service/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	dbConn *gorm.DB
	err    error
)

func Init(address string) {
	cfg, err := config.InitConfig(address)
	if err != nil {
		log.Fatal("read config fail")
	}
	host := cfg.Db.Master.Host
	port := cfg.Db.Master.Port
	user := cfg.Db.Master.User
	password := cfg.Db.Master.Password
	dbName := cfg.Db.Master.DBName
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", user, password, host, port, dbName)
	dbConn, err = gorm.Open("mysql", str)
	if err != nil {
		log.Fatal("connect to mysql error")
	}
}
