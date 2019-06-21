package main

import (
	_ "01-Golangbeego/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	// codefirst创建表
	orm.RunSyncdb("default", false, true)
	if beego.BConfig.RunMode == "dev" {
		//开启swagger
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		//调试模式打印查询语句
		orm.Debug = true
	}
	beego.Run()
}

