package model

import (
	"errors"

	"github.com/Lqlsoftware/mindmapper/src/config"
	"github.com/Lqlsoftware/mindmapper/src/orm"
	"github.com/astaxie/beego"
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

func GetUser(this interface{}) (User, error) {
	user := this.(*beego.Controller).GetSession("user")
	if user == nil {
		return User{}, errors.New("not login")
	} else {
		return user.(User), nil
	}
}