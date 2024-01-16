package main

import (
	"fmt"
	"os"

	"github.com/jcanonbenavi/app/internal/application"
)

func main() {
	cfg := application.ConfigAppDefault{
		ServerAddr: os.Getenv(""),
		DbFile:     os.Getenv(""),
	}
	app := application.NewApplicationDefault(&cfg)

	// - setup
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}

	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
