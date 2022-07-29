package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Passowrd string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Passowrd: "root",
		DBName:   "go-gin-restapi-mysql",
	}
	return &dbConfig
}

func DbURL(dbconfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbconfig.User,
		dbconfig.Passowrd,
		dbconfig.Host,
		dbconfig.Port,
		dbconfig.DBName,
	)

}
