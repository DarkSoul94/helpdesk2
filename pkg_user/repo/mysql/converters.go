package mysql

import (
	"github.com/DarkSoul94/helpdesk2/models"
)

func (r *Repo) toModelUser(user dbUser) *models.User {
	mUser := &models.User{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		Group: &models.Group{ID: user.GroupID},
	}

	if user.Department.Valid {
		mUser.Department = user.Department.String
	}

	return mUser
}

func (r *Repo) toDbUser(user *models.User) dbUser {
	dbUser := dbUser{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	if user.Group != nil {
		dbUser.GroupID = user.Group.ID
	}

	if len(user.Department) > 0 {
		dbUser.Department.Valid = true
		dbUser.Department.String = user.Department
	} else {
		dbUser.Department.Valid = false
	}

	return dbUser
}
