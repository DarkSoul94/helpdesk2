package pkg_user

import (
	"github.com/DarkSoul94/helpdesk2/models"
)

type IUserUsecase interface {
	CreateUser(user *models.User) (uint64, models.Err)
	UserUpdate(askUser *models.User, userID, groupID uint64) models.Err
	GetUserByEmail(email string) (*models.User, models.Err)
	GetUserByID(id uint64) (*models.User, models.Err)
	GetUsersList(askUser *models.User) ([]*models.User, models.Err)

	GetGroupList(askUser *models.User) ([]*models.Group, models.Err)
	CreateGroup(askUser *models.User, group *models.Group) (uint64, models.Err)
	GroupUpdate(askUser *models.User, group *models.Group) models.Err
}
