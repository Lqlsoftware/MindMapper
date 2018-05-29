package model

import (
	"github.com/Lqlsoftware/mindmapper/src/config"
	"github.com/Lqlsoftware/mindmapper/src/orm"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	// 用户ID
	Id 			int
	// 用户名
	Username 	string
	// 密码
	Password 	string
	// 状态
	State		uint8
}

func VaildUser(username, password string) (User, error) {

	// 获取用户信息
	user := User{}
	err := orm.GetDatabase().C(config.USER_CNAME).Find(bson.M{"Username": username, "Password": password}).One(&user)

	return user, err
}