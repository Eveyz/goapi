package models

type Class struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Teacher string `json:"teacher_id" gorm:"foreign_key"`
}
