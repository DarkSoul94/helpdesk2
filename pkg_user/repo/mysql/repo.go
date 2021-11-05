package mysql

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/jmoiron/sqlx"
)

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *Repo) CreateUser(user *models.User) (uint64, error) {
	var (
		dbUser dbUser
		query  string
		err    error
	)
	dbUser = r.toDBUser(user)
	if dbUser.ID == 0 {
		query = `
			INSERT INTO users SET
				name = :name,
				email = :email,
				department = :department
				ON DUPLICATE KEY UPDATE 
				name = :name,
				department = :department`
	} else {
		query = `
			UPDATE users SET
				name = :name,
				department = :department
			WHERE id = :id`
	}
	res, err := r.db.NamedExec(query, &dbUser)
	if err != nil {
		logger.LogError(ErrCreateUser.Error(), "user_manager/repo/mysql", dbUser.Email, err)
		return 0, err
	}
	lastID, _ := res.LastInsertId()
	if user.ID == 0 {
		user.ID = uint64(lastID)
	}
	return uint64(lastID), nil

}

func (r *Repo) UserUpdate(userID, groupID uint64) error {
	query := `
		UPDATE users SET
		group_id = ? 
		WHERE id = ?`
	_, err := r.db.Exec(query, groupID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetUserByEmail(email string) (models.User, error) {
	var (
		dbUser dbUser
		query  string
		err    error
	)

	query = `SELECT * FROM users WHERE email = ?`
	err = r.db.Get(&dbUser, query, email)
	if err != nil {
		logger.LogError(ErrReadUser.Error(), "user_manager/repo/mysql", email, err)
		return models.User{}, ErrReadUser
	}

	mUser := r.toModelUser(&dbUser)
	group, err := r.GetGroupByID(dbUser.GroupID)
	mUser.Group = &group
	if err != nil {
		return models.User{}, err
	}

	return mUser, nil
}

func (r *Repo) GetUserByID(id uint64) (models.User, error) {
	var (
		dbUser dbUser
		query  string
		err    error
	)

	query = `SELECT * FROM users WHERE id = ?`
	err = r.db.Get(&dbUser, query, id)
	if err != nil {
		logger.LogError(ErrReadUser.Error(), "user_manager/repo/mysql", fmt.Sprintf("user id: %d", id), err)
		return models.User{}, ErrReadUser
	}

	mUser := r.toModelUser(&dbUser)
	group, err := r.GetGroupByID(dbUser.GroupID)
	mUser.Group = &group
	if err != nil {
		return models.User{}, err
	}

	return mUser, nil
}

func (r *Repo) GetUsersList() ([]models.User, error) {
	var (
		dbUsersList []dbUser
		err         error
	)
	query := `SELECT * FROM users`

	if err = r.db.Select(&dbUsersList, query); err != nil {
		logger.LogError(ErrReadUsersList.Error(), "user_manager/repo/mysql", "", err)
		return []models.User{}, ErrReadUsersList
	}

	mUsersList := make([]models.User, 0)
	for _, val := range dbUsersList {
		mUsersList = append(mUsersList, r.toModelUser(&val))
	}
	return mUsersList, nil
}

func (r *Repo) GetGroupByID(groupID uint64) (models.Group, error) {
	var (
		group dbGroup
		query string
		err   error
	)
	query = `SELECT * FROM user_groups WHERE id = ?`
	err = r.db.Get(&group, query, groupID)
	if err != nil {
		logger.LogError(ErrReadGroup.Error(), "user_manager/repo/mysql", strconv.FormatUint(groupID, 10), err)
		return models.Group{}, ErrReadGroup
	}

	return r.toModelGroup(group), nil
}

func (r *Repo) GetGroupList() ([]models.Group, error) {
	var (
		dbGroupsList []dbGroup
		err          error
	)
	query := `SELECT * FROM user_groups`

	if err = r.db.Select(&dbGroupsList, query); err != nil {
		logger.LogError(ErrReadGroupsList.Error(), "user_manager/repo/mysql", "", err)
		return []models.Group{}, ErrReadGroupsList
	}

	mGroupsList := make([]models.Group, 0)
	for _, val := range dbGroupsList {
		mGroupsList = append(mGroupsList, r.toModelGroup(val))
	}
	return mGroupsList, nil

}

func (r *Repo) Close() error {
	r.db.Close()
	return nil
}
