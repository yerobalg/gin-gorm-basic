package main

import (
	"fmt"
	"gin-gorm-basic/sdk/config"
	"gin-gorm-basic/sdk/sql"
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

	_, err := sql.Init(sqlConfig)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected successfully!")

}
