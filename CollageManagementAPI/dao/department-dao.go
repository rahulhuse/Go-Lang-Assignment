package dao

import (
	"gorm-test/database"
	"gorm-test/models"

	"gorm.io/gorm"
)

type DepartmentService interface {
	GetAllDepartments(dept *[]models.Department) (err error)
	GetDepartmentByID(dept *models.Department, department_id string) (err error)
	CreateDepartment(dept *models.Department) (err error)
	UpadateDepartment(dept *models.Department, department_id string) (err error)
	DeleteDepartmentByID(dept *models.Department, department_id string) (err error)
}

type departmentRepo struct {
	Db *gorm.DB
}

func DepartmentNew() DepartmentService {
	db := database.InitDb()
	db.AutoMigrate(&models.Department{})
	return &departmentRepo{Db: db}
}

func (deptRepo *departmentRepo) GetAllDepartments(dept *[]models.Department) (err error) {
	if err = deptRepo.Db.Find(dept).Error; err != nil {
		return err
	}

	return nil

}

func (deptRepo *departmentRepo) CreateDepartment(dept *models.Department) (err error) {

	if err = deptRepo.Db.Create(dept).Error; err != nil {
		return err

	}
	return nil
}

func (deptRepo *departmentRepo) GetDepartmentByID(dept *models.Department, department_id string) (err error) {
	if err = deptRepo.Db.Where("department_id=?", department_id).First(dept).Error; err != nil {
		return err
	}
	return nil
}

func (deptRepo *departmentRepo) UpadateDepartment(dept *models.Department, department_id string) (err error) {

	deptRepo.Db.Where("department_id=?", department_id).Save(dept)
	return nil
}

func (deptRepo *departmentRepo) DeleteDepartmentByID(dept *models.Department, department_id string) (err error) {

	deptRepo.Db.Delete(dept, department_id)

	return nil
}
