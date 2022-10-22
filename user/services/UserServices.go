package services

import (
	"github.com/bardiesel/jan/interfaces"
	"github.com/bardiesel/jan/models"
)

type UserService struct {
	repository interfaces.RepositoryInterface[models.User]
}

func (u *UserService) FindAll() []models.User {
	return u.repository.FindAll()
}

func (u *UserService) FindById(id int) models.User {
	return u.repository.FindById(id)
}

func (u *UserService) Insert(data models.User) models.User {
	return u.repository.Insert(data)
}

func (u *UserService) Update(data models.User) models.User {
	return u.repository.Update(data)
}

func (u *UserService) Delete(id int) bool {
	return u.repository.Delete(id)
}

func NewUserService(repository interfaces.RepositoryInterface[models.User]) *UserService {
	return &UserService{repository}
}
