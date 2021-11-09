package group_manager

import "github.com/DarkSoul94/helpdesk2/models"

type GroupManager interface {
	GetGroupByID(groupID uint64) (*Group, models.Err)
	GetGroupList() ([]*Group, models.Err)

	GetPermissionList() []byte
}
