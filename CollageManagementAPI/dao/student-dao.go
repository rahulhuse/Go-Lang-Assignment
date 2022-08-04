package dao

import (
	"gorm-test/database"
	"gorm-test/models"

	"gorm.io/gorm"
)

type StudentRepositoryInterface interface {
	GetAllStudents(stu *[]models.Student) (err error)
	CreateStudent(stu *models.Student) (err error)
	GetStudentByID(stu *models.Student, student_id string) (err error)
	UpadateStudent(stu *models.Student, student_id string) (err error)
	DeleteStudentByID(stu *models.Student, student_id string) (err error)
}

type studentRepo struct {
	Db *gorm.DB
}

func StudentNew() StudentRepositoryInterface {
	db := database.InitDb()
	return &studentRepo{Db: db}
}

func (stuRepo *studentRepo) GetAllStudents(stu *[]models.Student) (err error) {
	if err = stuRepo.Db.Find(stu).Error; err != nil {
		return err
	}

	return nil

}

func (stuRepo *studentRepo) CreateStudent(stu *models.Student) (err error) {

	if err = stuRepo.Db.Create(stu).Error; err != nil {
		return err

	}
	return nil
}

func (stuRepo *studentRepo) GetStudentByID(stu *models.Student, student_id string) (err error) {
	if err = stuRepo.Db.Where("student_id=?", student_id).First(stu).Error; err != nil {
		return err
	}
	return nil
}

func (stuRepo *studentRepo) UpadateStudent(stu *models.Student, student_id string) (err error) {

	stuRepo.Db.Where("student_id=?", student_id).Save(stu)
	return nil
}

func (stuRepo *studentRepo) DeleteStudentByID(stu *models.Student, student_id string) (err error) {

	stuRepo.Db.Delete(stu, student_id)

	return nil
}
