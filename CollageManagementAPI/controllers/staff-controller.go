package controllers

import (
	"fmt"
	"gorm-test/models"
	"gorm-test/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllStaff(c *gin.Context) {

	var staff []models.Staff

	err := service.GetAllStaff(&staff)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Staff: ": "Not found"})
	} else {
		c.JSON(http.StatusOK, staff)
	}

}

func CreateStaff(c *gin.Context) {

	fmt.Println("create staff controller")
	var staff models.Staff
	c.BindJSON(&staff)
	err := service.CreateStaff(&staff)

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

	err := service.GetStaffByID(&staff, staff_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"staff_id: " + staff_id: "Not found"})
	} else {
		c.JSON(http.StatusOK, staff)
	}
}

func UpadateStaff(c *gin.Context) {
	var staff models.Staff
	staff_id := c.Params.ByName("staff_id")

	err := service.GetStaffByID(&staff, staff_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"staff_id: " + staff_id: "Not found"})
	} else {
		c.BindJSON(&staff)
		service.UpadateStaff(&staff, staff_id)
		c.JSON(http.StatusOK, staff)
	}

}

func DeleteStaffByID(c *gin.Context) {

	fmt.Println("In controller deletebyStaffId")

	var staff models.Staff
	staff_id := c.Params.ByName("staff_id")

	err := service.GetStaffByID(&staff, staff_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"staff_id: " + staff_id: "Not found"})

	} else {
		service.DeleteStaffByID(&staff, staff_id)
		c.JSON(http.StatusOK, gin.H{"staff_id " + staff_id: "is deleted"})
	}

}
