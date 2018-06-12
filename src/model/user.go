package model

import (
	"github.com/Lqlsoftware/mindmapper/src/config"
	"github.com/Lqlsoftware/mindmapper/src/orm"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	// 用户ID
	Id 			int 	`json:"id"`
	// 用户名
	Username 	string	`json:"username"`
	// 密码
	Password 	string	`json:"password"`
	// 状态
	State		uint8	`json:"state"`
}

func VaildUser(username, password string) (User, error) {
	// 获取用户信息
	user := User{}
	err := orm.GetDatabase().C(config.USER_CNAME).Find(bson.M{"username": username, "password": password}).One(&user)
	return user, err
}