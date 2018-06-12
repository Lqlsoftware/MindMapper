package main

import (
	"github.com/Lqlsoftware/mindmapper/src/orm"
	"github.com/astaxie/beego"
)

func init() {
	orm.InitDB()
	beego.SetStaticPath("/","/view")
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "gosessionid"
	bindRouter()
}

func main() {
	//dstCommit, _ := git.LoadCommit(1)
	//commit := git.Commit{
	//	Id:			git.GetLastCommitId(),
	//	Diff:		Tree.MapperDiff{},
	//	Time:		time.Now().Unix(),
	//	Title:		"third commit",
	//	Summary:	"hhhhh",
	//	Tree:		Tree.MindMapperTree{
	//		Tree: 	map[string]Tree.TreeNode{
	//			"0": {"0","-1", "1","root"},
	//			"1": {"1","0", "1","child1"},
	//			"2": {"2","0", "2","asdasd"},
	//			"3": {"3","1", "1","gg"},
	//			"4": {"4","2", "1","errer"},
	//		},
	//		Hash: 	"777",
	//	},
	//}
	//fmt.Println(dstCommit.MergeWith(&commit))
	//
	//return
	beego.Run()
}