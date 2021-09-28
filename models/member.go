package models

type Member struct {
	Id int
	Username string
	Nickname string
	License string
}

func (Member) TableName() string  {
	return "p_member"
}