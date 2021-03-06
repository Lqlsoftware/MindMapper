package git

import (
	"time"

	"github.com/Lqlsoftware/mindmapper/src/config"
	"github.com/Lqlsoftware/mindmapper/src/model"
	"github.com/Lqlsoftware/mindmapper/src/model/Tree"
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

func (branchSet *BranchSet)AddUser(username string) error {
	user := model.User{}
	err := orm.GetDatabase().C(config.USER_CNAME).Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return err
	}

	err = orm.GetDatabase().C(config.BRANCHSET_CNAME).Update(bson.M{"treeid": branchSet.TreeId},bson.M{"$set":bson.M{"memberids": append(branchSet.MemberIds, user.Id)}})
	if err != nil {
		return err
	}
	return nil
}

func GetBranchSets(userId int) []BranchSet {
	var branchSets []BranchSet
	err := orm.GetDatabase().C(config.BRANCHSET_CNAME).Find(bson.M{"memberids":bson.M{"$in": []int{userId}}}).All(&branchSets)
	if err != nil {
		return []BranchSet{}
	}
	return branchSets
}

func GetBranchSet(pid int) BranchSet {
	var branchSet BranchSet
	err := orm.GetDatabase().C(config.BRANCHSET_CNAME).Find(bson.M{"treeid":pid}).One(&branchSet)
	if err != nil {
		return BranchSet{}
	}
	return branchSet
}

func (branchSet *BranchSet)Save() error {
	return orm.GetDatabase().C(config.BRANCHSET_CNAME).Insert(branchSet)
}

func GetLastProjectId() int {
	branchset := BranchSet{}
	err := orm.GetDatabase().C(config.BRANCHSET_CNAME).Find(bson.M{"treeid":bson.M{"$gt":0}}).Sort("-treeid").Limit(1).One(&branchset)
	if err != nil {
		return 1
	} else {
		return branchset.TreeId + 1
	}
}

func NewBranchSet(name string, user model.User) BranchSet {
	pid := GetLastProjectId()

	// origin commit
	commit := Commit{
		Id:			GetLastCommitId(),
		Diff:		Tree.MapperDiff{
			Nodes: []Tree.MapperNodeDiff{},
		},
		Time:		time.Now().Unix(),
		Title:		"orign commit",
		Summary:	"orign commit",
		Tree:		Tree.MindMapperTree{
			Tree: map[string]Tree.TreeNode{"0":{"0","-1",0,1,-1,""}},
			Hash: "6666",
		},
		Submitter:	user.Username,
	}
	err := commit.Save()
	if err != nil {
		return BranchSet{}
	}

	// new branch
	master := Branch{
		Id:			GetLastBranchId(),
		Name: 		"master",
		HeadId: 	commit.Id,
		CommitIds:	[]int{commit.Id},
		StartTime: 	time.Now().Unix(),
		EndTime:	-1,
	}
	err = master.Save()
	if err != nil {
		return BranchSet{}
	}

	// project
	project := BranchSet{
		TreeId:			pid,
		Name:			name,
		MainBranchId:	master.Id,
		BranchIds:		[]int{master.Id},
		MemberIds:		[]int{user.Id},
	}
	err = project.Save()
	if err != nil {
		return BranchSet{}
	}
	return project
}