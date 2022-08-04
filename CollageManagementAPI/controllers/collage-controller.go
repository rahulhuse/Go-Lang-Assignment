package controllers

import (
	"fmt"
	"gorm-test/dao"
	"gorm-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCollagesCon(c *gin.Context) {

	fmt.Println("In controller start")
	var collage []models.Collage

	err := dao.CollageNew().GetAllCollages(&collage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Collages: ": "Not found"})
	} else {
		c.JSON(http.StatusOK, collage)
	}

	fmt.Println("In controller end")
}

func MyTest(c *gin.Context) {

	var err error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Collages: ": "Not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"Collages: ": "found"})
	}

}

func CreateCollageCon(c *gin.Context) {
	var collage []models.Collage
	errs := c.BindJSON(&collage)

	if errs != nil {
		return
	}

	err := dao.CollageNew().CreateCollage(&collage)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Collage": "Is not created"})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "Collage has been created"})
		return
	}
}

func GetCollageByIDCon(c *gin.Context) {
	collageid := c.Params.ByName("collageid")
	var collage models.Collage

	err := dao.CollageNew().GetCollageByID(&collage, collageid)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"CollageID: " + collageid: "Not found"})
	} else {
		c.JSON(http.StatusOK, collage)
	}

}

func UpdateCollageCon(c *gin.Context) {

	var collage models.Collage
	collageid := c.Params.ByName("collageid")

	err := dao.CollageNew().GetCollageByID(&collage, collageid)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"CollageID: " + collageid: "Not found"})
	} else {
		c.BindJSON(&collage)
		err = dao.CollageNew().UpdateCollage(&collage, collageid)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"CollageID: " + collageid: "Not found"})
		}
		c.JSON(http.StatusOK, collage)
	}
}

func DeleteCollageCon(c *gin.Context) {
	var collage models.Collage
	collageid := c.Params.ByName("collageid")

	err := dao.CollageNew().GetCollageByID(&collage, collageid)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"CollageID: " + collageid: "Not found"})
	} else {

		dao.CollageNew().DeleteCollage(&collage, collageid)

		c.JSON(http.StatusOK, gin.H{"collage_id " + collageid: "is deleted"})

	}
}
