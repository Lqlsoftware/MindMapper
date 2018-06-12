package main

import (
	"github.com/Lqlsoftware/mindmapper/src/handler"
	"github.com/astaxie/beego"
)

func bindRouter() {
	beego.Router("/login", &(handler.LoginController{}))
	beego.Router("/user", &(handler.UserController{}))
}