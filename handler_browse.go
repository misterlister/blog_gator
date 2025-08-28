package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/microcosm-cc/bluemonday"

	"github.com/misterlister/blog_gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	browseLimit := 2

	if len(cmd.args) == 1 {
		specifiedLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %w", err)
		}
		if specifiedLimit <= 0 {
			return fmt.Errorf("limit must be positive")
		}
		browseLimit = specifiedLimit
	}

	if len(cmd.args) > 1 {
		return errors.New("too many arguments provided")
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(browseLimit),
	}

	posts, err := s.db.GetPostsForUser(context.Background(), params)

	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("%v\n", cleanSnippet(post.Description.String, 200))
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("===============================")
	}

	return nil
}

var sanitize = bluemonday.StrictPolicy()

func cleanSnippet(html string, max int) string {
	txt := sanitize.Sanitize(html)
	txt = strings.TrimSpace(txt)
	if len(txt) > max {
		return txt[:max] + "..."
	}
	return txt
}
