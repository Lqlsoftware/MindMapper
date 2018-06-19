package main

import (
	"github.com/Lqlsoftware/mindmapper/src/handler"
	"github.com/astaxie/beego"
)

func bindRouter() {
	beego.Router("/login", &(handler.LoginController{}))
	beego.Router("/user", &(handler.UserController{}))
	beego.Router("/project", &(handler.ProjectController{}))
	beego.Router("/projectMember", &(handler.ProjectMemberController{}))
	beego.Router("/branch", &(handler.BranchController{}))
	beego.Router("/commit", &(handler.CommitController{}))
	beego.Router("/commits", &(handler.CommitsController{}))
	beego.Router("/merge", &(handler.MergerController{}))
}