package loader

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/jcanonbenavi/app/internal"
)

func LoadDataFromJson() (product []internal.Product, err error) {
	fileContent, err := os.ReadFile("products.json")
	if err != nil {
		err = errors.New("Error reading file")
	}
	errorjson := json.Unmarshal(fileContent, &product)
	if errorjson != nil {
		err = errors.New("Error unmarshaling file")
	}
	return
}
