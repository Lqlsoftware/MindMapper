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
	user, err := this.GetUser()
	if err != nil {
		this.Ctx.WriteString(utils.GetJsonResult("Not Login", -1, nil))
		return
	}

	projectId,_ := this.GetInt("pid")
	project := git.GetBranchSet(projectId)
	master := git.GetBranch(project.MainBranchId)
	var branchs []git.Branch
	for _,v := range project.BranchIds {
		if v != master.Id {
			branch := git.GetBranch(v)
			if branch.OwnerId == user.Id {
				branchs = append(branchs, branch)
			}
		}
	}
	this.Ctx.WriteString(utils.GetJsonResult("success", 1, master, branchs))
}

func (this *BranchController) Post() {
	user, err := this.GetUser()
	if err != nil {
		this.Ctx.WriteString(utils.GetJsonResult("Not Login", -1, nil))
		return
	}

	pid,_ := this.GetInt("pid")
	name := this.GetString("name")
	branch := git.NewBranch(pid, name, user)
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