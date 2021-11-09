package pkg_user

import "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"

type User struct {
	ID         uint64
	Email      string
	Name       string
	Group      *group_manager.Group
	Department string
}
