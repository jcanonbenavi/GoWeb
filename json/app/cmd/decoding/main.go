package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Cars []string `json:"cars"`
}

func main() {
	reader := strings.NewReader(`{"name":"john","age":30,"cars":["ford","chevy","bmw"]}`)
	decoding := json.NewDecoder(reader)

	var p Person

	if err := decoding.Decode(&p); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", p)
}
