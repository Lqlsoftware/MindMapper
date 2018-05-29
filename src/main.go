package main

import (
	"github.com/Lqlsoftware/mindmapper/src/orm"
	"github.com/astaxie/beego"
)

func main() {
	orm.InitDB()
	bindRouter()
	beego.Run()
}