package dao

import (
	"fmt"
	"gorm-test/database"
	"gorm-test/models"

	"gorm.io/gorm"
)

type CollageService interface {
	GetAllCollages(collage *[]models.Collage) (err error)
	GetCollageByID(collage *models.Collage, collageid string) (err error)
	CreateCollage(collage *[]models.Collage) (err error)
	UpdateCollage(collage *models.Collage, collageid string) (err error)
	DeleteCollage(collage *models.Collage, collage_id string) (err error)
}

type CollageRepo struct {
	Db *gorm.DB
}

func CollageNew() CollageService {
	db := database.InitDb()
	//db.AutoMigrate(&models.Collage{})
	return &CollageRepo{Db: db}
}

func (collageRepo *CollageRepo) GetAllCollages(collage *[]models.Collage) (err error) {

	if err = collageRepo.Db.Find(collage).Error; err != nil {
		return err

	}

	return nil
}

func (collageRepo *CollageRepo) GetCollageByID(collage *models.Collage, collageid string) (err error) {

	if err = collageRepo.Db.Preload("Departments").Preload("Staffs").Preload("Students").Where("collage_id = ?", collageid).First(collage).Error; err != nil {
		return err

	}

	return nil
}

func (collageRepo *CollageRepo) CreateCollage(collage *[]models.Collage) (err error) {

	// if err = config.DB.Create(collage).Error; err != nil {
	// 	return err
	// }

	for _, collages := range *collage {

		collageRepo.Db.Create(&collages)
	}

	if err != nil {

		return err
	}

	return nil
}

func (collageRepo *CollageRepo) UpdateCollage(collage *models.Collage, collageid string) (err error) {

	fmt.Println(collage)
	collageRepo.Db.Save(collage)

	return nil
}

func (collageRepo *CollageRepo) DeleteCollage(collage *models.Collage, collage_id string) (err error) {
	collageRepo.Db.Where("collage_id = ?", collage_id).Delete(collage)

	return nil

}
