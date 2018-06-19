package git

import (
	"errors"
	"time"

	"github.com/Lqlsoftware/mindmapper/src/config"
	"github.com/Lqlsoftware/mindmapper/src/model"
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
	MergeIds	[]int
	OwnerId		int		`json:"ownerId"`
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

func NewBranch(pid int, name string, user model.User) Branch {
	project := GetBranchSet(pid)
	master := GetBranch(project.MainBranchId)
	branch := Branch{
		Id:			GetLastBranchId(),
		Name: 		name,
		HeadId: 	master.HeadId,
		CommitIds:	[]int{master.HeadId},
		StartTime: 	time.Now().Unix(),
		EndTime:	-1,
		MergeIds:	[]int{master.HeadId},
		OwnerId:	user.Id,
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

func (branch *Branch)MergeWith(other *Branch, user model.User) (*Commit, *Conflict, error) {
	// 获取 Head Commit
	masterHead, err := LoadCommit(branch.HeadId)
	if err != nil {
		return nil, nil, errors.New("Wrong Commit Id ")
	}
	branchHead, err := LoadCommit(other.HeadId)
	if err != nil {
		return nil, nil, errors.New("Wrong Commit Id ")
	}

	// 获取公共祖先 base commit
	base, err := LoadCommit(other.MergeIds[len(other.MergeIds) - 1])


	// 三路合并
	commit, conflict := Merge(&masterHead, &branchHead, &base, user)
	if commit == nil {
		return nil, conflict, nil
	}
	return commit, nil, nil
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