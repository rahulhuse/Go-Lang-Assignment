package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host      string
	Port      int
	User      string
	DBName    string
	Passoword string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:      "localhost",
		Port:      3306,
		User:      "root",
		Passoword: "root",
		DBName:    "collage-management-app",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Passoword,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

}
