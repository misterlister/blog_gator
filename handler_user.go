package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("login requires a username argument")
	}

	if len(cmd.args) > 1 {
		return errors.New("too many arguments provided")
	}

	username := cmd.args[0]

	user, err := s.db.GetUserByName(context.Background(), username)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("user '%s' does not exist in the database", username)
		}
		return fmt.Errorf("error querying user '%s' : %w", username, err)
	}

	err = s.cfg.SetUser(user.Name)

	if err != nil {
		return fmt.Errorf("couldn't log in as user '%s': %w", user.Name, err)
	}

	fmt.Printf("user '%s' has been logged in!\n", user.Name)

	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("too many arguments provided")
	}

	users, err := s.db.GetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("could not retrieve users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
