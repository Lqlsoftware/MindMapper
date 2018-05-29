package handler

import (
	"github.com/astaxie/beego"
	"../model"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) POST() {
	username := this.GetString("username")
	password := this.GetString("password")

	user, err := model.VaildUser(username, password)
	if err != nil {
		//
		this.Ctx.WriteString("faild")
	}

	this.SetSession("user", user)

	this.Ctx.WriteString("success")

	//if this.GetSession("user") == nil {
	//
	//}
}