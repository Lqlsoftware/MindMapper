package git

import (
	"strconv"
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

// 三路合并算法
func Merge(dst, src, base *Commit) (*Commit, *Conflict) {
	dstDiff := dst.Tree.DiffWith(&base.Tree)
	srcDiff := src.Tree.DiffWith(&base.Tree)

	// 使用的Id
	last := 0
	usedId := map[string]bool{}
	for k := range dst.Tree.Tree {
		usedId[k] = true
	}
	for k := range src.Tree.Tree {
		usedId[k] = true
	}
	for k := range base.Tree.Tree {
		usedId[k] = true
	}
	// 分配不同的id
	for _,v1 := range srcDiff.Nodes {
		for _,v2 := range dstDiff.Nodes {
			// 存在id相同 父亲不同的节点
			if v1.Node.Idx == v2.Node.Idx && v1.Node.Father != v2.Node.Father {
				// 获取一个不重复的id
				for i := last + 1; ; i++ {
					if _, ok := usedId[strconv.Itoa(i)]; !ok {
						last = i
						break
					}
				}
				// 对v1分配新的id
				oldId := v2.Node.Idx
				node := src.Tree.Tree[oldId]
				delete(src.Tree.Tree, oldId)
				node.Idx = strconv.Itoa(last)
				src.Tree.Tree[node.Idx] = node
				// 更新孩子节点
				for k, v := range src.Tree.Tree {
					if v.Father == oldId {
						v.Father = node.Idx
						src.Tree.Tree[k] = v
					}
				}
			}
		}
	}

	// 求差异的交集
	srcDiff = src.Tree.DiffWith(&base.Tree)
	conflict := Conflict{[]Tree.MapperNodeDiff{}}
	apply := Tree.MapperDiff{Nodes:make([]Tree.MapperNodeDiff, len(srcDiff.Nodes))}
	copy(apply.Nodes, srcDiff.Nodes)
	for i,v1 := range srcDiff.Nodes {
		for _,v2 := range dstDiff.Nodes {
			// 存在相同差异节点
			if v1.Node.Idx == v2.Node.Idx {
				if v1.Equal(&v2) {
					apply.Nodes[i].Operate = Tree.Hide
				} else {
					conflict.Diff = append(conflict.Diff, v1, v2)
				}
			}
		}
	}
	if len(conflict.Diff) != 0 {
		return nil, &conflict
	}

	resTree := dst.Tree.ApplyDiff(&apply)
	res := resTree.DiffWith(&dst.Tree)
	commit := Commit{
		Id:	GetLastCommitId(),
		Diff: res,
		Time: time.Now().Unix(),
		Title: "merge from '" + src.Title + "'",
		Summary: "",
		Tree: resTree,
	}
	commit.Save()
	return &commit, nil
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