package dto

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type NewTicket struct {
	SectionID uint64 `json:"section_id"`
	Text      string `json:"ticket_text"`
	Author    OutUser
	Ip        string
	//TODO: Files        []*inpFile `json:"files"`
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
	ID             uint64                  `json:"ticket_id"`
	Date           time.Time               `json:"ticket_date"`
	Section        *OutSectionWithCategory `json:"category_section"`
	Text           string                  `json:"ticket_text"`
	Status         *OutTicketStatus        `json:"ticket_status"`
	Filial         string                  `json:"filial"`
	IP             string                  `json:"ip"`
	Author         *OutUserWithOutGroup    `json:"ticket_author"`
	Support        *OutUserWithOutGroup    `json:"support"`
	ResolvedUser   *OutUserWithOutGroup    `json:"resolved_user"`
	ServiceComment string                  `json:"service_comment"`
	Comments       []*OutComment           `json:"comments"`
	//TODO: Files           []*outFiles         `json:"files"`
}

type InpUpdateTicket struct {
	ID             uint64 `json:"ticket_id"`
	SectionID      uint64 `json:"section_id"`
	StatusID       uint64 `json:"ticket_status_id"`
	SupportID      uint64 `json:"support_id"`
	ServiceComment string `json:"service_comment"`
	//TODO: Files             []*inpFile `json:"files"`
}

type InpGenerateTicket struct {
	Text      string               `json:"text,omitempty"`
	SectionID uint64               `json:"section_id,omitempty"`
	Users     []inpUserForGenerate `json:"users,omitempty"`
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
		ID:             ticket.ID,
		Date:           ticket.Date,
		Section:        &section,
		Text:           ticket.Text,
		Status:         &status,
		Filial:         ticket.Filial,
		IP:             ticket.IP,
		ServiceComment: ticket.ServiceComment,
	}

	if ticket.Author != nil && ticket.Author.ID != 0 {
		author := ToOutUserWithOutGroup(ticket.Author)
		outTicket.Author = &author
	}

	if ticket.Support != nil && ticket.Support.ID != 0 {
		support := ToOutUserWithOutGroup(ticket.Support)
		outTicket.Support = &support
	}

	if ticket.ResolvedUser != nil && ticket.ResolvedUser.ID != 0 {
		resolvUser := ToOutUserWithOutGroup(ticket.ResolvedUser)
		outTicket.ResolvedUser = &resolvUser
	}

	if ticket.Comments != nil {
		for _, comment := range ticket.Comments {
			outTicket.Comments = append(outTicket.Comments, ToOutComment(comment))
		}
	}

	return outTicket
}

func UpdateTicketToModel(ticket InpUpdateTicket) *internal_models.Ticket {
	return &internal_models.Ticket{
		ID:             ticket.ID,
		CatSect:        &internal_models.SectionWithCategory{ID: ticket.SectionID},
		Status:         &internal_models.TicketStatus{ID: ticket.StatusID},
		Support:        &models.User{ID: ticket.SupportID},
		ServiceComment: ticket.ServiceComment,
	}
}

func ToModelGenerateTicekt(ticket InpGenerateTicket) internal_models.TicketGenerate {
	mTicket := internal_models.TicketGenerate{
		Text:      ticket.Text,
		SectionID: ticket.SectionID,
	}

	for _, user := range ticket.Users {
		mTicket.Users = append(mTicket.Users, ToModelUserForGenerate(user))
	}

	return mTicket
}
