package main

import (
	"log"

	"github.com/Lqlsoftware/mindmapper/src/config"
	"github.com/Lqlsoftware/mindmapper/src/model"
	"github.com/Lqlsoftware/mindmapper/src/model/Tree"
	"github.com/Lqlsoftware/mindmapper/src/model/git"
	"gopkg.in/mgo.v2"
)

func main() {
	// Connect mongodb
	session, err := mgo.Dial(config.DB_URL)
	if err != nil {
		panic(err)
	}
	db := session.DB(config.DB_NAME)

	// drop database
	db.DropDatabase()

	// create collection user
	c := db.C(config.USER_CNAME)
	err = c.Insert(&model.User{
		Id:			0,
		Username:	"admin",
		Password:	"admin",
		State:		1,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = c.Insert(&model.User{
		Id:			1,
		Username:	"alice",
		Password:	"123",
		State:		1,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = c.Insert(&model.User{
		Id:			2,
		Username:	"bob",
		Password:	"123",
		State:		1,
	})
	if err != nil {
		log.Fatal(err)
	}

	// create collection branchSet
	c = db.C(config.BRANCHSET_CNAME)
	err = c.Insert(&git.BranchSet{
		TreeId:			0,
		MainBranchId:	0,
		BranchIds:		[]int{},
	})
	if err != nil {
		log.Fatal(err)
	}

	// create collection branch
	c = db.C(config.BRANCH_CNAME)
	err = c.Insert(&git.Branch{
		Id:			0,
		Name:		"origin",
		HeadId:		0,
		CommitIds:	[]int{},
	})
	if err != nil {
		log.Fatal(err)
	}

	// create collection commit
	c = db.C(config.COMMIT_CNAME)
	err = c.Insert(&git.Commit{
		Id:			0,
		Diff:		Tree.MapperDiff{},
		Time:		0,
		Title:		"initial commit",
		Summary:	"init",
		Tree:		Tree.MindMapperTree{},
	})
	if err != nil {
		log.Fatal(err)
	}
}