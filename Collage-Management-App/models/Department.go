package models

import "github.com/jinzhu/gorm"

type Department struct {
	gorm.Model
	DepartmentID   uint   `json:"departmentid"`
	DepartmentName string `json:"departmentname"`
	CollageID      uint   `json:"collageid"`

}

func (d *Department) TableName() string {
	return "department"
}
