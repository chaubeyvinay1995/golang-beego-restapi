package main

import (
	_ "domio-api/routers"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Get the Database Configuration from conf/app.cong
	dbUser := beego.AppConfig.String("mySqlUser")
	dbPwd := beego.AppConfig.String("mySqlPass")
	dbName := beego.AppConfig.String("mySqlDb")
	// Create the database string for connection
	dbString := dbUser + ":" + dbPwd + "@/" + dbName + "?charset=utf8"
	// Register Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// Register default database
	orm.RegisterDataBase("default", "mysql", dbString)
	// autosync
	// db alias
	name := "default"
	// drop table and re-create
	force := false
	// print log
	verbose := true
	// error
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil{
		fmt.Println("Error in Database Connection", err)
	}
	// To print the Sql queries
	orm.Debug = true
	// Set the Logger
	logs.SetLogger(logs.AdapterFile,`{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,
												"daily":true,"maxdays":10,"color":true}`)
	//Run SyncDB
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
