package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]any{
		"name":   "john",
		"age":    30,
		"height": 2.0,
		"wight":  80.5,
	}
	fmt.Printf("type: %T\n", m) //MAP
	fmt.Println(m)

	// convert into json bytes.
	bytes, err := json.Marshal(m)
	fmt.Printf("type: %T\n", bytes) //[]uint8
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert into json string.
	data := string(bytes)
	fmt.Printf("type: %T\n", data) //string
	fmt.Println(data)
}
