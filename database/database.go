package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func Connect() {
	db, err := sql.Open("mysql", "wish_list:LVA7ECV3ucv3MGDp@tcp(mysql-db.bredbrains.tech)/wish_list")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
