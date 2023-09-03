package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

var Con *sql.DB

func init() {

	// TODO: 環境変数から取得する
	config := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "db:3306",
		DBName:               "GM",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	con, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}

	Con = con
}
