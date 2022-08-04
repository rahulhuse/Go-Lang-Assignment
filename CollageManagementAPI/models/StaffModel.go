package models

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	StaffID           uint   `json:"staffid"`
	StaffName         string `json:"staffname"`
	StaffMobileNumber string `json:"staffmobilenumber"`
	CollageID         uint   `json:"collageid"`
}

func (s *Staff) TableName() string {
	return "staff"
}
