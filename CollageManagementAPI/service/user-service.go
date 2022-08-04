package service

import (
	"gorm-test/dao"
	"gorm-test/models"
)

type UserService interface {
	CreateUser() func()
	GetUsersService(User *[]models.User) (func(), error)
	GetUser() func()
	UpdateUser() func()
	DeleteUser() func()
}

type myUserService struct {
	dao dao.UserRepo
}

func New(dao dao.UserRepo) UserService {
	return &myUserService{
		dao: dao,
	}
}
func (service myUserService) CreateUser() func() {

	return func() {

	}
}
func (service myUserService) GetUsersService(User *[]models.User) (func(), error) {

	var err error

	return func() {
		service.dao.GetUsers(User)
	}, err
}

func (service myUserService) GetUser() func() {

	return func() {

	}
}

func (service myUserService) UpdateUser() func() {

	return func() {

	}
}

func (service myUserService) DeleteUser() func() {

	return func() {

	}
}
