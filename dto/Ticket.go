package dto

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type NewTicket struct {
	SectionID uint64 `json:"section_id"`
	Text      string `json:"ticket_text"`
	Author    OutUser
	Ip        string
	//Files        []*inpFile `json:"files"`
}

type OutTicketForList struct {
	ID           uint64    `json:"ticket_id"`
	Date         time.Time `json:"ticket_date"`
	Significant  bool      `json:"significant"`
	Category     string    `json:"category"`
	Section      string    `json:"section"`
	Text         string    `json:"ticket_text"`
	Status       string    `json:"status"`
	StatusID     uint64    `json:"ticket_status_id"`
	Filial       string    `json:"filial"`
	Author       string    `json:"ticket_author"`
	Support      string    `json:"support"`
	Grade        uint      `json:"grade"`
	SortPriority uint      `json:"sort_priority,omitempty"`
}

func NewTicketToModelTicket(tick NewTicket) *internal_models.Ticket {
	return &internal_models.Ticket{
		CatSect: &internal_models.SectionWithCategory{ID: tick.SectionID},
		Text:    tick.Text,
		Author:  ToModelUser(tick.Author),
		Status:  &internal_models.TicketStatus{},
	}
}

func ToOutTicketForList(tick *internal_models.Ticket, priority map[uint]uint) OutTicketForList {
	outTick := OutTicketForList{
		ID:       tick.ID,
		Date:     tick.Date,
		Category: tick.CatSect.Cat.Name,
		Section:  tick.CatSect.Name,
		Text:     tick.Text,
		Status:   tick.Status.Name,
		StatusID: tick.Status.ID,
		Filial:   tick.Filial,
		Author:   tick.Author.Name,
		Grade:    tick.Grade,
	}

	if tick.CatSect.Significant || tick.CatSect.Cat.Significant {
		outTick.Significant = true
	} else {
		outTick.Significant = false
	}

	if tick.Support != nil {
		outTick.Support = tick.Support.Name
	}

	if priority != nil {
		outTick.SortPriority = priority[uint(tick.Status.ID)]
	} else {
		outTick.SortPriority = 1
	}

	return outTick
}
