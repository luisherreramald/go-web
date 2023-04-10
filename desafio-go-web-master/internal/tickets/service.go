package tickets

import (
	"context"
)

type Service interface {
	GetTotalTickets(context context.Context, destination string) (int, error)
	AverageDestination(context context.Context, destination string) (float64, error)
}

func NewService(rp Repository) Service {
	return ServiceTickets{rp: rp}
}

type ServiceTickets struct {
	rp Repository
}

func (ticketsService ServiceTickets) GetTotalTickets(context context.Context, destination string) (numTickets int, err error) {
	ticketsByDestination, err := ticketsService.rp.GetTicketByDestination(context, destination)

	if err != nil {
		return 0, err
	}

	return len(ticketsByDestination), nil
}

func (ticketsService ServiceTickets) AverageDestination(context context.Context, destination string) (averageTickets float64, err error) {
	ticketsByDestination, err := ticketsService.rp.GetTicketByDestination(context, destination)

	if err != nil {
		return 0, err

	}

	ticketsAll, err := ticketsService.rp.GetAll(context)

	if err != nil {
		return 0.0, nil
	}

	return float64(len(ticketsByDestination)) / float64(len(ticketsAll)), nil
}
