package user

import (
	"errors"
	"fmt"

	command_bus "github.com/AntonioMartinezFernandez/go-command-bus-example/pkg/command-bus"
)

// RemoveUserCommand is a command to remove a user
type RemoveUserCommand struct {
	id   string
	data map[string]interface{}
}

func NewRemoveUserCommand(id string, data map[string]interface{}) *RemoveUserCommand {
	return &RemoveUserCommand{id: id, data: data}
}

func (c *RemoveUserCommand) ID() string {
	return c.id
}

func (c *RemoveUserCommand) Data() map[string]interface{} {
	return c.data
}

// RemoveUserHandler is a handler that remove users
type RemoveUserHandler struct{}

func (h *RemoveUserHandler) Handle(c command_bus.Command) error {
	cmd, ok := c.(*RemoveUserCommand)
	if !ok {
		return errors.New("invalid command")
	}

	fmt.Printf("Removing user with ID: %s, Data: %v\n", cmd.ID(), cmd.Data())
	return nil
}
