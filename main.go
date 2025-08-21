package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/misterlister/blog_gator/internal/config"
	"github.com/misterlister/blog_gator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	var st state

	st.cfg = cfg

	db, err := sql.Open("postgres", st.cfg.DBUrl)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	dbQueries := database.New(db)

	st.db = dbQueries

	var cmds commands

	cmds.cmdList = make(map[string]func(*state, command) error)

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

	if len(os.Args) < 2 {
		fmt.Println("error: no command specified")
		os.Exit(1)
	}

	var currCmd command

	currCmd.name = os.Args[1]

	if len(os.Args) > 2 {
		currCmd.args = os.Args[2:]
	}

	err = cmds.run(&st, currCmd)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
