package service

import (
	"Collage-Management-App/config"
	"Collage-Management-App/models"
)

func GetAllStaff(staff *[]models.Staff) (err error) {
	if err = config.DB.Find(staff).Error; err != nil {
		return err
	}

	return nil

}

func CreateStaff(staff *models.Staff) (err error) {

	if err = config.DB.Create(staff).Error; err != nil {
		return err

	}
	return nil
}

func GetStaffByID(staff *models.Staff, staff_id string) (err error) {
	if err = config.DB.Where("staff_id=?", staff_id).First(staff).Error; err != nil {
		return err
	}
	return nil
}

func UpadateStaff(staff *models.Staff, staff_id string) (err error) {

	config.DB.Where("staff_id=?", staff_id).Save(staff)
	return nil
}

func DeleteStaffByID(staff *models.Staff, staff_id string) (err error) {

	config.DB.Delete(staff, staff_id)

	return nil
}
