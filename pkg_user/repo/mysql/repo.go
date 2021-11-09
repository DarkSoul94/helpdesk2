package mysql

import (
	"database/sql"
	"fmt"

	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	"github.com/jmoiron/sqlx"
)

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *Repo) CreateUser(user *pkg_user.User) (uint64, error) {
	var (
		query string
		err   error
	)
	if user.ID == 0 {
		query = `
			INSERT INTO users SET
				user_name = :user_name,
				email = :email,
				department = :department
				ON DUPLICATE KEY UPDATE 
				user_name = :user_name,
				department = :department`
	} else {
		query = `
			UPDATE users SET
				user_name = :user_name,
				department = :department
			WHERE user_id = :user_id`
	}
	res, err := r.db.NamedExec(query, r.toDbUser(user))
	if err != nil {
		logger.LogError(ErrCreateUser.Error(), "user_manager/repo/mysql", user.Email, err)
		return 0, err
	}
	lastID, _ := res.LastInsertId()
	if user.ID == 0 {
		user.ID = uint64(lastID)
	}

	return uint64(lastID), nil
}

func (r *Repo) UpdateUser(userID, groupID uint64) error {
	query := `
		UPDATE users SET
		group_id = ? 
		WHERE user_id = ?`

	_, err := r.db.Exec(query, groupID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetUserByEmail(email string) (*pkg_user.User, error) {
	var (
		dbUser dbUser
		query  string
		err    error
	)

	query = `SELECT * FROM users WHERE email = ?`
	err = r.db.Get(&dbUser, query, email)
	if err != nil {
		logger.LogError(ErrReadUser.Error(), "user_manager/repo/mysql", email, err)
		return &pkg_user.User{}, ErrReadUser
	}

	return r.toModelUser(dbUser), nil
}

func (r *Repo) GetUserByID(id uint64) (*pkg_user.User, error) {
	var (
		dbUser dbUser
		query  string
		err    error
	)

	query = `SELECT * FROM users WHERE id = ?`

	err = r.db.Get(&dbUser, query, id)
	if err != nil {
		logger.LogError(ErrReadUser.Error(), "user_manager/repo/mysql", fmt.Sprintf("user id: %d", id), err)
		return &pkg_user.User{}, ErrReadUser
	}

	return r.toModelUser(dbUser), nil
}

func (r *Repo) GetUsersList() ([]*pkg_user.User, error) {
	var (
		dbUsersList []dbUser
		err         error
	)

	query := `SELECT * FROM users`

	if err = r.db.Select(&dbUsersList, query); err != nil {
		logger.LogError(ErrReadUsersList.Error(), "user_manager/repo/mysql", "", err)
		return []*pkg_user.User{}, ErrReadUsersList
	}

	mUsersList := make([]*pkg_user.User, 0)
	for _, val := range dbUsersList {
		mUsersList = append(mUsersList, r.toModelUser(val))
	}

	return mUsersList, nil
}

func (r *Repo) Close() error {
	r.db.Close()
	return nil
}
