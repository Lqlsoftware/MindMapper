package model

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
type MapperNode struct {
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
type MindMapper struct {
	Tree	map[string]MapperNode
	Hash	string
}

func (mindMapper *MindMapper)ToJson() (string,error) {
	res,err := json.Marshal(*mindMapper)
	return string(res),err
}

func (mindMapper *MindMapper)DiffWith(other *MindMapper) MapperDiff {
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