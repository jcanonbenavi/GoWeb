package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	height   float64 `json:"height,omitempty"` //not exported
	Weight   float64 `json:"weight"`
	Password string  `json:"-"`
}

func main() {
	// create person struct
	p := Person{
		Name:     "john",
		Age:      30,
		height:   0.0,
		Weight:   80.5,
		Password: "123456",
	}

	fmt.Println(p)
	fmt.Printf("type: %T\n", p) //type: main.Person

	// convert into json bytes.
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert into json string.
	data := string(bytes)
	fmt.Printf("type: %T\n", data) //string
	fmt.Println(data)
}
