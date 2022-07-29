package main

import (
	"Collage-Management-App/config"
	"Collage-Management-App/models"
	"Collage-Management-App/routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {

	config.DB, err = gorm.Open("mysql",
		config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Collage{}, &models.Department{}, &models.Staff{}, &models.Student{})
	r := routes.SetupRouter()
	r.Run()
}
