package service

import (
	"gorm-test/dao"
	"gorm-test/models"
)

func GetAllCollages(collage *[]models.Collage) (err error) {

	dao.CollageNew().GetAllCollages(collage)

	if err != nil {
		return err
	}
	return nil
}

func GetCollageByID(collage *models.Collage, collageid string) (err error) {

	err = dao.CollageNew().GetCollageByID(collage, collageid)

	if err != nil {
		return err
	}

	return nil
}

func CreateCollage(collage *[]models.Collage) (err error) {

	err = dao.CollageNew().CreateCollage(collage)

	if err != nil {

		return err
	}

	return nil
}

func UpdateCollage(collage *models.Collage, collageid string) (err error) {

	err = dao.CollageNew().UpdateCollage(collage, collageid)

	if err != nil {

		return err
	}

	return nil
}

func DeleteCollage(collage *models.Collage, collageid string) (err error) {

	dao.CollageNew().DeleteCollage(collage, collageid)

	if err != nil {

		return err
	}

	return nil

}
