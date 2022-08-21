package config

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "bookstore_db",
		AllowNativePasswords: true,
	}
	d, err := gorm.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(any(err))
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
