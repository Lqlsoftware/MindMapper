package git

import (
	"github.com/Lqlsoftware/mindmapper/src/config"
	"github.com/Lqlsoftware/mindmapper/src/orm"
	"gopkg.in/mgo.v2/bson"
)

type BranchSet struct {
	TreeId			int		`json:"treeId"`
	Name			string	`json:"name"`
	MainBranchId	int		`json:"mainBranchId"`
	BranchIds		[]int	`json:"branchIds"`
	MemberIds		[]int 	`json:"memberIds"`
}

func GetBranchSets(userId int) []BranchSet {
	var branchSets []BranchSet
	err := orm.GetDatabase().C(config.BRANCHSET_CNAME).Find(bson.M{"memberids":bson.M{"$in": []int{userId}}}).All(&branchSets)
	if err != nil {
		return []BranchSet{}
	}
	return branchSets
}

func (branchSet *BranchSet)Save() error {
	return orm.GetDatabase().C(config.BRANCHSET_CNAME).Insert(branchSet)
}

func GetLastProjectId() int {
	branchSet := BranchSet{}
	err := orm.GetDatabase().C(config.BRANCHSET_CNAME).Find(bson.M{"id":"$max"}).One(&branchSet)
	if err != nil {
		return 1
	} else {
		return branchSet.TreeId + 1
	}
}

func NewBranchSet(name string, userId int) *BranchSet {
	pid, bid := GetLastProjectId(),GetLastBranchId()
	project := BranchSet{
		TreeId:			pid,
		Name:			name,
		MainBranchId:	bid,
		BranchIds:		[]int{bid},
		MemberIds:		[]int{userId},
	}
	err := project.Save()
	if err != nil {
		return nil
	}
	return &project
}