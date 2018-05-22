package handler

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Ctx.WriteString("hello world")
}