package models

type Teacher struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}
