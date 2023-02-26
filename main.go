package main

import (
	"gin-gorm-basic/sdk/config"
)

func main() {
	conf := config.Init()
	if err := conf.Load(".env"); err != nil {
		panic(err)
	}

	
}
