package controllers

import (
	"fmt"
	"gorm-test/dao"
	"gorm-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllStaff(c *gin.Context) {

	var staff []models.Staff

	err := dao.StaffNew().GetAllStaff(&staff)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Staff: ": "Not found"})
	} else {
		c.JSON(http.StatusOK, staff)
	}

}

func CreateStaff(c *gin.Context) {
	var staff models.Staff
	c.BindJSON(&staff)
	err := dao.StaffNew().CreateStaff(&staff)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Staff": "Is not created"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "New Staff created"})
	}
}

func GetStaffByID(c *gin.Context) {
	staff_id := c.Params.ByName("staff_id")
	var staff models.Staff

	err := dao.StaffNew().GetStaffByID(&staff, staff_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"staff_id: " + staff_id: "Not found"})
	} else {
		c.JSON(http.StatusOK, staff)
	}
}

func UpadateStaff(c *gin.Context) {
	var staff models.Staff
	staff_id := c.Params.ByName("staff_id")

	err := dao.StaffNew().GetStaffByID(&staff, staff_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"staff_id: " + staff_id: "Not found"})
	} else {
		c.BindJSON(&staff)
		dao.StaffNew().UpadateStaff(&staff, staff_id)
		c.JSON(http.StatusOK, staff)
	}

}

func DeleteStaffByID(c *gin.Context) {
	var staff models.Staff
	staff_id := c.Params.ByName("staff_id")

	err := dao.StaffNew().GetStaffByID(&staff, staff_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"staff_id: " + staff_id: "Not found"})

	} else {
		dao.StaffNew().DeleteStaffByID(&staff, staff_id)
		c.JSON(http.StatusOK, gin.H{"staff_id " + staff_id: "is deleted"})
	}

}
