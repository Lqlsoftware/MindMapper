package model

import (
	"github.com/Lqlsoftware/mindmapper/src/config"
	"github.com/Lqlsoftware/mindmapper/src/orm"
	"gopkg.in/mgo.v2/bson"
)

type Team struct {
	Id				int 	`json:"id"`
	Name			string 	`json:"name"`
	AdminId			int 	`json:"adminId"`
	MemberIds		[]int 	`json:"memberIds"`
	State			uint8 	`json:"state"`
	TeamTreeIds		[]int 	`json:"teamTreeIds"`
}

func GetTeamsInfoByMemberId(id int) ([]Team, error) {
	var team []Team
	err := orm.GetDatabase().C(config.TEAM_CNAME).Find(bson.M{"memberids": bson.M{"$in": []int{id}}}).All(&team)
	if err != nil {
		return nil, err
	}
	return team, nil
}