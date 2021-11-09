package pkg_user

import "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"

type UserManagerUC interface {
	CreateUser(user *User) (uint64, error)
	UserUpdate(author *User, userID, groupID uint64) error
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id uint64) (*User, error)
	GetUsersList(askUser *User) ([]*User, error)

	GetGroupByID(user *User, groupID uint64) (*group_manager.Group, error)
	GetGroupList(user *User) ([]*group_manager.Group, error)
	GroupUpdate(id uint64, permission []byte) error
	CreateGroup(name string, permissions []byte) (uint64, error)

	CheckPermissionForAction(user *User, actions ...string) bool
}
