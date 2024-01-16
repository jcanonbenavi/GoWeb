package internal

import "context"

type ServiceTicket interface {
	Get(ctx context.Context) (t map[int]TicketAttributes, err error)
	// GetTotalAmountTickets returns the total amount of tickets
	GetTotalTickets() (total int, err error)

	// GetTicketsAmountByDestinationCountry returns the amount of tickets filtered by destination country
	// ...
	GetTicketsAmountByDestinationCountry(country string) (t map[int]TicketAttributes, err error)

	// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
	// ...
	GetPercentageTicketsByDestinationCountry(country string) (percentage float64, err error)
}
