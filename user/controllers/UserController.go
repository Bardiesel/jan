package controllers

import (
	"github.com/bardiesel/jan/models"
	"github.com/bardiesel/jan/services"
)

type UserController struct {
	service *services.UserService
}

func (u *UserController) FindAll() []models.User {
	return u.service.FindAll()
}

func (u *UserController) FindById(id int) models.User {
	return u.service.FindById(id)
}

func (u *UserController) Insert(data models.User) models.User {
	return u.service.Insert(data)
}

func (u *UserController) Update(data models.User) models.User {
	return u.service.Update(data)
}

func (u *UserController) Delete(id int) bool {
	return u.service.Delete(id)
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service}
}
