package comment_manager

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type ICommentUsecase interface {
	CreateComment(comment *internal_models.Comment) (uint64, models.Err)
	GetTicketComments(ticketID uint64) ([]*internal_models.Comment, models.Err)
}
