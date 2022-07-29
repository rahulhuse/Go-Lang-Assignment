package main

import (
	"fmt"
	"golang-workspace/src/go-gin-restapi-mysql/Config"
	"golang-workspace/src/go-gin-restapi-mysql/Models"
	"golang-workspace/src/go-gin-restapi-mysql/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})

	r := Routes.SetupRouter()
	r.Run()

}
