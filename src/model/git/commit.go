package git

import (
	"time"

	"github.com/Lqlsoftware/mindmapper/src/config"
	"github.com/Lqlsoftware/mindmapper/src/model"
	"github.com/Lqlsoftware/mindmapper/src/model/Tree"
	"github.com/Lqlsoftware/mindmapper/src/orm"
	"gopkg.in/mgo.v2/bson"
)

type Commit struct {
	Id			int					`json:"id"`
	Diff		Tree.MapperDiff		`json:"diff"`
	Time		int64				`json:"time"`
	Title		string				`json:"title"`
	Summary		string				`json:"summary"`
	Tree		Tree.MindMapperTree	`json:"tree"`
	Submitter	string				`json:"submitter"`
}

func LoadCommit(Id int) (Commit, error) {
	commit := Commit{}
	err := orm.GetDatabase().C(config.COMMIT_CNAME).Find(bson.M{"id": Id}).One(&commit)
	return commit, err
}

func (commit *Commit)Save() error {
	return orm.GetDatabase().C(config.COMMIT_CNAME).Insert(commit)
}

func GetLastCommitId() int {
	commit := Commit{}
	err := orm.GetDatabase().C(config.COMMIT_CNAME).Find(bson.M{"id":bson.M{"$gt":0}}).Sort("-id").Limit(1).One(&commit)
	if err != nil {
		return 1
	} else {
		return commit.Id + 1
	}
}

type Conflict struct {
	Diff 	[]Tree.MapperNodeDiff
}

func (commit *Commit)MergeWith(other *Commit) (*Commit, *Conflict) {
	conflict := Conflict{[]Tree.MapperNodeDiff{}}
	commitDiff := other.Tree.DiffWith(&commit.Tree)
	for _,diff := range commitDiff.Nodes {
		if diff.Operate != Tree.Add {
			conflict.Diff = append(conflict.Diff, diff)
		}
	}
	if len(conflict.Diff) == 0 {
		DstTree := commit.Tree
		DstTree.MergeFrom(commitDiff)
		commit := Commit{
			Id:	GetLastCommitId(),
			Diff: commitDiff,
			Time: time.Now().Unix(),
			Title: "merge from " + other.Tree.Hash,
			Summary: "",
			Tree: DstTree,
		}
		commit.Save()
		return &commit, nil
	}
	return nil, &conflict
}

func NewCommit(bid int, tree []Tree.TreeNode, title, summary string, user model.User) *Commit {
	branch := GetBranch(bid)
	preCommit, _ := LoadCommit(branch.HeadId)

	// 构造新树
	newTree := Tree.MindMapperTree{Tree: map[string]Tree.TreeNode{}, Hash: "6666"}
	for _, v := range tree {
		newTree.Tree[v.Idx] = v
	}

	// 差异
	commitDiff := newTree.DiffWith(&preCommit.Tree)

	newCommit := Commit{
		Id:			GetLastCommitId(),
		Diff:		commitDiff,
		Time:		time.Now().Unix(),
		Title:		title,
		Summary:	summary,
		Tree:		newTree,
		Submitter:	user.Username,
	}
	newCommit.Save()

	err := orm.GetDatabase().C(config.BRANCH_CNAME).Update(bson.M{"id":bid},bson.M{"$set":bson.M{"commitids":append(branch.CommitIds, newCommit.Id),"headid": newCommit.Id}})
	if err != nil {
		return &Commit{}
	}
	return &newCommit
}