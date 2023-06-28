package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name    string `json:"name"`
	Teacher string `json:"teacher_id" gorm:"foreign_key"`
}
