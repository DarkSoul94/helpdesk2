package mysql

import (
	"database/sql"
	"strconv"

	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
	"github.com/jmoiron/sqlx"
)

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *Repo) GetGroupByID(groupID uint64) (*group_manager.Group, error) {
	var (
		group dbGroup
		query string
		err   error
	)
	query = `SELECT * FROM user_groups WHERE group_id = ?`
	err = r.db.Get(&group, query, groupID)
	if err != nil {
		logger.LogError(ErrReadGroup.Error(), "user_manager/repo/mysql", strconv.FormatUint(groupID, 10), err)
		return &group_manager.Group{}, ErrReadGroup
	}

	return r.toModelGroup(group), nil
}

func (r *Repo) GetGroupList() ([]*group_manager.Group, error) {
	var (
		dbGroupsList []dbGroup
		err          error
	)
	query := `SELECT * FROM user_groups`

	if err = r.db.Select(&dbGroupsList, query); err != nil {
		logger.LogError(ErrReadGroupsList.Error(), "user_manager/repo/mysql", "", err)
		return []*group_manager.Group{}, ErrReadGroupsList
	}

	mGroupsList := make([]*group_manager.Group, 0)
	for _, val := range dbGroupsList {
		mGroupsList = append(mGroupsList, r.toModelGroup(val))
	}
	return mGroupsList, nil

}

func (r *Repo) Close() error {
	r.db.Close()
	return nil
}
