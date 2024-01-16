package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jcanonbenavi/app/internal"

	"github.com/jcanonbenavi/app/platform/web/response"
)

type DefaultTicket struct {
	service internal.ServiceTicket
}

func NewDefaulTicket(ticket internal.ServiceTicket) *DefaultTicket {
	return &DefaultTicket{
		service: ticket,
	}
}

// Get returns all the tickets
func (p *DefaultTicket) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		elements, err := p.service.Get(context.Background())
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrAllTickets):
				response.Text(w, http.StatusNotFound, "Tickets not found")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"Tickets": elements,
		})
	}
}

// GetTicketsAmountByDestinationCountry returns the tickets filtered by destination country
func (p *DefaultTicket) GetTicketsAmountByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		country := chi.URLParam(r, "country")
		countries, err := p.service.GetTicketsAmountByDestinationCountry(country)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrCountryNotFound):
				response.Text(w, http.StatusNotFound, "Country not found")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			country: countries,
		})

	}
}

// GetTotalTickets returns the total number of tickets
func (p *DefaultTicket) GetTotalTickets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		elements, err := p.service.GetTotalTickets()
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrAllTickets):
				response.Text(w, http.StatusNotFound, "Tickets not found")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"Total of tickets": elements,
		})
	}
}

// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
func (p *DefaultTicket) GetPercentageTicketsByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		country := chi.URLParam(r, "country")
		percentage, err := p.service.GetPercentageTicketsByDestinationCountry(country)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrCountryNotFound):
				response.Text(w, http.StatusNotFound, "Country not found")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			country: percentage,
		})

	}
}
