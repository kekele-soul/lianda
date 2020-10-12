package db_mysql

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
func ConnectDB() {
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbName := config.String("db_name")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_pssword")
	dbIp := config.String("db_ip")

	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	db, err := sql.Open(dbDriver, connUrl)
	if err != nil {
		panic("数据库连接错误，请检查配置")
	}
	DB = db
}