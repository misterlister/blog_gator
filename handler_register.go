package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/misterlister/blog_gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("register requires a username argument")
	}

	if len(cmd.args) > 1 {
		return errors.New("too many arguments provided")
	}

	params := database.CreateUserParams{
		Name:      cmd.args[0],
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err := s.db.CreateUser(context.Background(), params)

	if err != nil {
		return fmt.Errorf("could not create user '%s'", cmd.args[0])
	}

	err = s.cfg.SetUser(user.Name)

	if err != nil {
		return fmt.Errorf("couldn't log in as user '%s': %w", user.Name, err)
	}

	fmt.Printf("user '%s' has been created and logged in!\n", user.Name)

	fmt.Println("ID:", user.ID)
	fmt.Println("Name:", user.Name)
	fmt.Println("Created Time:", user.CreatedAt)
	fmt.Println("Updated Time:", user.UpdatedAt)

	return nil
}
