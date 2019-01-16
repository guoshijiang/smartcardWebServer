package db

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"fmt"
)

var db *sql.DB

func initDB(addr string) error {
	var err error
	db, err = sql.Open("mysql", addr)
	return err
}

func Init(user, passwd, ip, port, db string) error {
	dbAddr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		user, passwd, ip, port, db)
	return initDB(dbAddr)
}

func GetDB() *sql.DB {
	return db
}

