package usecase

import (
	"github.com/DarkSoul94/helpdesk2/global_const"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
)

type Usecase struct {
	repo group_manager.GroupRepo
}

func NewManager(repo group_manager.GroupRepo) (*Usecase, error) {
	return &Usecase{repo: repo}, nil
}

func (u *Usecase) GetPermissionList() []byte {
	return nil
}

func (u *Usecase) CheckPermission(groupID uint64, actions ...string) models.Err {
	group, err := u.repo.GetGroupByID(groupID)
	if err != nil {
		return err
	}
	errArray := make([]models.Err, 0)
	for _, action := range actions {
		switch action {
		case global_const.AdminTA_UserUpdate:
			errArray = append(errArray, checkUserUpdate(group))
		case global_const.AdminTA_GroupCreate,
			global_const.AdminTA_GroupUpdate:
			errArray = append(errArray, checkGroupChange(group))
		case global_const.AdminTA_GroupGet:
			errArray = append(errArray, checkGroupGet(group))
		}
	}
	return models.Concat(errArray...)
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
