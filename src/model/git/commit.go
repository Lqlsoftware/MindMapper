package git

import (
	"time"

	"github.com/Lqlsoftware/mindmapper/src/config"
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