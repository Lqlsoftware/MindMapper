package Tree

import (
	"encoding/json"

	"github.com/sergi/go-diff/diffmatchpatch"
)

/*
	MindMapper Tree Node

	Idx 	: string	节点的编号
	Father	: string	父亲节点编号
	Rank	: string	当前父亲中的排名
	Value 	: string   	节点文字内容
 */
type TreeNode struct {
	Idx		string
	Father	string
	Rank	string
	Value 	string
}

/*
	MindMapper Liner Tree

	·-------------------------------------·
    |  Idx: string |      MapperNode      |
	·-------------------------------------·
    |  Idx: string |      MapperNode      |
	·-------------------------------------·
	|	... ...							  |
	·-------------------------------------·
    |  Idx: string |      MapperNode      |
	·-------------------------------------·
    |  Idx: string |      MapperNode      |
	·-------------------------------------·
 */
type MindMapperTree struct {
	Tree	map[string]TreeNode
	Hash	string
}

func (mindMapper *MindMapperTree)ToJson() (string,error) {
	res,err := json.Marshal(*mindMapper)
	return string(res),err
}

func (mindMapper *MindMapperTree)addNode(Node *TreeNode) {
	mindMapper.Tree[Node.Idx] = *Node
}

func (mindMapper *MindMapperTree)updateHash() {
	mindMapper.Hash = "0"
}

func (mindMapper *MindMapperTree)DiffWith(other *MindMapperTree) MapperDiff {
	engine := diffmatchpatch.New()
	Diffs := MapperDiff{[]MapperNodeDiff{}}
	for _,curr := range mindMapper.Tree {
		if last, exist := other.Tree[curr.Idx];!exist {
			// add
			valueDiff := engine.DiffMain("", curr.Value, true)
			diff := MapperNodeDiff{
				Node:		&curr,
				Operate:	Add,
				Different:	valueDiff,
			}
			Diffs.Nodes = append(Diffs.Nodes, diff)
		} else {
			// modify
			valueDiff := engine.DiffMain(curr.Value, last.Value, true)
			diff := MapperNodeDiff{
				Node:		&curr,
				Operate:	Modify,
				Different:	valueDiff,
			}
			Diffs.Nodes = append(Diffs.Nodes, diff)
		}
	}
	for _,last := range other.Tree {
		if _, exist := mindMapper.Tree[last.Idx];!exist {
			// delete
			valueDiff := engine.DiffMain(last.Value, "", true)
			diff := MapperNodeDiff{
				Node:		&last,
				Operate:	Delete,
				Different:	valueDiff,
			}
			Diffs.Nodes = append(Diffs.Nodes, diff)
		}
	}
	return Diffs
}

func (mindMapper *MindMapperTree)MergeFrom(Diff MapperDiff) {
	for _, add := range Diff.Nodes {
		mindMapper.addNode(add.Node)
	}
	mindMapper.updateHash()
}