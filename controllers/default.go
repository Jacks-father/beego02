package controllers

import (
	"beego03/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//获取name,age,sex,class
	//做数据对比
	//根据对比结果进行判断
	//1.获取请求
	//usersName := c.Ctx.Input.Query("user")
	//passerword := c.Ctx.Input.Query("psd")
	//2.使用固定数据进行数据校验
	//if usersName !="jack" || passerword !="123456789"{
		//c.Ctx.ResponseWriter.Write([]byte("对不起，数据错误"))
		//return
	//}
	//c.Ctx.ResponseWriter.Write([]byte("欢迎来到实力至上主义的教室"))

	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "18398563782@qq.com"
	c.TplName = "index.tpl"
}
func (c *MainController) post(){
	var person models.Person
	dataBytes,err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err !=nil{
		c.Ctx.WriteString("数据解析失败，请重试")
	}
	err = json.Unmarshal(dataBytes,&person)
	if err != nil {
		c.Ctx.WriteString("数据解析失败")
		return
	}
	fmt.Println("姓名",person.Name)
	fmt.Println("生日",person.Birthday)
	fmt.Println("地址",person.Address)
	fmt.Println("昵称",person.Nick)
	c.Ctx.WriteString("数据解析成功")
	//for i:=0;i<10 ;i++  {
	//fmt.Printf("第%d次打印\n",i)

	//}
	//name :=c.Ctx.Request.FormValue("N")
	//age :=c.Ctx.Request.FormValue("A")
	//sex :=c.Ctx.Request.FormValue("S")
	//fmt.Println(name,age,sex)
//if name !="flag"&& age !="20" {
	//c.Ctx.WriteString("数据错误" )
	//return
//}
	//c.Ctx.WriteString("成功")
	//var person models.Person
	//dateBytes,err :=ioutil.ReadAll(c.Ctx.Request.Body)
	//if err !=nil {
	//	c.Ctx.WriteString("重试")
	//}
	//err =json.Unmarshal(dateBytes,&person)
	//if err !=nil {
	//	c.Ctx.WriteString("重试")
	//}
	//fmt.Println("姓名：",person.Name)
	//c.Ctx.WriteString("SUCCESS")

}
