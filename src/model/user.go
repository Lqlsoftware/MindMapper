package model

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

