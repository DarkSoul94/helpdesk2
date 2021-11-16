package group_manager

import "github.com/DarkSoul94/helpdesk2/models"

type IPermManager interface {
	CheckPermission(groupID uint64, actions ...string) models.Err
	CheckUpdatedPermissions(group *models.Group, actions ...string) models.Err
}
