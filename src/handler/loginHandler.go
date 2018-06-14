package handler

import (
	"fmt"
	"time"

	"github.com/Lqlsoftware/mindmapper/src/model"
	"github.com/Lqlsoftware/mindmapper/src/model/Tree"
	"github.com/Lqlsoftware/mindmapper/src/model/git"
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

func (this *LoginController) Get() {
	dstCommit, _ := git.LoadCommit(0)
	commit := git.Commit{
		Id:			git.GetLastCommitId(),
		Diff:		Tree.MapperDiff{},
		Time:		time.Now().Unix(),
		Title:		"third commit",
		Summary:	"hhhhh",
		Tree:		Tree.MindMapperTree{
			Tree: 	map[string]Tree.TreeNode{
				"0": {"0","-1", "1","root"},
				"1": {"1","0", "1","child1"},
				"2": {"2","0", "2","asdasd"},
				"3": {"3","1", "1","gg"},
				"4": {"4","2", "1","errer"},
			},
			Hash: 	"777",
		},
	}
	resCommit, conflict := dstCommit.MergeWith(&commit)
	fmt.Println(resCommit)
	fmt.Println(conflict)
	this.Ctx.WriteString(utils.GetJsonResult("fuck you", 1, resCommit))
	return
}