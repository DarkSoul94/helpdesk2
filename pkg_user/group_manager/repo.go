package group_manager

import "github.com/DarkSoul94/helpdesk2/models"

type GroupRepo interface {
	GetGroupByID(groupID uint64) (*Group, models.Err)
	GetGroupList() ([]*Group, models.Err)

	Close()
}
