package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/misterlister/blog_gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return errors.New("please provide the URL of the feed to follow")
	}

	if len(cmd.args) > 1 {
		return errors.New("too many arguments provided")
	}

	feedURL := cmd.args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)

	if err != nil {
		return err
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), params)

	if err != nil {
		return err
	}

	fmt.Printf("%s is now following %s\n", user.Name, feed.Name)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return errors.New("too many arguments provided")
	}

	feedList, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return err
	}

	if len(feedList) == 0 {
		fmt.Printf("%s is following no feeds\n", user.Name)
		return nil
	}

	fmt.Printf("%s is following:\n", user.Name)

	for _, feed := range feedList {
		fmt.Printf(" - %s\n", feed.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("no feed to unfollow was provided")
	}

	if len(cmd.args) > 1 {
		return errors.New("too many arguments provided")
	}

	feedURL := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)

	if err != nil {
		return err
	}

	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.db.DeleteFeedFollow(context.Background(), params)

	if err != nil {
		return err
	}

	return nil
}
