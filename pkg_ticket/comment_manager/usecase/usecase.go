package usecase

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/comment_manager"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type CommentUsecase struct {
	repo comment_manager.ICommentRepo
}

func NewCommentUsecase(repo comment_manager.ICommentRepo) *CommentUsecase {
	return &CommentUsecase{
		repo: repo,
	}
}

func (u *CommentUsecase) CreateComment(comment *internal_models.Comment) (uint64, models.Err) {
	comment.Date = time.Now().Truncate(time.Microsecond)

	id, err := u.repo.CreateComment(comment)
	if err != nil {
		return 0, models.InternalError(err.Error())
	}

	return id, nil
}

func (u *CommentUsecase) GetTicketComments(ticketID uint64) ([]*internal_models.Comment, models.Err) {
	comments, err := u.repo.GetTicketComments(ticketID)
	if err != nil {
		return nil, models.InternalError(err.Error())
	}

	return comments, nil
}
