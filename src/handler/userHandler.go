package handler

import (
	"github.com/Lqlsoftware/mindmapper/src/model"
	"github.com/Lqlsoftware/mindmapper/src/utils"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	user, err := model.GetUser(this)
	if err != nil {
		this.Ctx.WriteString(utils.GetJsonResult("Not Login", -1, nil))
	} else {
		this.Ctx.WriteString(utils.GetJsonResult("Welcome Back", 1, user))
	}
}