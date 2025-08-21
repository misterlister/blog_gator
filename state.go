package main

import (
	"github.com/misterlister/blog_gator/internal/config"
	"github.com/misterlister/blog_gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}
