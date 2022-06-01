package model

type User struct {
	Base
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}
