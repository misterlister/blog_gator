package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("too many arguments provided")
	}

	err := s.db.ResetUsersTable(context.Background())

	if err != nil {
		return fmt.Errorf("could not reset users: %w", err)
	}

	fmt.Println("Reset users table!")

	return nil
}
