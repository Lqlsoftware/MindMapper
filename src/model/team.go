package model

type Team struct {
	Id				int 	`json:"id"`
	Name			string 	`json:"name"`
	AdminId			int 	`json:"adminId"`
	MemberIds		[]int 	`json:"memberIds"`
	State			uint8 	`json:"state"`
	TeamTreeIds		[]int 	`json:"teamTreeIds"`
}
