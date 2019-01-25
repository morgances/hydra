package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	databaseConn *sql.DB
)

func init() {
	var (
		err error
	)

	databaseConn, err = sql.Open("mysql", "root:111111@tcp(127.0.0.1:3306)/main?charset=utf8")
	if err != nil {
		panic(err)
	}
}
