package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/misterlister/blog_gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("must provide a name and URL to add a feed")
	}

	if len(cmd.args) > 2 {
		return errors.New("too many arguments provided. Please supply only a name and URL")
	}

	feedName := cmd.args[0]
	feedURL := cmd.args[1]
	username := s.cfg.CurrentUserName

	user, err := s.db.GetUserByName(context.Background(), username)

	if err != nil {
		return err
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	}

	newFeed, err := s.db.CreateFeed(context.Background(), params)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", newFeed)

	return nil
}
