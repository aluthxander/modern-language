package app

import (
	"belajar-12-resful-api/helper"
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql" // driver MySQL
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajargodb")
	helper.PanicError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}