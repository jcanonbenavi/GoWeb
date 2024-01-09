package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
}

func main() {
	// json bytes
	bytes := []byte(`{"name":"john","age":30,"Height":2.0,"weight":null}`)

	// convert into struct
	var p Person
	if err := json.Unmarshal(bytes, &p); err != nil {
		fmt.Println(err)
		return
	}

	// print struct
	fmt.Println(p)
}
