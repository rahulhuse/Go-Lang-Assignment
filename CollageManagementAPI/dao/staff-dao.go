package dao

import (
	"gorm-test/database"
	"gorm-test/models"

	"gorm.io/gorm"
)

type StaffService interface {
	GetAllStaff(staff *[]models.Staff) (err error)
	CreateStaff(staff *models.Staff) (err error)
	GetStaffByID(staff *models.Staff, staff_id string) (err error)
	UpadateStaff(staff *models.Staff, staff_id string) (err error)
	DeleteStaffByID(staff *models.Staff, staff_id string) (err error)
}

type stafftRepo struct {
	Db *gorm.DB
}

func StaffNew() StaffService {
	db := database.InitDb()
	db.AutoMigrate(&models.Staff{})
	return &stafftRepo{Db: db}
}

func (staffRepo *stafftRepo) GetAllStaff(staff *[]models.Staff) (err error) {
	if err = staffRepo.Db.Find(staff).Error; err != nil {
		return err
	}

	return nil

}

func (staffRepo *stafftRepo) CreateStaff(staff *models.Staff) (err error) {

	if err = staffRepo.Db.Create(staff).Error; err != nil {
		return err

	}
	return nil
}

func (staffRepo *stafftRepo) GetStaffByID(staff *models.Staff, staff_id string) (err error) {
	if err = staffRepo.Db.Where("staff_id=?", staff_id).First(staff).Error; err != nil {
		return err
	}
	return nil
}

func (staffRepo *stafftRepo) UpadateStaff(staff *models.Staff, staff_id string) (err error) {

	staffRepo.Db.Where("staff_id=?", staff_id).Save(staff)
	return nil
}

func (staffRepo *stafftRepo) DeleteStaffByID(staff *models.Staff, staff_id string) (err error) {

	staffRepo.Db.Delete(staff, staff_id)

	return nil
}
