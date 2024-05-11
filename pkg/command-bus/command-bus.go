package command_bus

import "fmt"

// CommandBus
type CommandBus struct {
	handlers map[string]Handler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{make(map[string]Handler)}
}

func (b *CommandBus) Register(command Command, handler Handler) {
	name := fmt.Sprintf("%T", command) // use the type of the command as its name
	fmt.Printf("Registering handler for command %s\n", name)
	b.handlers[name] = handler
}

func (b *CommandBus) Send(command Command) error {
	name := fmt.Sprintf("%T", command) // use the type of the command as its name
	fmt.Printf("Sending command %s\n", name)
	if h, ok := b.handlers[name]; ok {
		return h.Handle(command)
	}
	return fmt.Errorf("no handler found for command %s", name)
}

// Command interface
type Command interface {
	ID() string
	Data() map[string]interface{}
}

// Command Handler interface
type Handler interface {
	Handle(command Command) error
}
