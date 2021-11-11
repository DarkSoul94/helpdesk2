package usecase

import (
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
	"github.com/DarkSoul94/helpdesk2/pkg_user/support_manager"
)

type Usecase struct {
	repo    pkg_user.UserManagerRepo
	group   group_manager.GroupManager
	support support_manager.SupportUsecase
}
