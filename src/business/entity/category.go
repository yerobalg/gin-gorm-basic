package entity

type Category struct {
	ID   uint
	Name string `gorm:"type:VARCHAR(30)"`
}