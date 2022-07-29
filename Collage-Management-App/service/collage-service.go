package service

import (
	"Collage-Management-App/models"
	"fmt"

	"Collage-Management-App/config"

	_ "github.com/go-sql-driver/mysql"
)

//collage management

func GetAllCollages(collage *[]models.Collage) (err error) {

	if err = config.DB.Find(collage).Error; err != nil {
		return err

	}

	return nil
}

func GetCollageByID(collage *models.Collage, collageid string) (err error) {

	if err = config.DB.Preload("Departments").Preload("Staffs").Preload("Students").Where("collage_id = ?", collageid).First(collage).Error; err != nil {
		return err

	}

	return nil
}

func CreateCollage(collage *[]models.Collage) (err error) {

	// if err = config.DB.Create(collage).Error; err != nil {
	// 	return err
	// }

	if err != nil {

		return err
	}

	for _, collages := range *collage {

		config.DB.Create(collages)
	}

	return nil
}

func UpdateCollage(collage *models.Collage, collageid string) (err error) {

	fmt.Println(collage)
	config.DB.Save(collage)

	return nil
}

func DeleteCollage(collage *models.Collage, collage_id string) (err error) {
	if err = config.DB.Delete(collage, collage_id).Error; err != nil {
		return err
	}
	return nil

}
