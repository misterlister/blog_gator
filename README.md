# blog_gator

A command line tool for aggregating RSS feeds and viewing the posts.

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You can then install `gator` with:

```bash
git clone https://github.com/misterlister/blog_gator
cd blog_gator
go build -o gator 
```

## Config

Before running `gator`, you need to set up your database configuration:

1. Create a `.gatorconfig.json` file in your home directory
2. Add your Postgres connection details:

```json
{
  "db_url": "postgres://username:password@localhost:5432/database_name?sslmode=disable"
}
```

Replace the placeholders with your actual database credentials:

username: your Postgres username
password: your Postgres password (leave empty if none)
database_name: the name of your database
Example:

```json
{
  "db_url": "postgres://john:mypassword@localhost:5432/gator_db?sslmode=disable"
}
```

## Usage

```bash
##################
# User Management:
##################

# Create a new user account
gator register <name>

# Log in as an existing user
gator login <name>

# List all registered users
gator users

##################
# Feed Management:
##################

# Add a new RSS feed to the database
gator addfeed <url>

# List all available feeds
gator feeds

# Follow an existing feed
gator follow <url>

# Unfollow a feed you're currently following
gator unfollow <url>

########################
# Content & Aggregation:
########################

# Start the RSS aggregator (fetches posts every 30 seconds)
gator agg 30s

# Browse recent posts (optionally limit the number shown)
gator browse [limit]

# Reset the database (removes all data)
gator reset
```
