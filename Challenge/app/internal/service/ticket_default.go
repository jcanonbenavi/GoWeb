package service

import (
	"context"
	"fmt"

	"github.com/jcanonbenavi/app/internal"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp internal.RepositoryTicket
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp internal.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// Get returns all the tickets
func (s *ServiceTicketDefault) Get(ctx context.Context) (t map[int]internal.TicketAttributes, err error) {
	t, err = s.rp.Get(context.Background())
	if err != nil {
		err = internal.ErrAllTickets
	}
	return
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalTickets() (total int, err error) {
	total, err = s.rp.GetTotalTickets()
	if err != nil {
		err = internal.ErrAllTickets
	}
	return
}

// GetTicketsAmountByDestinationCountry returns the tickets filtered by destination country
func (s *ServiceTicketDefault) GetTicketsAmountByDestinationCountry(country string) (total map[int]internal.TicketAttributes, err error) {
	total, err = s.rp.GetTicketsByDestinationCountry(context.Background(), country)
	if err != nil {
		switch err {
		case internal.ErrCountryNotFound:
			err = fmt.Errorf("%w: Country", internal.ErrCountryNotFound)
		}
	}
	return
}

// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(country string) (percentage float64, err error) {
	percentage, err = s.rp.GetPercentageTicketsByDestinationCountry(country)
	if err != nil {
		switch err {
		case internal.ErrCountryNotFound:
			err = fmt.Errorf("%w: Country", internal.ErrCountryNotFound)
		}
	}
	return
}
