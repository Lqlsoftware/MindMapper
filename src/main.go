package main

import (
	"github.com/Lqlsoftware/mindmapper/src/orm"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	orm.InitDB()
	bindRouter()
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
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

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "gosessionid"
	beego.Run()
}