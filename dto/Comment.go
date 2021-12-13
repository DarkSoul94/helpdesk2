package dto

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type InpComment struct {
	TicketId uint64 `json:"ticket_id"`
	Text     string `json:"comment_text"`
}

type OutComment struct {
	ID     uint64    `json:"comment_id"`
	Date   time.Time `json:"comment_date"`
	Author string    `json:"comment_author"`
	Text   string    `json:"comment_text"`
}

func ToModelComment(comment InpComment) *internal_models.Comment {
	return &internal_models.Comment{
		TicketId: comment.TicketId,
		Text:     comment.Text,
		Author:   &models.User{},
	}
}

func ToOutComment(comment *internal_models.Comment) *OutComment {
	return &OutComment{
		ID:     comment.ID,
		Date:   comment.Date,
		Author: comment.Author.Name,
		Text:   comment.Text,
	}
}
