package permissions

import (
	"github.com/DarkSoul94/helpdesk2/global_const/actions"
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

func (u *Usecase) CheckPermission(groupID uint64, actions ...string) bool {
	group, err := u.repo.GetGroupByID(groupID)
	if err != nil {
		return false
	}
	for _, action := range actions {
		if !u.check(group, action) {
			return false
		}
	}
	return true
}

func (u *Usecase) CheckUpdatedPermissions(group *models.Group, actions ...string) bool {
	for _, action := range actions {
		if !u.check(group, action) {
			return false
		}
	}
	return true
}

func (u *Usecase) check(group *models.Group, action string) bool {
	switch action {
	case actions.AdminTA_UserUpdate:
		return checkUserUpdate(group)

	case actions.AdminTA_GroupCreate,
		actions.AdminTA_GroupUpdate:
		return checkGroupChange(group)

	case actions.AdminTA_GroupGet:
		return checkGroupGet(group)

	case actions.AdminTA:
		return checkIsAdmin(group)

	case actions.TicketTA_FullSearch:
		return checkFullSearch(group)

	case actions.TicketTA_Work:
		return checkWorkOnTicket(group)

	case actions.TicketTA_Resolve:
		return checkResolveTicket(group)

	default:
		return false
	}
}
