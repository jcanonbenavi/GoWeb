package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	stringJSON := `{"name":"John","age":30,"cars":["Ford","BMW","Fiat"]}`
	//fmt.Printf("type: %T\n", stringJSON) //string
	bytesJSON := []byte(stringJSON)

	// convert into maps
	var mapJSON map[string]any
	err := json.Unmarshal(bytesJSON, &mapJSON)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", mapJSON)
	//fmt.Printf("type: %T\n", mapJSON) //MAP

	// fetch key age
	value, ok := mapJSON["age"]
	//fmt.Printf("type: %T\n", value) //float64
	if !ok {
		fmt.Println("age key not found")
		return
	}
	// type assertion
	ageFloat, ok := value.(float64)
	if !ok {
		fmt.Println("age is not a float64")
		return
	}

	// type casting
	age := int(ageFloat)

	age++

	fmt.Println("age:", age)

}
