package models

type User struct {
	ID       int
	Name     string
	Username string
	Password string
	Role     string
}

func (u *User) TableName() string {
	return "users"
}
