package service

import (
	"gorm-test/dao"
	"gorm-test/models"
)

func GetAllStudents(stu *[]models.Student) (err error) {

	err = dao.StudentNew().GetAllStudents(stu)
	if err != nil {
		return err
	}

	return nil

}

func CreateStudent(stu *models.Student) (err error) {

	err = dao.StudentNew().CreateStudent(stu)
	if err != nil {
		return err

	}
	return nil
}

func GetStudentByID(stu *models.Student, student_id string) (err error) {

	err = dao.StudentNew().GetStudentByID(stu, student_id)
	if err != nil {
		return err
	}
	return nil
}

func UpadateStudent(stu *models.Student, student_id string) (err error) {

	err = dao.StudentNew().UpadateStudent(stu, student_id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteStudentByID(stu *models.Student, student_id string) (err error) {

	err = dao.StudentNew().DeleteStudentByID(stu, student_id)
	if err != nil {
		return err
	}

	return nil
}
