package usecase

import (
	"github.com/DarkSoul94/helpdesk2/user_manager"
	"github.com/DarkSoul94/helpdesk2/user_manager/perm_manager"
)

type Usecase struct {
	repo        user_manager.UserManagerRepo
	permManager perm_manager.Manager
}
