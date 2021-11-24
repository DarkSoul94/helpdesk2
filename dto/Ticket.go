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

type OutTicket struct {
	ID              uint64                  `json:"ticket_id"`
	Date            time.Time               `json:"ticket_date"`
	CategorySection *OutSectionWithCategory `json:"category_section"`
	Text            string                  `json:"ticket_text"`
	Status          *OutTicketStatus        `json:"ticket_status"`
	Filial          string                  `json:"filial"`
	IP              string                  `json:"ip"`
	Author          *OutUserWithOutGroup    `json:"ticket_author"`
	Support         *OutUserWithOutGroup    `json:"support"`
	ResolvedUser    *OutUserWithOutGroup    `json:"resolved_user"`
	ServiceComment  string                  `json:"service_comment"`
	//Comment         []*outComment       `json:"comments"`
	//Files           []*outFiles         `json:"files"`
}

func NewTicketToModelTicket(tick NewTicket) *internal_models.Ticket {
	return &internal_models.Ticket{
		CatSect: &internal_models.SectionWithCategory{ID: tick.SectionID},
		Text:    tick.Text,
		Author:  ToModelUser(tick.Author),
		Status:  &internal_models.TicketStatus{},
		IP:      tick.Ip,
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

func ToOutTicket(ticket *internal_models.Ticket) OutTicket {
	section := ToOutSectionWithCategory(ticket.CatSect)
	status := ToOutTicketStatus(ticket.Status)

	outTicket := OutTicket{
		ID:              ticket.ID,
		Date:            ticket.Date,
		CategorySection: &section,
		Text:            ticket.Text,
		Status:          &status,
		Filial:          ticket.Filial,
		IP:              ticket.IP,
		ServiceComment:  ticket.ServiceComment,
	}

	if ticket.Author != nil {
		author := ToOutUserWithOutGroup(ticket.Author)
		outTicket.Author = &author
	}

	if ticket.Support != nil {
		support := ToOutUserWithOutGroup(ticket.Support)
		outTicket.Support = &support
	}

	if ticket.ResolvedUser != nil {
		resolvUser := ToOutUserWithOutGroup(ticket.ResolvedUser)
		outTicket.ResolvedUser = &resolvUser
	}

	return outTicket
}
