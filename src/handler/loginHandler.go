package handler

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (this *LoginController) POST() {
	username := this.GetStrings("username")
	password := this.GetStrings("password")


	this.Ctx.WriteString("hello world")
}