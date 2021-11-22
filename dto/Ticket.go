package dto

import (
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type NewTicket struct {
	SectionID uint64 `json:"section_id"`
	Text      string `json:"ticket_text"`
	Author    OutUser
	Ip        string
	//Files        []*inpFile `json:"files"`
}

func NewTicketToModelTicket(tick NewTicket) *internal_models.Ticket {
	return &internal_models.Ticket{
		CatSect: &internal_models.CategorySection{ID: tick.SectionID},
		Text:    tick.Text,
		Author:  ToModelUser(tick.Author),
		Status:  &internal_models.TicketStatus{},
	}
}
