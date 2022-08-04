package controllers

import (
	"fmt"
	"gorm-test/dao"
	"gorm-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {

	var stu []models.Student

	err := dao.StudentNew().GetAllStudents(&stu)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Student: ": "Not found"})
	} else {
		c.JSON(http.StatusOK, stu)
	}

}

func CreateStudent(c *gin.Context) {
	var stu models.Student
	c.BindJSON(&stu)
	err := dao.StudentNew().CreateStudent(&stu)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Student": "Is not created"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "New Student created"})
	}
}

func GetStudentByID(c *gin.Context) {
	student_id := c.Params.ByName("student_id")
	var stu models.Student

	err := dao.StudentNew().GetStudentByID(&stu, student_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"student_id: " + student_id: "Not found"})
	} else {
		c.JSON(http.StatusOK, stu)
	}
}

func UpadateStudent(c *gin.Context) {

	var stu models.Student
	student_id := c.Params.ByName("student_id")

	err := dao.StudentNew().GetStudentByID(&stu, student_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"student_id: " + student_id: "Not found"})
	} else {
		c.BindJSON(&stu)
		dao.StudentNew().UpadateStudent(&stu, student_id)
		c.JSON(http.StatusOK, stu)
	}

}

func DeleteStudentByID(c *gin.Context) {
	var stu models.Student
	student_id := c.Params.ByName("student_id")

	err := dao.StudentNew().GetStudentByID(&stu, student_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"student_id: " + student_id: "Not found"})

	} else {
		dao.StudentNew().DeleteStudentByID(&stu, student_id)
		c.JSON(http.StatusOK, gin.H{"student_id " + student_id: "is deleted"})
	}

}
