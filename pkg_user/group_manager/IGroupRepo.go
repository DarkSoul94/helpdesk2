package group_manager

import "github.com/DarkSoul94/helpdesk2/models"

type IGroupRepo interface {
	CreateGroup(group *models.Group) (uint64, models.Err)
	GroupUpdate(group *models.Group) models.Err
	GetGroupByID(groupID uint64) (*models.Group, models.Err)
	GetGroupList() ([]*models.Group, models.Err)
	GetUsersByGroup(groupID uint64) ([]uint64, models.Err)

	Close()
}
