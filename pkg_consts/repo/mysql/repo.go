package mysql

import (
	"database/sql"
	"fmt"

	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type ConstsRepo struct {
	db *sqlx.DB
}

func NewConstsRepo(db *sql.DB) *ConstsRepo {
	return &ConstsRepo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *ConstsRepo) SetConst(key string, data interface{}) error {
	var (
		constant dBConst
		query    string
		err      error
	)

	constant.ToConst(key, data)

	query = `
		INSERT INTO consts SET
			name = :name,
			data = :data,
			data_type = :data_type,
			table_name = :table_name
		ON DUPLICATE KEY UPDATE
			data = :data,
			data_type = :data_type,
			table_name = :table_name`

	_, err = r.db.NamedExec(query, constant)
	if err != nil {
		logger.LogError(
			"Failed modify const",
			"pkg_consts/repo/mysql",
			fmt.Sprintf("name: %s;", key),
			err,
		)
	}
	return err
}

func (r *ConstsRepo) GetConst(key string, target interface{}) error {
	var (
		constant dBConst
		query    string
		err      error
	)

	query = `SELECT * FROM consts
				WHERE name = ?
				LIMIT 1`
	err = r.db.Get(&constant, query, key)
	if err != nil {
		logger.LogError(
			"Failed get const",
			"pkg_consts/repo/mysql",
			fmt.Sprintf("name: %s;", key),
			err,
		)
		return err
	}

	err = constant.FromConst(target)
	if err != nil {
		logger.LogError(
			fmt.Sprintf("Failed convert db_const to %s", constant.DataType),
			"pkg_consts/repo/mysql",
			fmt.Sprintf("name: %s;", key),
			err,
		)
		return err
	}

	return nil
}

func (r *ConstsRepo) CreateHistory(name string, data interface{}) error {
	var history dbHistory
	history.ToHistory(name, data)
	query := `
		INSERT INTO const_change_history SET
			date = :date,
			name = :name,
			val = :val,
			val_type = :val_type
		ON DUPLICATE KEY UPDATE
			val = :val,
			val_type = :val_type`
	if _, err := r.db.NamedExec(query, history); err != nil {
		logger.LogError(
			"Failed modify const",
			"pkg_consts/repo/mysql",
			fmt.Sprintf("name: %s;", name),
			err,
		)
		return err
	}
	return nil
}

func (r *ConstsRepo) Close() error {
	return r.db.Close()
}
