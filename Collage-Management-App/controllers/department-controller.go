package controllers

import (
	"Collage-Management-App/models"
	"Collage-Management-App/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDepartments(c *gin.Context) {

	var dept []models.Department

	err := service.GetAllDepartments(&dept)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Departments: ": "Not found"})
	} else {
		c.JSON(http.StatusOK, dept)
	}

}

func CreateDepartment(c *gin.Context) {
	var dept models.Department
	c.BindJSON(&dept)
	err := service.CreateDepartment(&dept)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Deparments": "Is not created"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "New department created"})
	}
}

func GetDepartmentByID(c *gin.Context) {
	department_id := c.Params.ByName("department_id")
	var dept models.Department

	err := service.GetDepartmentByID(&dept, department_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"DepartmentID: " + department_id: "Not found"})
	} else {
		c.JSON(http.StatusOK, dept)
	}
}

func UpadateDepartment(c *gin.Context) {
	var dept models.Department
	department_id := c.Params.ByName("department_id")

	err := service.GetDepartmentByID(&dept, department_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"DepartmentID: " + department_id: "Not found"})
	} else {
		c.BindJSON(&dept)
		service.UpadateDepartment(&dept, department_id)
		c.JSON(http.StatusOK, dept)
	}

}

func DeleteDepartmentByID(c *gin.Context) {
	var dept models.Department
	department_id := c.Params.ByName("department_id")

	err := service.GetDepartmentByID(&dept, department_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"DepartmentID: " + department_id: "Not found"})

	} else {
		service.DeleteDepartmentByID(&dept, department_id)
		c.JSON(http.StatusOK, gin.H{"department_id " + department_id: "is deleted"})
	}

}
