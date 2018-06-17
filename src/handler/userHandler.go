package handler

import (
	"errors"

	"github.com/Lqlsoftware/mindmapper/src/model"
	"github.com/Lqlsoftware/mindmapper/src/utils"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	user, err := this.GetUser()
	if err != nil {
		this.Ctx.WriteString(utils.GetJsonResult("Not Login", -1, nil))
	} else {
		this.Ctx.WriteString(utils.GetJsonResult("Welcome Back", 1, user))
	}
}

func (this *UserController)GetUser() (model.User, error) {
	user := this.GetSession("user")
	if user == nil {
		return model.User{}, errors.New("not login")
	} else {
		return user.(model.User), nil
	}
}