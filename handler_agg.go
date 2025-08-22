package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("too many arguments provided")
	}

	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		return fmt.Errorf("could not fetch feed - %w", err)
	}

	fmt.Printf("%+v\n", rssFeed)

	return nil
}
