package group_manager

import (
	"github.com/DarkSoul94/helpdesk2/models"
)

type IGroupUsecase interface {
	CreateGroup(group *models.Group) (uint64, models.Err)
	GetGroupByID(groupID uint64) (*models.Group, models.Err)
	GetGroupList() ([]*models.Group, models.Err)
	GroupUpdate(group *models.Group) models.Err
	GetUsersByGroup(groupID uint64) ([]uint64, models.Err)
}
