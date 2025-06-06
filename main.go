package main

import (
	"fmt"

	"github.com/misterlister/blog_gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		return
	}

	var currentUser = "Hayden"

	err = cfg.SetUser(currentUser)

	if err != nil {
		return
	}

	newCfg, err := config.Read()
	if err != nil {
		return
	}

	fmt.Printf("%+v\n", newCfg)

}
