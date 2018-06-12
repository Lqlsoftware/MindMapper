package handler

import (
	"github.com/astaxie/beego"
)

type BranchController struct {
	beego.Controller
}
//
//func (this *BranchController) GET() {
//	branchId := this.GetString("branchId")
//
//	user, err := model.VaildUser(username, password)
//	if err != nil {
//		// return msg
//		this.Ctx.WriteString(utils.GetJsonResult("login failed", -1, nil))
//	}
//	// save to session
//	this.SetSession("user", user)
//	// return msg
//	this.Ctx.WriteString(utils.GetJsonResult("login success", 1, user))
//}