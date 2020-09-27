package main

import (
	_ "beego03/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	config := beego.AppConfig
	appName := config.String("appname")
	fmt.Println("应用名称：",appName)
	port,err := config.Int("httpport")
	if err != nil{
		panic("项目配置文件失败，请检查配置文件")
	}
	//dbDriver := config.String("db_driverName")
	//dbUser := config.String("db_user")
	//dbPasserword := config.String()
		fmt.Println(port)

	beego.Run()
}