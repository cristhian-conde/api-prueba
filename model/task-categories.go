package model

type TaskCategories struct {
	id         int        `gorm:"type:int;primary_key;unique;auto_increment"`
	task       Task       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	categories categories `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
