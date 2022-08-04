package dao

import (
	"fmt"
	"gorm-test/database"
	"gorm-test/models"

	"gorm.io/gorm"
)

type DaoInterface interface {
	CreateUser(User *models.User) (err error)
	GetUsers(User *[]models.User) (err error)
	GetUser(User *models.User, id int) (err error)
	UpdateUser(User *models.User) (err error)
	DeleteUser(User *models.User, id int) (err error)
}

type UserRepo struct {
	Db *gorm.DB
}

func New() DaoInterface {
	db := database.InitDb()
	db.AutoMigrate(&models.User{})
	return &UserRepo{Db: db}
}

//get users
func (repository *UserRepo) GetUsers(User *[]models.User) (err error) {

	fmt.Println("Inside get users")
	err = repository.Db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

//create a user
func (repository *UserRepo) CreateUser(User *models.User) (err error) {

	err = repository.Db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get user by id
func (repository *UserRepo) GetUser(User *models.User, id int) (err error) {
	err = repository.Db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

//update user
func (repository *UserRepo) UpdateUser(User *models.User) (err error) {
	repository.Db.Save(User)
	return nil
}

//delete user
func (repository *UserRepo) DeleteUser(User *models.User, id int) (err error) {
	repository.Db.Where("id = ?", id).Delete(User)
	return nil
}
