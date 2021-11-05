package models

import "github.com/DarkSoul94/helpdesk2/pkg_user/perm_manager"

type Group struct {
	ID          uint64
	Name        string
	Permissions perm_manager.PermLayer
}
