package handler

import (
	"github.com/Lqlsoftware/mindmapper/src/model"
	"github.com/Lqlsoftware/mindmapper/src/utils"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")

	user, err := model.VaildUser(username, password)
	if err != nil {
		// return msg
		this.Ctx.WriteString(utils.GetJsonResult("login failed", -1, nil))
		return
	}
	// save to session
	this.SetSession("user", user)
	// return msg
	this.Ctx.WriteString(utils.GetJsonResult("login success", 1, user))
}