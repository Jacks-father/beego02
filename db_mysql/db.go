package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

func init()  {
	fmt.Println("连接mysql数据库")
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")

	connUrl := dbUser +":" +dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf-8
	db,err :=sql.Open(dbDriver,connUrl)
	if err != nil{
		panic("数据库错误，请检查配置")
	}
	Db = db
}