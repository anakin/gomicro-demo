package dbops

import (
	"demo4/middleware"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DBConn *gorm.DB
	err    error
)

func Init() {
	cfg := middleware.G_cfg

	host := cfg.Db.Master.Host
	port := cfg.Db.Master.Port
	user := cfg.Db.Master.User
	password := cfg.Db.Master.Password
	dbName := cfg.Db.Master.DBName
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", user, password, host, port, dbName)
	DBConn, err = gorm.Open("mysql", str)
	if err != nil {
		log.Println("connect to mysql error")
	}
	DBConn.DB().SetMaxOpenConns(100)
	DBConn.DB().SetConnMaxLifetime(time.Hour)
	DBConn.DB().SetMaxIdleConns(10)
	DBConn.Set("gorm:table_options", "ENGINE=InnoDB,CHARSET=utf8mb4").AutoMigrate(User{}).AddUniqueIndex("idx_email", "email")
}
