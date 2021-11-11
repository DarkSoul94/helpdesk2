package group_manager

import (
	"github.com/DarkSoul94/helpdesk2/models"
)

type GroupManager interface {
	CheckPermission(groupID uint64, actions ...string) models.Err
	CreateGroup(group *models.Group) (uint64, models.Err)
	GetGroupByID(groupID uint64) (*models.Group, models.Err)
	GetGroupList() ([]*models.Group, models.Err)

	GetPermissionList() []byte
}
