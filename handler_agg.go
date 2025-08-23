package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("no timeBetweenRequests value was passed")
	}

	if len(cmd.args) > 1 {
		return errors.New("too many arguments provided")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])

	if err != nil {
		return fmt.Errorf("couldn't parse time duration - %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}
	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)

	if err != nil {
		fmt.Println(err)
		return
	}

	feedData, err := fetchFeed(context.Background(), nextFeed.Url)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range feedData.Channel.Item {
		fmt.Println(item.Title)
	}
}
