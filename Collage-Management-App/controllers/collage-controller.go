package controllers

import (
	"Collage-Management-App/models"
	"Collage-Management-App/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCollages(c *gin.Context) {
	var collage []models.Collage
	err := service.GetAllCollages(&collage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Collages: ": "Not found"})
	} else {
		c.JSON(http.StatusOK, collage)
	}
}

func CreateCollage(c *gin.Context) {
	var collage []models.Collage
	c.BindJSON(&collage)
	err := service.CreateCollage(&collage)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Collage": "Is not created"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "Collage has been created"})
	}
}

func GetCollageByID(c *gin.Context) {
	collageid := c.Params.ByName("collageid")
	var collage models.Collage

	err := service.GetCollageByID(&collage, collageid)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"CollageID: " + collageid: "Not found"})
	} else {
		c.JSON(http.StatusOK, collage)
	}

}

func UpdateCollage(c *gin.Context) {

	var collage models.Collage
	collageid := c.Params.ByName("collageid")

	err := service.GetCollageByID(&collage, collageid)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"CollageID: " + collageid: "Not found"})
	} else {
		c.BindJSON(&collage)
		err = service.UpdateCollage(&collage, collageid)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"CollageID: " + collageid: "Not found"})
		}
		c.JSON(http.StatusOK, collage)
	}
}

func DeleteCollage(c *gin.Context) {
	var collage models.Collage
	collage_id := c.Params.ByName("collage_id")

	err := service.GetCollageByID(&collage, collage_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"CollageID: " + collage_id: "Not found"})
	} else {

		_ = service.DeleteCollage(&collage, collage_id)

		c.JSON(http.StatusOK, gin.H{"collage_id " + collage_id: "is deleted"})

	}
}
