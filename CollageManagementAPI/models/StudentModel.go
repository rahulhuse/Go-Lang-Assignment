package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	StudentID           uint   `json:"studentid"`
	StudentName         string `json:"studentname"`
	StudentMobileNumber string `json:"studentmobilenumber"`
	CollageID           uint   `json:"collageid"`
}

func (s *Student) TableName() string {
	return "student"
}
