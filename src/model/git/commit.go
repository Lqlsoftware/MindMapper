package git

import (
	"github.com/Lqlsoftware/mindmapper/src/model/Tree"
)

type Commit struct {
	Id			int
	Diff		Tree.MapperDiff
	Time		uint32
	Title		string
	Summary		string
	Tree		Tree.MindMapperTree
}
