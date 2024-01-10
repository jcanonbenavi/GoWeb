package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	myEncoder := json.NewEncoder(os.Stdout)
	//standar output
	type MyData struct {
		ProductID string
		Price     float64
	}

	data := MyData{
		ProductID: "123",
		Price:     100.5,
	}
	if err := myEncoder.Encode(data); err != nil {
		fmt.Println(err)
		return
	}

}
