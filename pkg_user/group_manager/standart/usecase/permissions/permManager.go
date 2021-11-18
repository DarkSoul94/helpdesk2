package permissions

import (
	"github.com/DarkSoul94/helpdesk2/global_const"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
)

type Usecase struct {
	repo group_manager.IGroupRepo
}

func NewPermManager(repo group_manager.IGroupRepo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) CheckPermission(groupID uint64, actions ...string) models.Err {
	group, err := u.repo.GetGroupByID(groupID)
	if err != nil {
		return err
	}
	errArray := make([]models.Err, 0)
	for _, action := range actions {
		errArray = append(errArray, u.check(group, action))
	}
	return models.Concat(errArray...)
}

func (u *Usecase) CheckUpdatedPermissions(group *models.Group, actions ...string) models.Err {
	errArray := make([]models.Err, 0)
	for _, action := range actions {
		errArray = append(errArray, u.check(group, action))
	}
	return models.Concat(errArray...)
}

func (u *Usecase) check(group *models.Group, action string) models.Err {
	switch action {
	case global_const.AdminTA_UserUpdate:
		return checkUserUpdate(group)

	case global_const.AdminTA_GroupCreate,
		global_const.AdminTA_GroupUpdate:
		return checkGroupChange(group)

	case global_const.AdminTA_GroupGet:
		return checkGroupGet(group)

	case global_const.AdminTA:
		return checkIsAdmin(group)

	case global_const.TicketTA_FullSearch:
		return checkFullSearch(group)

	case global_const.TicketTA_Work:
		return checkWorkOnTicket(group)

	default:
		return nil
	}
}
