package repositories

import (
	"github.com/bardiesel/jan/models"
)

type UserRepository struct {
	model *models.User
}

func (u *UserRepository) GetDB() {
	u.model.TableName()
}

func (u *UserRepository) FindAll() []models.User {
	return []models.User{}
}

func (u *UserRepository) FindById(id int) models.User {
	return models.User{}
}

func (u *UserRepository) Insert(data models.User) models.User {
	data.ID++
	return data
}

func (u *UserRepository) Update(data models.User) models.User {
	return data
}

func (u *UserRepository) Delete(id int) bool {
	return true
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		model: &models.User{},
	}
}
