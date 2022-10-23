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

func (u *UserController) Insert(name string) models.User {
	return u.service.Insert(name)
}

func (u *UserController) Update(id int, name string) models.User {
	return u.service.Update(id, name)
}

func (u *UserController) Delete(id int) bool {
	return u.service.Delete(id)
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service}
}
