package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmdList map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmdFunc, ok := c.cmdList[cmd.name]
	if !ok {
		return fmt.Errorf("command '%s' does not exist", cmd.name)
	}

	return cmdFunc(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdList[name] = f
}
