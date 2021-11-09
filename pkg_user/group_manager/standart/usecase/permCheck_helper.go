package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
)

func checkUserUpdate(group *group_manager.Group) models.Err {
	if !group.ChangeSettings {
		return models.BadRequest("Insufficient rights for update user")
	}
	return nil
}

func checkGroupChange(group *group_manager.Group) models.Err {
	if !group.ChangeSettings {
		return models.BadRequest("Insufficient rights for change users groups")
	}
	return nil
}

func checkGroupGet(group *group_manager.Group) models.Err {
	if !group.ChangeSettings {
		return models.BadRequest("Insufficient rights for get group")
	}
	return nil
}
