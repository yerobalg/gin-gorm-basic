package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID uint
	Content string `gorm:"type:VARCHAR(255)"`
}