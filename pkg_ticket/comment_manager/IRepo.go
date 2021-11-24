package comment_manager

import "github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"

type ICommentRepo interface {
	CreateComment(comment *internal_models.Comment) (uint64, error)
	GetTicketComments(ticketID uint64) ([]*internal_models.Comment, error)

	Close() error
}
