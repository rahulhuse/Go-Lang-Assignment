package service

import (
	"gorm-test/dao"
	"gorm-test/models"
)

func GetAllDepartments(dept *[]models.Department) (err error) {

	err = dao.DepartmentNew().GetAllDepartments(dept)

	if err != nil {
		return err
	}

	return nil

}

func CreateDepartment(dept *models.Department) (err error) {

	err = dao.DepartmentNew().CreateDepartment(dept)

	if err != nil {
		return err
	}
	return nil
}

func GetDepartmentByID(dept *models.Department, department_id string) (err error) {

	err = dao.DepartmentNew().GetDepartmentByID(dept, department_id)
	if err != nil {
		return err
	}
	return nil
}

func UpadateDepartment(dept *models.Department, department_id string) (err error) {

	err = dao.DepartmentNew().UpadateDepartment(dept, department_id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDepartmentByID(dept *models.Department, department_id string) (err error) {

	err = dao.DepartmentNew().DeleteDepartmentByID(dept, department_id)
	if err != nil {
		return err
	}
	return nil
}
