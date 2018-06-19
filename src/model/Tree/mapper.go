package Tree

import (
	"encoding/json"
	"sort"
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
	Rank	int 	`json:"rank"`
	PreBro	int		`json:"preBro"`
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
	for i := len(diff.Nodes) - 1;i >= 0;i-- {
		v := diff.Nodes[i]
		switch v.Operate {
		case Add:
			// father
			father := res.Tree[v.Node.Father]
			father.EdgeNum++
			v.Node.Rank = father.EdgeNum
			v.Node.EdgeNum = 0
			res.Tree[father.Idx] = father
			// prebro
			v.Node.PreBro = -1
			for _,v1 := range res.Tree {
				if v1.Father == father.Idx && v1.Rank == v.Node.Rank - 1 {
					v.Node.PreBro,_ = strconv.Atoi(v1.Idx)
					break
				}
			}
			res.Tree[v.Node.Idx] = v.Node
		case Modify:
			if _,ok := res.Tree[v.Node.Idx];!ok {
				continue
			}
			target := res.Tree[v.Node.Idx]
			target.Value = v.Node.Value
			res.Tree[v.Node.Idx] = target
		case Delete:
			if _,ok := res.Tree[v.Node.Idx];!ok {
				continue
			}
			// 递归删除元素
			res.DelChild(v.Node.Idx)

			father := res.Tree[v.Node.Father]
			father.EdgeNum--
			res.Tree[father.Idx] = father
			for k,v1 := range res.Tree {
				if v1.Father == father.Idx {
					if v1.Rank > v.Node.Rank {
						v1.Rank--
						res.Tree[k] = v1
					}
					if v1.Rank == v.Node.Rank + 1 {
						v1.PreBro = v.Node.PreBro
						res.Tree[k] = v1
					}
				}
			}
		default:
		}
	}
	return res
}

func (mindMapper *MindMapperTree)DelChild(idx string) {
	node := mindMapper.Tree[idx]
	delete(mindMapper.Tree, idx)
	for k,v1 := range mindMapper.Tree {
		if v1.Father == node.Father {
			mindMapper.DelChild(k)
		}
	}
}


func (mindMapper *MindMapperTree)DiffWith(other *MindMapperTree) MapperDiff {
	engine := diffmatchpatch.New()

	var list []TreeNode
	for _, curr := range mindMapper.Tree {
		list = append(list, curr)
	}
	sort.Slice(list, func(i, j int) bool {
		v1,_ := strconv.Atoi(list[i].Idx)
		v2,_ := strconv.Atoi(list[j].Idx)
		return v1 < v2
	})
	
	Diffs := MapperDiff{[]MapperNodeDiff{}}
	for _, curr := range list {
		if last, exist := other.Tree[curr.Idx]; !exist {
			// add
			valueDiff := engine.DiffMain("", curr.Value, true)
			diff := MapperNodeDiff{
				Node:      curr,
				Operate:   Add,
				Different: valueDiff,
			}
			Diffs.Nodes = append(Diffs.Nodes, diff)
		} else if  last.Father != curr.Father {
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

	var list2 []TreeNode
	for _, curr := range other.Tree {
		list2 = append(list2, curr)
	}
	sort.Slice(list2, func(i, j int) bool {
		v1,_ := strconv.Atoi(list2[i].Idx)
		v2,_ := strconv.Atoi(list2[j].Idx)
		return v1 < v2
	})
	for _, last := range list2 {
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