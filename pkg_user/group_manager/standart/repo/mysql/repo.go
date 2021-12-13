package mysql

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/jmoiron/sqlx"
)

func NewGroupRepo(db *sql.DB) *GroupRepo {
	return &GroupRepo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

/*CreateGroup - принимает данные новой группы, если такая есть в базе выдает ошибку.
Если нет записывает в базу.*/
func (r *GroupRepo) CreateGroup(group *models.Group) (uint64, models.Err) {
	var err error

	dbGroup := r.toDbGroup(group)
	fmt.Println(group)
	query := `INSERT INTO user_groups SET
							group_name = :group_name,
							create_ticket = :create_ticket,
							get_all_tickets = :get_all_tickets,
							see_additional_info = :see_additional_info,
							can_resolve_ticket = :can_resolve_ticket,
							work_on_tickets = :work_on_tickets,
							change_settings = :change_settings,
							can_reports = :can_reports,
							full_search = :full_search`

	res, err := r.db.NamedExec(query, dbGroup)
	if err != nil {
		logger.LogError("Failed create group", "pkg_user/group_manager/standart/repo/mysql", "", err)
		if strings.Contains(err.Error(), "Duplicate") {
			return 0, GroupErr_Exist
		}
		return 0, GroupErr_Create
	}

	lastID, _ := res.LastInsertId()
	return uint64(lastID), nil
}

func (r *GroupRepo) GetGroupByID(groupID uint64) (*models.Group, models.Err) {
	var (
		dbGroup dbGroup
		err     error
	)
	query := `SELECT * FROM user_groups
						WHERE group_id = ?`

	if err = r.db.Get(&dbGroup, query, groupID); err != nil {
		logger.LogError("Group not found", "pkg_user/group_manager/standart/repo/mysql", strconv.FormatUint(groupID, 10), err)
		return nil, GroupErr_NotFound
	}
	return r.toModelsGroup(&dbGroup), nil
}

func (r *GroupRepo) GetGroupList() ([]*models.Group, models.Err) {
	var dbGroups []dbGroup

	if err := r.db.Select(&dbGroups, `SELECT * FROM user_groups`); err != nil {
		logger.LogError("Failed read groups from DB", "pkg_user/group_manager/standart/repo/mysql", "", err)
		return nil, GroupErr_NotFound
	}

	modelGroups := make([]*models.Group, 0)
	for _, group := range dbGroups {
		modelGroups = append(modelGroups, r.toModelsGroup(&group))
	}

	return modelGroups, nil
}

func (r *GroupRepo) GroupUpdate(group *models.Group) models.Err {
	dbGrp := r.toDbGroup(group)
	query := `UPDATE user_groups SET
							group_name = :group_name,
							create_ticket = :create_ticket,
							get_all_tickets = :get_all_tickets,
							see_additional_info = :see_additional_info,
							can_resolve_ticket = :can_resolve_ticket,
							work_on_tickets = :work_on_tickets,
							change_settings = :change_settings,
							can_reports = :can_reports,
							full_search = :full_search
						WHERE group_id = :group_id`
	if _, err := r.db.NamedExec(query, dbGrp); err != nil {
		logger.LogError("Failed update group", "pkg_user/group_manager/standart/repo/mysql", "", err)
		return GroupErr_Update
	}
	return nil
}

func (r *GroupRepo) GetUsersByGroup(groupID uint64) ([]uint64, models.Err) {
	users := make([]uint64, 0)
	query := `
		SELECT user_id FROM users
		WHERE group_id = ?`

	if err := r.db.Select(&users, query, groupID); err != nil {
		logger.LogError("Group not found", "pkg_user/group_manager/standart/repo/mysql", strconv.FormatUint(groupID, 10), err)
		return nil, GroupErr_NotFound
	}
	return users, nil
}

func (r *GroupRepo) Close() {
	r.db.Close()
}
