package main

type UserInfo struct {

	ID          uint           `gorm:"primaryKey"`
	UserName    string         `form:"username" json:"user_name" binding:"required"`
	PassWord    string         `form:"password" json:"pass_word" binding:"required"`
	Phone		string         `form:"phone" json:"phone" binding:"required"`
}
