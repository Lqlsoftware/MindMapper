package Tree

import (
	"encoding/json"

	"github.com/sergi/go-diff/diffmatchpatch"
)

type MapperNodeOperate uint8

const (
	Add			MapperNodeOperate	= 0x01
	Delete		MapperNodeOperate	= 0x02
	Modify		MapperNodeOperate 	= 0x03
)

func (operate MapperNodeOperate)String() string {
	switch operate {
	case Add:
		return "+ ADD"
	case Delete:
		return "- DELETE"
	case Modify:
		return "~ MODIFY"
	}
	return ""
}

/*
	MindMapper Node Different

	NodeIdx 	: *MapperNode			变更节点
	Operate		: MapperNodeOperate		操作
	Different	: []Diff				字符差异
 */
type MapperNodeDiff struct {
	Node	 	TreeNode	 			`json:"node"`
	Operate		MapperNodeOperate	 	`json:"operate"`
	Different	[]diffmatchpatch.Diff 	`json:"different"`
}

type MapperDiff struct {
	Nodes		[]MapperNodeDiff 		`json:"nodes"`
}

func (diffs *MapperDiff)ToJson() (string,error) {
	res,err := json.Marshal(*diffs)
	return string(res), err
}
