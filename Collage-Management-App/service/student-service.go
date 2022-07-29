package service

import (
	"Collage-Management-App/config"
	"Collage-Management-App/models"
)

func GetAllStudents(stu *[]models.Student) (err error) {
	if err = config.DB.Find(stu).Error; err != nil {
		return err
	}

	return nil

}

func CreateStudent(stu *models.Student) (err error) {

	if err = config.DB.Create(stu).Error; err != nil {
		return err

	}
	return nil
}

func GetStudentByID(stu *models.Student, student_id string) (err error) {
	if err = config.DB.Where("student_id=?", student_id).First(stu).Error; err != nil {
		return err
	}
	return nil
}

func UpadateStudent(stu *models.Student, student_id string) (err error) {

	config.DB.Where("student_id=?", student_id).Save(stu)
	return nil
}

func DeleteStudentByID(stu *models.Student, student_id string) (err error) {

	config.DB.Delete(stu, student_id)

	return nil
}
