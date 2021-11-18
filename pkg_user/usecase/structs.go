package usecase

import (
	"github.com/DarkSoul94/helpdesk2/pkg_support"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
)

type Usecase struct {
	repo    pkg_user.IUserRepo
	group   group_manager.IGroupUsecase
	perm    group_manager.IPermManager
	support pkg_support.ISupportUsecase
}
