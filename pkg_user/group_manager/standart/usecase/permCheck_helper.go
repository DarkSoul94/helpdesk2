package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
)

func checkUserUpdate(group *models.Group) models.Err {
	if !group.ChangeSettings {
		return models.BadRequest("Insufficient rights for update user")
	}
	return nil
}

func checkGroupChange(group *models.Group) models.Err {
	if !group.ChangeSettings {
		return models.BadRequest("Insufficient rights for change users groups")
	}
	return nil
}

func checkGroupGet(group *models.Group) models.Err {
	if !group.ChangeSettings {
		return models.BadRequest("Insufficient rights for get group")
	}
	return nil
}
