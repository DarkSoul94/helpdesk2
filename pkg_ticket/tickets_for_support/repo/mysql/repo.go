package mysql

import (
	"database/sql"
	"fmt"

	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func NewTicketRepoForSupport(db *sql.DB) *Repo {
	return &Repo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *Repo) GetTicketsCount(supportID, statusID uint64) int {
	var count int
	if supportID != 0 {
		query := getCountWithSupp()
		if err := r.db.Get(&count, query, supportID, statusID); err != nil {
			logger.LogError(
				"Failed read tickets count",
				"pkg_ticket/repo/mysql",
				fmt.Sprintf("support id: %d, status id: %d", supportID, statusID),
				err)
		}
		return count
	}
	query := getCountWithoutSupp()
	if err := r.db.Get(&count, query, statusID); err != nil {
		logger.LogError(
			"Failed read tickets count",
			"pkg_ticket/repo/mysql",
			fmt.Sprintf("support id: %d, status id: %d", supportID, statusID),
			err)
	}
	return count
}

func (r *Repo) GetTodayTicketsCount(supportID, statusID uint64) int {
	var count int
	if supportID != 0 {
		query := getTodaysCountWithSupp()
		if err := r.db.Get(&count, query, supportID, statusID); err != nil {
			logger.LogError(
				"Failed read tickets count",
				"pkg_ticket/repo/mysql",
				fmt.Sprintf("support id: %d, status id: %d", supportID, statusID),
				err)
		}
		return count
	}

	query := getTodaysCountWithoutSupp()
	if err := r.db.Get(&count, query, statusID); err != nil {
		logger.LogError(
			"Failed read tickets count",
			"pkg_ticket/repo/mysql",
			fmt.Sprintf("support id: %d, status id: %d", supportID, statusID),
			err)
	}
	return count
}

func (r *Repo) Close() error {
	r.db.Close()
	return nil
}
