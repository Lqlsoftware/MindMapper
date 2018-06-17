package git

import (
	"errors"

	"github.com/Lqlsoftware/mindmapper/src/config"
	"github.com/Lqlsoftware/mindmapper/src/orm"
	"gopkg.in/mgo.v2/bson"
)

type Branch struct {
	Id			int		`json:"id"`
	Name		string	`json:"name"`
	HeadId		int		`json:"headId"`
	CommitIds	[]int	`json:"commitIds"`
	StartTime	uint32	`json:"startTime"`
	EndTime		uint32	`json:"endTime"`
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
	err := orm.GetDatabase().C(config.BRANCHSET_CNAME).Find(bson.M{"id":"$max"}).One(&branch)
	if err != nil {
		return 1
	} else {
		return branch.Id + 1
	}
}