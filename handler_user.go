package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("login requires a username argument")
	}

	if len(cmd.args) > 1 {
		return errors.New("login requires only a username argument")
	}

	username := cmd.args[0]

	err := s.cfg.SetUser(username)

	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("user '%s' has been logged in!\n", username)

	return nil
}
