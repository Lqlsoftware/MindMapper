package main

import (
	"github.com/Lqlsoftware/mindmapper/src/orm"
	"github.com/astaxie/beego"
)

func init() {
	orm.InitDB()
	beego.SetStaticPath("/static","view")
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "gosessionid"
	bindRouter()
}

func main() {
	//base := git.Commit{
	//	Tree: Tree.MindMapperTree{Tree:map[string]Tree.TreeNode{
	//		"1":{"1","0",1,"1","text"},
	//		"2":{"2","1",1,"1","text"},
	//		"3":{"3","1",1,"2","text"},
	//	}},
	//}
	//masterHead := git.Commit{
	//	Tree: Tree.MindMapperTree{Tree:map[string]Tree.TreeNode{
	//		"1":{"1","0",1,"1","text"},
	//		"2":{"2","1",0,"1","text"},
	//		"3":{"3","1",1,"2","text"},
	//		"4":{"4","3",0,"1","text"},
	//	}},
	//}
	//branchHead := git.Commit{
	//	Tree: Tree.MindMapperTree{Tree:map[string]Tree.TreeNode{
	//		"1":{"1","0",1,"1","text"},
	//		"2":{"2","1",1,"1","text"},
	//		"3":{"3","1",1,"2","text"},
	//		"4":{"4","2",0,"1","text"},
	//		"5":{"5","3",0,"1","text"},
	//	}},
	//}
	//commit, conflict := git.Merge(&masterHead, &branchHead, &base)
	//fmt.Println(commit, conflict)
	//return
	beego.Run()
}