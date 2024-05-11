package main

import (
	"fmt"

	user "github.com/AntonioMartinezFernandez/go-command-bus-example/internal/user"
	command_bus "github.com/AntonioMartinezFernandez/go-command-bus-example/pkg/command-bus"
)

func main() {
	// Create a new bus
	bus := command_bus.NewCommandBus()

	// Register handlers
	createUserHandler := &user.CreateUserHandler{}
	removeUserHandler := &user.RemoveUserHandler{}
	bus.Register(&user.CreateUserCommand{}, createUserHandler)
	bus.Register(&user.RemoveUserCommand{}, removeUserHandler)

	// Send a command to create a user
	cuErr := bus.Send(user.NewCreateUserCommand("123", map[string]interface{}{"name": "John Doe"}, map[string]interface{}{"service": "user"}))
	if cuErr != nil {
		fmt.Println("Error: ", cuErr)
	}

	// Send a command to remove a user
	ruErr := bus.Send(user.NewRemoveUserCommand("123", map[string]interface{}{"name": "John Doe"}))
	if ruErr != nil {
		fmt.Println("Error: ", ruErr)
	}
}
