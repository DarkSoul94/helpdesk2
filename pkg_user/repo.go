package pkg_user

import "github.com/DarkSoul94/helpdesk2/models"

type UserManagerRepo interface {
	CreateUser(user *models.User) (uint64, error)
	UserUpdate(userID, groupID uint64) error
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id uint64) (models.User, error)
	GetUsersList() ([]models.User, error)

	GetGroupByID(groupID uint64) (models.Group, error)
	GetGroupList() ([]models.Group, error)

	Close() error
}
