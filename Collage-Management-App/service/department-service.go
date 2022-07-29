package service

import (
	"Collage-Management-App/config"
	"Collage-Management-App/models"
)

func GetAllDepartments(dept *[]models.Department) (err error) {
	if err = config.DB.Find(dept).Error; err != nil {
		return err
	}

	return nil

}

func CreateDepartment(dept *models.Department) (err error) {

	if err = config.DB.Create(dept).Error; err != nil {
		return err

	}
	return nil
}

func GetDepartmentByID(dept *models.Department, department_id string) (err error) {
	if err = config.DB.Where("department_id=?", department_id).First(dept).Error; err != nil {
		return err
	}
	return nil
}

func UpadateDepartment(dept *models.Department, department_id string) (err error) {

	config.DB.Where("department_id=?", department_id).Save(dept)
	return nil
}

func DeleteDepartmentByID(dept *models.Department, department_id string) (err error) {

	config.DB.Delete(dept, department_id)

	return nil
}
