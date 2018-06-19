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
	//	Tree: Tree.MindMapperTree{Tree: map[string]Tree.TreeNode{
	//		"0": {"0", "-1",2,1,-1,""},
	//		"1": {"1","0",0,1,-1,""},
	//		"2": {"2","0",1,2,1,""},
	//		"4": {"4","2",1,1,-1,"21"},
	//		"5": {"5","4",0,1,-1,"5"},
	//	}},
	//}
	//master := git.Commit{
	//	Tree: Tree.MindMapperTree{Tree: map[string]Tree.TreeNode{
	//		"0": {"0", "-1",2,1,-1,""},
	//		"1": {"1","0",0,1,-1,""},
	//		"2": {"2","0",1,2,1,""},
	//		"4": {"4","2",1,1,-1,"21"},
	//		"5": {"5","4",0,1,-1,"5"},
	//	}},
	//}
	//head := git.Commit{
	//	Tree: Tree.MindMapperTree{Tree: map[string]Tree.TreeNode{
	//		"0": {"0", "-1",2,1,-1,""},
	//		"1": {"1","0",0,1,-1,""},
	//		"2": {"2","0",1,2,1,""},
	//	}},
	//}
	//commit, conflict := git.Merge(&master, &head, &base, model.User{})
	//fmt.Println(commit, "\n", conflict)
	//return
	beego.Run()
}