package main

import (
	"fmt"

	"github.com/jcanonbenavi/code/internal/application"
)

func main() {
	app := application.NewDefaultHTTP(":8080")

	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
