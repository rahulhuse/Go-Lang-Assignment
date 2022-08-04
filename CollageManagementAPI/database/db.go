package database

import (
	"fmt"
	"gorm-test/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "root"
const DB_NAME = "collage-management-app"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = ConnectDB()
	return Db
}

func ConnectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.Collage{}, &models.Department{}, &models.Staff{}, &models.Student{}, &models.User{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}
   

	return db
}

func CloseDbConnextion(Db *gorm.DB) {
	dbSQL, err := Db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
