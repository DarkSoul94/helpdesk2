package group

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
)

type Usecase struct {
	repo group_manager.IGroupRepo
}

func NewGroupManager(repo group_manager.IGroupRepo) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) CreateGroup(group *models.Group) (uint64, models.Err) {
	return u.repo.CreateGroup(group)
}

func (u *Usecase) GetGroupByID(groupID uint64) (*models.Group, models.Err) {
	return u.repo.GetGroupByID(groupID)
}

func (u *Usecase) GetGroupList() ([]*models.Group, models.Err) {
	return u.repo.GetGroupList()
}

func (u *Usecase) GetUsersByGroup(groupID uint64) ([]uint64, models.Err) {
	return u.repo.GetUsersByGroup(groupID)
}

func (u *Usecase) GroupUpdate(group *models.Group) models.Err {
	return u.repo.GroupUpdate(group)
}
