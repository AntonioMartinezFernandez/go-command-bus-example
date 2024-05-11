package user

import (
	"errors"
	"fmt"

	command_bus "github.com/AntonioMartinezFernandez/go-command-bus-example/pkg/command-bus"
)

// CreateUserCommand is a command to create a user
type CreateUserCommand struct {
	id   string
	data map[string]interface{}
	meta map[string]interface{}
}

func NewCreateUserCommand(id string, data map[string]interface{}, meta map[string]interface{}) *CreateUserCommand {
	return &CreateUserCommand{id: id, data: data, meta: meta}
}

func (c *CreateUserCommand) ID() string {
	return c.id
}

func (c *CreateUserCommand) Data() map[string]interface{} {
	return c.data
}

func (c *CreateUserCommand) Meta() map[string]interface{} {
	return c.meta
}

// CreateUserHandler is a handler that creates users
type CreateUserHandler struct{}

func (h *CreateUserHandler) Handle(c command_bus.Command) error {
	cmd, ok := c.(*CreateUserCommand)
	if !ok {
		return errors.New("invalid command")
	}

	fmt.Printf("Creating user with ID: %s, Data: %v, Meta: %v\n", cmd.ID(), cmd.Data(), cmd.Meta())
	return nil
}
