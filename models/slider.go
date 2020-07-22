package models

import "github.com/jinzhu/gorm"

type Slider struct {
	gorm.Model
	id       int
	img      string
	status   Status `gorm:"foreignkey:statusID"`
	statusID int
}
