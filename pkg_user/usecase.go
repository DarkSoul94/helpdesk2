package pkg_user

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
)

type UserManagerUC interface {
	CreateUser(user *User) (uint64, models.Err)
	UserUpdate(author *User, userID, groupID uint64) models.Err
	GetUserByEmail(email string) (*User, models.Err)
	GetUserByID(id uint64) (*User, models.Err)
	GetUsersList(askUser *User) ([]*User, models.Err)

	GetGroupByID(user *User, groupID uint64) (*group_manager.Group, models.Err)
	GetGroupList(user *User) ([]*group_manager.Group, models.Err)
	GroupUpdate(id uint64, permission []byte) models.Err
	CreateGroup(user *User, group *group_manager.Group) (uint64, models.Err)

	CheckPermissionForAction(user *User, actions ...string) models.Err
}
