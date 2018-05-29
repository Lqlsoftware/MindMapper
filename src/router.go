package main

import (
	"github.com/astaxie/beego"
	"./handler"
)

func bindRouter() {
	beego.Router("/login", &handler.LoginController{})
}