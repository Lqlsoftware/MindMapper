package handler

import (
	"errors"

	"github.com/Lqlsoftware/mindmapper/src/model"
	"github.com/astaxie/beego"
)

type CommitController struct {
	beego.Controller
}

func (this *CommitController) Get() {
	_, err := this.GetUser()
	if err != nil {
		this.Ctx.WriteString(utils.GetJsonResult("Not Login", -1, nil))
		return
	}
}

func (this *CommitController)GetUser() (model.User, error) {
	user := this.GetSession("user")
	if user == nil {
		return model.User{}, errors.New("not login")
	} else {
		return user.(model.User), nil
	}
}