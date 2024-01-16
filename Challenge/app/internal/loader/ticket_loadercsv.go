package loader

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/jcanonbenavi/app/internal"
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {
	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type LoaderTicketCSV struct {
	filePath string
}

// Load loads the tickets from the CSV file
func (t *LoaderTicketCSV) Load() (ticket map[int]internal.TicketAttributes, err error) {
	// open the file
	file, err := os.Open(t.filePath)

	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}
	defer file.Close()

	// read the records
	ticket = make(map[int]internal.TicketAttributes)

	// read the file
	r := csv.NewReader(file)
	record, err := r.ReadAll()

	for key, values := range record {
		price, _ := strconv.ParseFloat(values[5], 64)
		ticket[key+1] = internal.TicketAttributes{
			Name:    values[1],
			Email:   values[2],
			Country: values[3],
			Hour:    values[4],
			Price:   price,
		}
	}
	return
}
