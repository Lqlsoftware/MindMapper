package main

import (
	"fmt"
	"time"

	"github.com/Lqlsoftware/mindmapper/src/model/Tree"
	"github.com/Lqlsoftware/mindmapper/src/model/git"
	"github.com/Lqlsoftware/mindmapper/src/orm"
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
	beego.Run()
}