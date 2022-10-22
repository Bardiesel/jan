package main

import (
	"fmt"

	"github.com/bardiesel/jan/controllers"
	"github.com/bardiesel/jan/models"
	"github.com/bardiesel/jan/repositories"
	"github.com/bardiesel/jan/services"
)

func main() {
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	var cmd string
	fmt.Scanf("%s", &cmd)

	switch cmd {
	case "findAll":
		fmt.Println(userController.FindAll())
	case "findById":
		var id int
		fmt.Scanf("%d", &id)
		fmt.Println(userController.FindById(id))
	case "insert":
		var name string
		fmt.Scanf("%s", &name)
		fmt.Println(userController.Insert(models.User{Name: name}))
	case "update":
		var id int
		var name string
		fmt.Scanf("%d %s", &id, &name)
		fmt.Println(userController.Update(models.User{ID: id, Name: name}))
	case "delete":
		var id int
		fmt.Scanf("%d", &id)
		fmt.Println(userController.Delete(id))
	default:
		fmt.Println("Command not found")
	}
}
