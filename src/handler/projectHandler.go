package handler

import (
	"errors"

	"github.com/Lqlsoftware/mindmapper/src/model"
	"github.com/Lqlsoftware/mindmapper/src/model/git"
	"github.com/Lqlsoftware/mindmapper/src/utils"
	"github.com/astaxie/beego"
)

type ProjectController struct {
	beego.Controller
}

func (this *ProjectController) Get() {
	user, err := this.GetUser()
	if err != nil {
		this.Ctx.WriteString(utils.GetJsonResult("Not Login", -1, nil))
		return
	}

	// 获取project列表
	projects := git.GetBranchSets(user.Id)
	this.Ctx.WriteString(utils.GetJsonResult("list of project", 1, projects))
}

func (this *ProjectController) Post() {
	user, err := this.GetUser()
	if err != nil {
		this.Ctx.WriteString(utils.GetJsonResult("Not Login", -1, nil))
		return
	}

	name := this.GetString("name")
	project := git.NewBranchSet(name, user)
	this.Ctx.WriteString(utils.GetJsonResult("new project", 1, project))
}

func (this *ProjectController)GetUser() (model.User, error) {
	user := this.GetSession("user")
	if user == nil {
		return model.User{}, errors.New("not login")
	} else {
		return user.(model.User), nil
	}
}

type ProjectMemberController struct {
	beego.Controller
}

func (this *ProjectMemberController) Post() {
	//_, err := this.GetUser()
	//if err != nil {
	//	this.Ctx.WriteString(utils.GetJsonResult("Not Login", -1, nil))
	//	return
	//}

	username := this.GetString("name")
	pid,_ := this.GetInt("pid")
	project := git.GetBranchSet(pid)
	err := project.AddUser(username)
	if err != nil {
		this.Ctx.WriteString(utils.GetJsonResult("invalid username", -1, nil))
		return
	}
	this.Ctx.WriteString(utils.GetJsonResult("ok", 1, nil))
}

func (this *ProjectMemberController)GetUser() (model.User, error) {
	user := this.GetSession("user")
	if user == nil {
		return model.User{}, errors.New("not login")
	} else {
		return user.(model.User), nil
	}
}