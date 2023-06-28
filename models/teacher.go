package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name   string `json:"name"`
	Gender string `json:"gender"`
}
