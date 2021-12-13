package mysql

import (
	"database/sql"
	"fmt"

	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/jmoiron/sqlx"
)

func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *CommentRepo) CreateComment(comment *internal_models.Comment) (uint64, error) {
	var (
		res   sql.Result
		query string
		err   error
	)

	query = `INSERT INTO comment_history SET
				ticket_id = :ticket_id,
				comment_date = :comment_date,
				comment_author_id = :comment_author_id,
				comment_text = :comment_text`

	res, err = r.db.NamedExec(query, r.toDbComment(comment))
	if err != nil {
		logger.LogError(
			"Failed create comment",
			"pkg_ticket/comment_manager/repo/mysql",
			fmt.Sprintf("ticket id: %d; date: %s; author id: %d; text: %s;", comment.TicketId, comment.Date, comment.Author.ID, comment.Text),
			err,
		)
		return 0, err
	}

	id, _ := res.LastInsertId()

	return uint64(id), nil
}

func (r *CommentRepo) GetTicketComments(ticketID uint64) ([]*internal_models.Comment, error) {
	var (
		dbComments []dbComment
		mComments  []*internal_models.Comment
		query      string
		err        error
	)

	query = `SELECT * FROM comment_history
				WHERE ticket_id = ?
				ORDER BY comment_id`

	err = r.db.Select(&dbComments, query, ticketID)
	if err != nil {
		logger.LogError(
			"Failed read ticket comments",
			"pkg_ticket/comment_manager/repo/mysql",
			fmt.Sprintf("ticket id: %d", ticketID),
			err,
		)

		return nil, err
	}

	for _, comment := range dbComments {
		mComments = append(mComments, r.toModelComment(comment))
	}

	return mComments, nil
}

func (r *CommentRepo) Close() error {
	return r.db.Close()
}
