package pkg_user

import "github.com/DarkSoul94/helpdesk2/models"

type IUserRepo interface {
	CreateUser(user *models.User) (uint64, models.Err)
	UpdateUser(userID, groupID uint64) models.Err
	GetUserByEmail(email string) (*models.User, models.Err)
	GetUserByID(id uint64) (*models.User, models.Err)
	GetUsersList() ([]*models.User, models.Err)

	Close() error
}
