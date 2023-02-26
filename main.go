package main

import (
	"fmt"
	"gin-gorm-basic/sdk/config"
	"gin-gorm-basic/sdk/sql"
	"gin-gorm-basic/src/business/entity"
	"gin-gorm-basic/src/handler"

	"gorm.io/gorm"
)

func main() {
	conf := config.Init()
	if err := conf.Load(".env"); err != nil {
		panic(err)
	}

	sqlConfig := sql.Config{
		Username: conf.Get("DB_USERNAME"),
		Password: conf.Get("DB_PASSWORD"),
		Host:     conf.Get("DB_HOST"),
		Port:     conf.Get("DB_PORT"),
		Database: conf.Get("DB_DATABASE"),
	}

	sql, err := sql.Init(sqlConfig)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected successfully!")

	db := sql.GetInstance()
	db.AutoMigrate(entity.Post{}, &entity.Comment{}, &entity.Category{})

	if err := seedCategory(db); err != nil {
		panic("GAGAL SEED BROOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	}

	handler := handler.Init(conf, db)
	handler.Run()
}


func seedCategory(sql *gorm.DB) error{
	var categories []entity.Category

	if err := sql.First(&categories).Error; err != gorm.ErrRecordNotFound {
		return err
	}
	categories = []entity.Category{
		{
			Name: entity.BusinessCategory,
		},
		{
			Name: entity.TravelCategory,
		},
		{
			Name: entity.HealthCategory,
		},
		{
			Name: entity.TechnologyCategory,
		},
	}

	if err := sql.Create(&categories).Error; err != nil {
		return err
	}
	return nil
}