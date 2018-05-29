package main

import (
	"log"

	"./config"
	"./model"
	"./model/Tree"
	"./model/git"
	"gopkg.in/mgo.v2"
)

func main() {
	// Connect mongodb
	session, err := mgo.Dial(config.DB_URL)
	if err != nil {
		panic(err)
	}
	db := session.DB(config.DB_NAME)

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

	// create collection team
	c = db.C(config.TEAM_CNAME)
	err = c.Insert(&model.Team{
		Id:			0,
		Name:		"admin",
		AdminId:	0,
		MemberIds:	[]int{},
		State:		1,
		TeamTreeIds:[]int{},
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