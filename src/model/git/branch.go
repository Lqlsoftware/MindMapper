package git

import (
	"errors"
	"time"

	"github.com/Lqlsoftware/mindmapper/src/config"
	"github.com/Lqlsoftware/mindmapper/src/orm"
	"gopkg.in/mgo.v2/bson"
)

type Branch struct {
	Id			int		`json:"id"`
	Name		string	`json:"name"`
	HeadId		int		`json:"headId"`
	CommitIds	[]int	`json:"commitIds"`
	StartTime	int64	`json:"startTime"`
	EndTime		int64	`json:"endTime"`
	MergeId		[]int
}

func GetBranch(branchId int) Branch {
	branch := Branch{}
	err := orm.GetDatabase().C(config.BRANCH_CNAME).Find(bson.M{"id":branchId}).One(&branch)
	if err != nil {
		return Branch{}
	} else {
		return branch
	}
}

func NewBranch(pid int, name string) Branch {
	project := GetBranchSet(pid)
	master := GetBranch(project.MainBranchId)
	branch := Branch{
		Id:			GetLastBranchId(),
		Name: 		name,
		HeadId: 	master.HeadId,
		CommitIds:	[]int{master.HeadId},
		StartTime: 	time.Now().Unix(),
		EndTime:	-1,
	}

	err := branch.Save()
	if err != nil {
		return Branch{}
	}

	err = orm.GetDatabase().C(config.BRANCHSET_CNAME).Update(bson.M{"treeid":pid},bson.M{"$set":bson.M{"branchids":append(project.BranchIds, branch.Id)}})
	if err != nil {
		return Branch{}
	}
	return branch
}

func (branch *Branch)MergeWith(other *Branch) (Commit, error) {
	// 获取 Head Commit
	dstHead, err := LoadCommit(branch.HeadId)
	if err != nil {
		return Commit{}, errors.New("Wrong Commit Id ")
	}
	srcHead, err := LoadCommit(other.HeadId)
	if err != nil {
		return Commit{}, errors.New("Wrong Commit Id ")
	}

	// 尝试 Merge 两颗树
	dstHead.MergeWith(&srcHead)
	return dstHead, nil
}

func GetLastBranchId() int {
	branch := Branch{}
	err := orm.GetDatabase().C(config.BRANCH_CNAME).Find(bson.M{"id":bson.M{"$gt":0}}).Sort("-id").Limit(1).One(&branch)
	if err != nil {
		return 1
	} else {
		return branch.Id + 1
	}
}

func (branch *Branch)Save() error {
	return orm.GetDatabase().C(config.BRANCH_CNAME).Insert(branch)
}