package models

import "github.com/jinzhu/gorm"

type Status struct {
	gorm.Model
	Name string `gorm:"type:varchar(20)"`
}
