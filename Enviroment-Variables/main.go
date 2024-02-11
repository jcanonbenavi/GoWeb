package main

import (
	"fmt"
	"os"
)

func main() {
	// read enviroment variable
	goPath := os.Getenv("FILE_PATH")
	fmt.Println(goPath)

	// set enviroment variable
	os.Setenv("KEY", "value")
	println(os.Getenv("KEY"))
}
