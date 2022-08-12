package service

import (
	"fmt"
	"gorm-test/dao"
	"gorm-test/models"
)

func GetAllStaff(staff *[]models.Staff) (err error) {

	err = dao.StaffNew().GetAllStaff(staff)

	if err != nil {
		return err
	}

	return nil

}

func CreateStaff(staff *models.Staff) (err error) {

	fmt.Println("create staff service")
	err = dao.StaffNew().CreateStaff(staff)

	if err != nil {
		return err
	}

	return nil
}

func GetStaffByID(staff *models.Staff, staff_id string) (err error) {

	err = dao.StaffNew().GetStaffByID(staff, staff_id)
	if err != nil {
		return err
	}

	return nil
}

func UpadateStaff(staff *models.Staff, staff_id string) (err error) {

	err = dao.StaffNew().UpadateStaff(staff, staff_id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteStaffByID(staff *models.Staff, staff_id string) (err error) {

	err = dao.StaffNew().DeleteStaffByID(staff, staff_id)
	if err != nil {
		return err
	}

	return nil
}
