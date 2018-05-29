package git

import (
	"time"

	"github.com/Lqlsoftware/mindmapper/src/config"
	"github.com/Lqlsoftware/mindmapper/src/model/Tree"
	"github.com/Lqlsoftware/mindmapper/src/orm"
	"gopkg.in/mgo.v2/bson"
)

type Commit struct {
	Id			int
	Diff		Tree.MapperDiff
	Time		int64
	Title		string
	Summary		string
	Tree		Tree.MindMapperTree
}

func LoadCommit(Id int) (Commit, error) {
	commit := Commit{}
	err := orm.GetDatabase().C(config.COMMIT_CNAME).Find(bson.M{"Id": Id}).One(&commit)
	return commit, err
}

func GetLastCommitId() int {
	commit := Commit{}
	id := 1

	commits := orm.GetDatabase().C(config.COMMIT_CNAME)
	err := commits.Find(bson.M{"Id":"$max"}).One(&commit)
	if err != nil {
		id = 1
	} else {
		id = commit.Id + 1
	}
	return id
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
		err := orm.GetDatabase().C(config.COMMIT_CNAME).Insert(commit)
		if err != nil {
			return nil, nil
		}
		return &commit, nil
	}
	return nil, &conflict
}