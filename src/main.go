package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}




func main() {
	beego.Router("/login", &LoginController{})
	beego.Run()
	//t1 := model.MindMapper{
	//	Tree:	make(map[string]model.MapperNode),
	//	Hash:	"",
	//}
	//t2 := model.MindMapper{
	//	Tree:	make(map[string]model.MapperNode),
	//	Hash:	"",
	//}
	//t1.Tree["1"] = model.MapperNode{"1","0","1","fuck"}
	//t1.Tree["2"] = model.MapperNode{"2","1","1","git"}
	//
	//t2.Tree["1"] = model.MapperNode{"1","0","1","fucka"}
	//t2.Tree["3"] = model.MapperNode{"3","1","1","gitSS"}
	//diff := t1.DiffWith(&t2)
	//for _,v := range diff.Nodes {
	//	fmt.Println(v.Operate)
	//	fmt.Println(diffmatchpatch.New().DiffPrettyText(v.Different))
	//}
	//
	//fmt.Println(t1.ToJson())
	//
	//
	//fmt.Println(diff.ToJson())
}