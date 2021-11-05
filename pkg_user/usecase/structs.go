package usecase

import (
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	"github.com/DarkSoul94/helpdesk2/pkg_user/perm_manager"
)

type Usecase struct {
	repo        pkg_user.UserManagerRepo
	permManager perm_manager.Manager
}
