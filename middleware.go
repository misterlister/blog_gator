package main

import (
	"context"
	"errors"

	"github.com/misterlister/blog_gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		username := s.cfg.CurrentUserName

		if username == "" {
			return errors.New("not logged in")
		}

		user, err := s.db.GetUserByName(context.Background(), username)

		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
