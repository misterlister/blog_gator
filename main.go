package main

import (
	"fmt"
	"os"

	"github.com/misterlister/blog_gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		return
	}

	var st state

	st.cfg = cfg

	var cmds commands

	cmds.cmdList = make(map[string]func(*state, command) error)

	cmds.register("login", handlerLogin)

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
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
