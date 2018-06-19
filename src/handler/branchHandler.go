package handler

import (
	"errors"

	"github.com/Lqlsoftware/mindmapper/src/model"
	"github.com/Lqlsoftware/mindmapper/src/model/git"
	"github.com/Lqlsoftware/mindmapper/src/utils"
	"github.com/astaxie/beego"
)

type BranchController struct {
	beego.Controller
}

func (this *BranchController) Get() {
	_, err := this.GetUser()
	if err != nil {
		this.Ctx.WriteString(utils.GetJsonResult("Not Login", -1, nil))
		return
	}

	projectId,_ := this.GetInt("pid")
	project := git.GetBranchSet(projectId)
	master := git.GetBranch(project.MainBranchId)
	var branch []git.Branch
	for _,v := range project.BranchIds {
		if v != master.Id {
			branch = append(branch, git.GetBranch(v))
		}
	}
	this.Ctx.WriteString(utils.GetJsonResult("success", 1, master, branch))
}

func (this *BranchController) Post() {
	_, err := this.GetUser()
	if err != nil {
		this.Ctx.WriteString(utils.GetJsonResult("Not Login", -1, nil))
		return
	}

	pid,_ := this.GetInt("pid")
	name := this.GetString("name")
	branch := git.NewBranch(pid, name)
	this.Ctx.WriteString(utils.GetJsonResult("success", 1, branch))
}

func (this *BranchController)GetUser() (model.User, error) {
	user := this.GetSession("user")
	if user == nil {
		return model.User{}, errors.New("not login")
	} else {
		return user.(model.User), nil
	}
}