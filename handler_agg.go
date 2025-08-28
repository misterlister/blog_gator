package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/misterlister/blog_gator/internal/database"
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

		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{Time: t, Valid: true}
		} else if t, err := time.Parse(time.RFC1123, item.PubDate); err == nil {
			publishedAt = sql.NullTime{Time: t, Valid: true}
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		params := database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: publishedAt,
			FeedID:      nextFeed.ID,
		}
		_, err := s.db.CreatePost(context.Background(), params)

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
}
