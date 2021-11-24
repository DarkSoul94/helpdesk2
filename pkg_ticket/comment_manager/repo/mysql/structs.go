package mysql

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type CommentRepo struct {
	db *sqlx.DB
}

type dbComment struct {
	ID       uint64    `db:"comment_id"`
	TicketId uint64    `db:"ticket_id"`
	Date     time.Time `db:"comment_date"`
	AuthorID uint64    `db:"comment_author_id"`
	Text     string    `db:"comment_text"`
}
