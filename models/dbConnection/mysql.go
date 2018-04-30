package dbConnection

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/astaxie/beego"
)

func ConnectToMysql() (*sql.DB, error) {
	dbUsername := beego.AppConfig.String("dbUsername")
	dbName := beego.AppConfig.String("dbName")
	db, err := sql.Open("mysql", dbUsername+":@/"+dbName)
	return db, err
}

