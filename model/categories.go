package model

type categories struct {
	id       int    `gorm:"type:int;primary_key;unique;auto_increment"`
	category string `gorm:"type:varchar(255)"`
}
