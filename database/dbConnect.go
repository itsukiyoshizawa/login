package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func GormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(localhost:8889)"
	DBNAME := "login"

	CONNECT := DBMS + "://" + USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "charset=utf8"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
