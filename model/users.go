package model

type User struct {
	id       int    `gorm:"type:int;primary_key;unique;auto_increment"`
	username string `gorm:"type:varchar(255)"`
	password string `gorm:"type:varchar(255)"`
}
