package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint
}

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint
}

type UpdatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
