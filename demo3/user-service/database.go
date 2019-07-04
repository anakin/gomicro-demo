package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func CreateConnection() (*gorm.DB, error) {
	//host := os.Getenv("DB_HOST")
	//user := os.Getenv("DB_USER")
	//DBName := os.Getenv("DB_NAME")
	//password := os.Getenv("DB_PASSWORD")
	host := "127.0.0.1"
	user := "testuser"
	password := "test123"
	DBName := "test"
	str := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", user, password, host, DBName)
	db, err := gorm.Open("mysql", str)
	if err != nil {
		return nil, err
	}
	return db, nil
}
