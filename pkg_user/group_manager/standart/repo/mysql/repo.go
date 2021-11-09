package mysql

import (
	"database/sql"
	"strconv"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
	"github.com/jmoiron/sqlx"
)

func NewGroupRepo(db *sql.DB) *GroupRepo {
	return &GroupRepo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *GroupRepo) GetGroupByID(groupID uint64) (*group_manager.Group, models.Err) {
	var (
		dbGroup dbGroup
		err     error
	)

	err = r.db.Get(&dbGroup, `
		SELECT * FROM user_groups
		WHERE group_id = ?`,
		groupID)
	if err != nil {
		logger.LogError("Group not found", "helpdesk/repo/mysql", strconv.FormatUint(groupID, 10), err)
		return nil, models.BadRequest("Group not found")
	}
	return r.toModelsGroup(&dbGroup), nil
}

func (r *GroupRepo) GetGroupList() ([]*group_manager.Group, models.Err) {
	return nil, nil
}

func (r *GroupRepo) Close() {
	r.db.Close()
}
