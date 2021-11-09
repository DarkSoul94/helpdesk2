package group_manager

import (
	"github.com/DarkSoul94/helpdesk2/models"
)

type GroupManager interface {
	CheckPermission(groupID uint64, actions ...string) models.Err
	CreateGroup(group *Group) (uint64, models.Err)
	GetGroupByID(groupID uint64) (*Group, models.Err)
	GetGroupList() ([]*Group, models.Err)

	GetPermissionList() []byte
}
