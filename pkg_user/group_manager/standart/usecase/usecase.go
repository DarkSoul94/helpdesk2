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

func (u *Usecase) CheckPermission(user *models.User, actions ...string) models.Err {
	group, err := u.repo.GetGroupByID(user.Group)
	if err != nil {
		return err
	}
	errArray := make([]models.Err, 0)
	for _, action := range actions {
		switch action {
		case global_const.AdminTA_UserUpdate:
			if err := checkUserUpdate(group); err != nil {
				errArray = append(errArray, err)
			}
		case global_const.AdminTA_GroupCreate,
			global_const.AdminTA_GroupUpdate:
			if err := checkGroupChange(group); err != nil {
				errArray = append(errArray, err)
			}
		}
	}
	if len(errArray) > 0 {
		return models.Concat(errArray...)
	}
	return nil
}
