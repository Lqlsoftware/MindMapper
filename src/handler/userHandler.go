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
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "localhost")
	user := this.GetSession("user")
	if user == nil {
		this.Ctx.WriteString(utils.GetJsonResult("Not Login", -1, user))
	} else {
		this.Ctx.WriteString(utils.GetJsonResult("Welcome Back", 1, user.(model.User)))
	}
}
