package handler

import (
	"errors"

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

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Get() {
	_, err := this.GetUser()
	if err != nil {
		// return msg
		this.Ctx.WriteString(utils.GetJsonResult("not login", -1, nil))
		return
	}
	// save to session
	this.DelSession("user")
	// return msg
	this.Ctx.WriteString(utils.GetJsonResult("logout success", 1, nil))
}

func (this *LogoutController)GetUser() (model.User, error) {
	user := this.GetSession("user")
	if user == nil {
		return model.User{}, errors.New("not login")
	} else {
		return user.(model.User), nil
	}
}