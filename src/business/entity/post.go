package entity

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `json:"title" gorm:"type:varchar(255)"`
	Content string `json:"content" gorm:"type:text"`
	Comments []Comment
	Categories []Category `gorm:"many2many:post_categories;"`
}

type PostBody struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	CategoriesID []uint `json:"categoriesID"`
}

type PostParam struct {
	PostID int64 `uri:"post_id" gorm:"column:id"`
	PaginationParam
}
