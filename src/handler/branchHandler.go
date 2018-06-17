package handler

import (
	"github.com/Lqlsoftware/mindmapper/src/utils"
	"github.com/astaxie/beego"
)

type BranchController struct {
	beego.Controller
}

func (this *BranchController) Get() {
	user := this.GetSession("user")
	if user == nil {
		this.Ctx.WriteString(utils.GetJsonResult("Not Login", -1, user))
		return
	}

	//branchId := this.GetString("branchId")
}