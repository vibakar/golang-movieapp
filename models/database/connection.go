package database

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/movieapp")
	return db, err
}

