package internal

import (
	"context"
	"errors"
)

var (
	// ErrCountryNotFound is the error returned when a country is not found
	ErrCountryNotFound = errors.New("Country not found")
	ErrAllTickets      = errors.New("Not tickets found")
)

// RepositoryTicket represents the repository interface for tickets
type RepositoryTicket interface {
	// GetAll returns all the tickets
	Get(ctx context.Context) (t map[int]TicketAttributes, err error)

	// GetTotalAmountTickets returns the total amount of tickets
	GetTotalTickets() (total int, err error)

	// GetTicketByDestinationCountry returns the tickets filtered by destination country
	GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]TicketAttributes, err error)

	// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
	// ...
	GetPercentageTicketsByDestinationCountry(country string) (percentage float64, err error)
}
