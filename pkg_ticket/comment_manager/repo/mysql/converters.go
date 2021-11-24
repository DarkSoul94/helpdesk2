package mysql

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

func (r *CommentRepo) toDbComment(comment *internal_models.Comment) dbComment {
	return dbComment{
		ID:       comment.ID,
		TicketId: comment.TicketId,
		Date:     comment.Date,
		AuthorID: comment.Author.ID,
		Text:     comment.Text,
	}
}

func (r *CommentRepo) toModelComment(comment dbComment) *internal_models.Comment {
	return &internal_models.Comment{
		ID:       comment.ID,
		TicketId: comment.TicketId,
		Date:     comment.Date,
		Author:   &models.User{ID: comment.AuthorID},
		Text:     comment.Text,
	}
}
