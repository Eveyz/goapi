package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
}

type CreateStudentInput struct {
	Name   string `json:"title" binding:"required"`
	Age    string `json:"age" binding:"required"`
	Gender string `json:"gender" binding:"required"`
}

type UpdateStudentInput struct {
	Name   string `json:"name"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
}
