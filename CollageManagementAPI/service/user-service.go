package service

import (
	"fmt"
	"gorm-test/dao"
	"gorm-test/models"
)

func GetUsers(User *[]models.User) error {

	fmt.Println("Inside service layer")

	err := dao.New().GetUsers(User)
	if err != nil {
		return err
	}

	return nil

}

func CreateUser(User *models.User) error {

	err := dao.New().CreateUser(User)

	if err != nil {
		return err
	}

	return nil
}

func GetUser(User *models.User, id int) (err error) {
	err = dao.New().GetUser(User, id)
	if err != nil {
		return err
	}
	return nil
}

//update user
func UpdateUser(User *models.User) (err error) {
	err = dao.New().UpdateUser(User)

	if err != nil {
		return err
	}
	return nil
}

//delete user
func DeleteUser(User *models.User, id int) (err error) {
	err = dao.New().DeleteUser(User, id)
	if err != nil {
		return err
	}
	return nil
}
