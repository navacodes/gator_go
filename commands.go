package main

import (
	"errors"
	"github.com/navacodes/gator_go/internal/config"
)

type State struct {
	cfg *config.Config
}

type Command struct {
	Name string
	Args []string
}

type CommandRegistry struct {
	registeredCommands map[string]func(*State, Command) error
}

func (c *CommandRegistry) register(name string, f func(*State, Command) error) {
	c.registeredCommands[name] = f
}

func (c *CommandRegistry) run(s *State, cmd Command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}
