package model

type Task struct {
	id   int    `gorm:"type:int;primary_key;unique;auto_increment"`
	task string `gorm:"type:varchar(255)"`
	user User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
