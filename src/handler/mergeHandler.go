package handler

import (
	"errors"
	"time"

	"github.com/Lqlsoftware/mindmapper/src/config"
	"github.com/Lqlsoftware/mindmapper/src/model"
	"github.com/Lqlsoftware/mindmapper/src/model/git"
	"github.com/Lqlsoftware/mindmapper/src/orm"
	"github.com/Lqlsoftware/mindmapper/src/utils"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

type MergerController struct {
	beego.Controller
}

func (this *MergerController) Get() {
	user, err := this.GetUser()
	if err != nil {
		this.Ctx.WriteString(utils.GetJsonResult("Not Login", -1, nil))
		return
	}

	pid, _ := this.GetInt("pid")
	bid, _ := this.GetInt("bid")
	project := git.GetBranchSet(pid)
	branch := git.GetBranch(bid)
	master := git.GetBranch(project.MainBranchId)
	head,_ := git.LoadCommit(master.HeadId)
	commit := git.Commit{
		Id:	git.GetLastCommitId(),
		Diff: head.Diff,
		Time: time.Now().Unix(),
		Title: "pull from 'master'",
		Tree: head.Tree,
		Submitter: user.Username,
	}
	commit.Save()
	orm.GetDatabase().C(config.BRANCH_CNAME).Update(bson.M{"id":branch.Id}, bson.M{"$set":bson.M{"headid": commit.Id,"commitids": append(branch.CommitIds, commit.Id)}})
	this.Ctx.WriteString(utils.GetJsonResult("success", 1, nil))
}

func (this *MergerController) Post() {
	user, err := this.GetUser()
	if err != nil {
		this.Ctx.WriteString(utils.GetJsonResult("Not Login", -1, nil))
		return
	}

	pid, _ := this.GetInt("pid")
	bid, _ := this.GetInt("bid")

	project := git.GetBranchSet(pid)
	branch := git.GetBranch(bid)
	master := git.GetBranch(project.MainBranchId)
	commit, conflict, err := master.MergeWith(&branch, user)
	if err != nil {
		this.Ctx.WriteString(utils.GetJsonResult("err occur", -2, nil))
		return
	} else if commit == nil {
		this.Ctx.WriteString(utils.GetJsonResult("conflict need to be fix", 2, conflict))
		return
	}
	orm.GetDatabase().C(config.BRANCH_CNAME).Update(bson.M{"id":master.Id}, bson.M{"$set":bson.M{"headid": commit.Id,"commitids": append(master.CommitIds, commit.Id)}})
	orm.GetDatabase().C(config.BRANCH_CNAME).Update(bson.M{"id":bid}, bson.M{"$set":bson.M{"mergeids": append(branch.MergeIds, commit.Id)}})
	this.Ctx.WriteString(utils.GetJsonResult("success", 1, commit))
	return
}

func (this *MergerController)GetUser() (model.User, error) {
	user := this.GetSession("user")
	if user == nil {
		return model.User{}, errors.New("not login")
	} else {
		return user.(model.User), nil
	}
}