package repository

import (
	"context"

	"github.com/jcanonbenavi/app/internal"
)

// NewRepositoryTicketMap creates a new repository for tickets in a map
func NewRepositoryTicketMap(db map[int]internal.TicketAttributes, lastId int) *RepositoryTicketMap {
	return &RepositoryTicketMap{
		db:     db,
		lastId: lastId,
	}
}

// RepositoryTicketMap implements the repository interface for tickets in a map
type RepositoryTicketMap struct {
	// db represents the database in a map
	// - key: id of the ticket
	// - value: ticket
	db map[int]internal.TicketAttributes

	// lastId represents the last id of the ticket
	lastId int
}

// GetAll returns all the tickets
func (r *RepositoryTicketMap) Get(ctx context.Context) (t map[int]internal.TicketAttributes, err error) {
	// create a copy of the map
	t = make(map[int]internal.TicketAttributes, len(r.db))
	for k, v := range r.db {
		// assign the value to the new map
		t[k] = v
	}
	// check if the map is empty
	if len(t) == 0 {
		err = internal.ErrAllTickets
	}

	return
}

func (r *RepositoryTicketMap) GetTotalTickets() (total int, err error) {
	// get the length of the map
	total = len(r.db)
	// check if the map is empty
	if total == 0 {
		err = internal.ErrAllTickets
	}
	return
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *RepositoryTicketMap) GetTicketsByDestinationCountry(ctx context.Context, country string) (ticket map[int]internal.TicketAttributes, err error) {
	ticket = make(map[int]internal.TicketAttributes)
	for k, v := range r.db {
		// check if the country is the same
		if v.Country == country {
			ticket[k] = v
		}
	}
	if len(ticket) == 0 {
		err = internal.ErrCountryNotFound
	}
	return
}

// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
func (r *RepositoryTicketMap) GetPercentageTicketsByDestinationCountry(country string) (percentage float64, err error) {
	// get the tickets by destination country
	ticket_by_country, err := r.GetTicketsByDestinationCountry(context.Background(), country)
	if err != nil {
		return
	}
	total_tickets, err := r.GetTotalTickets()
	// calculate the percentage
	percentage = (float64(len(ticket_by_country)) / float64(total_tickets)) * 100
	return
}
