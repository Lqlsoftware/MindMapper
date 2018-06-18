package main

import (
	"fmt"

	"github.com/Lqlsoftware/mindmapper/src/orm"
	"github.com/Lqlsoftware/mindmapper/src/utils"
	"github.com/astaxie/beego"
)

func init() {
	orm.InitDB()
	beego.SetStaticPath("/static","view")
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "gosessionid"
	bindRouter()
}

func main() {
	fmt.Println(utils.GetJsonResult("s", 1, "sda"))
	fmt.Println(utils.GetJsonResult("s", 1, "sda","sda","sda"))
	return
	beego.Run()
}