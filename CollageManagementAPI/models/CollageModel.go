package models

import "gorm.io/gorm"

type Collage struct {
	gorm.Model
	CollageID      uint         `json:"collageid"`
	CollageName    string       `json:"collagename"`
	CollageEmail   string       `json:"collageemail"`
	CollageMobile  string       `json:"collagemobile"`
	CollageAddress string       `json:"collageaddress"`
	Departments    []Department `gorm:"foreignkey:collage_id"`
	Staffs         []Staff      `gorm:"foreignkey:collage_id"`
	Students       []Student    `gorm:"foreignkey:collage_id"`
}

func (c *Collage) TableName() string {
	return "collage"
}
