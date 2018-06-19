package Tree

import (
	"encoding/json"
	"strconv"

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
	Idx		string 	`json:"idx"`
	Father	string 	`json:"father"`
	EdgeNum int		`json:"edgeNum"`
	Rank	string 	`json:"rank"`
	Value 	string 	`json:"value"`
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
	Tree	map[string]TreeNode  	`json:"tree"`
	Hash	string 					`json:"hash"`
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

func (mindMapper *MindMapperTree)ApplyDiff(diff *MapperDiff) MindMapperTree {
	// deep copy
	res := MindMapperTree{map[string]TreeNode{},"6666"}
	for k,v := range mindMapper.Tree {
		res.Tree[k] = v
	}

	// apply diff
	for _,v := range diff.Nodes {
		switch v.Operate {
		case Add:
			father := res.Tree[v.Node.Father]
			father.EdgeNum++
			v.Node.Rank = strconv.Itoa(father.EdgeNum)
			res.Tree[father.Idx] = father
			res.Tree[v.Node.Idx] = v.Node
		case Modify:
			res.Tree[v.Node.Idx] = v.Node
		case Delete:
			delete(res.Tree, v.Node.Idx)
		default:
		}
	}
	return res
}

func (mindMapper *MindMapperTree)DiffWith(other *MindMapperTree) MapperDiff {
	engine := diffmatchpatch.New()
	Diffs := MapperDiff{[]MapperNodeDiff{}}
	for _, curr := range mindMapper.Tree {
		if last, exist := other.Tree[curr.Idx]; !exist {
			// add
			valueDiff := engine.DiffMain("", curr.Value, true)
			diff := MapperNodeDiff{
				Node:      curr,
				Operate:   Add,
				Different: valueDiff,
			}
			Diffs.Nodes = append(Diffs.Nodes, diff)
		} else if curr.Value != last.Value {
			// modify
			valueDiff := engine.DiffMain(curr.Value, last.Value, true)
			diff := MapperNodeDiff{
				Node:      curr,
				Operate:   Modify,
				Different: valueDiff,
			}
			Diffs.Nodes = append(Diffs.Nodes, diff)
		}
	}
	for _, last := range other.Tree {
		if _, exist := mindMapper.Tree[last.Idx]; !exist {
			// delete
			valueDiff := engine.DiffMain(last.Value, "", true)
			diff := MapperNodeDiff{
				Node:      last,
				Operate:   Delete,
				Different: valueDiff,
			}
			Diffs.Nodes = append(Diffs.Nodes, diff)
		}
	}
	return Diffs
}