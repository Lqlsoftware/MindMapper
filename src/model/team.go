package model

type Team struct {
	Id				int
	Name			string
	AdminId			int
	MemberIds		[]int
	State			uint8
	TeamTreeIds		[]int
}
