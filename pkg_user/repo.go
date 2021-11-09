package pkg_user

import "github.com/DarkSoul94/helpdesk2/models"

type UserManagerRepo interface {
	CreateUser(user *User) (uint64, models.Err)
	UpdateUser(userID, groupID uint64) models.Err
	GetUserByEmail(email string) (*User, models.Err)
	GetUserByID(id uint64) (*User, models.Err)
	GetUsersList() ([]*User, models.Err)

	Close() error
}
