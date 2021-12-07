package internal_models

import "time"

type ReturnedTicket struct {
	TicketID    uint64
	TicketDate  time.Time
	Category    string
	Section     string
	TicketText  string
	Status      string
	Author      string
	Support     string
	TicketGrade uint64
}
