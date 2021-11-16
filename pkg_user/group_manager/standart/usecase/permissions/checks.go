package permissions

import (
	"github.com/DarkSoul94/helpdesk2/models"
)

func checkUserUpdate(group *models.Group) models.Err {
	if !group.ChangeSettings {
		return models.Forbidden("Insufficient rights to update user")
	}
	return nil
}

func checkGroupChange(group *models.Group) models.Err {
	if !group.ChangeSettings {
		return models.Forbidden("Insufficient rights to change users groups")
	}
	return nil
}

func checkGroupGet(group *models.Group) models.Err {
	if !group.ChangeSettings {
		return models.Forbidden("Insufficient rights to get group")
	}
	return nil
}

func checkIsAdmin(group *models.Group) models.Err {
	if !group.ChangeSettings {
		return models.Forbidden("Insufficient rights")
	}
	return nil
}

func checkFullSearch(group *models.Group) models.Err {
	if !group.FullSearch {
		return models.Forbidden("Insufficient rights to perform a full search")
	}
	return nil
}

func checkWorkOnTicket(group *models.Group) models.Err {
	if !group.WorkOnTickets {
		return models.Forbidden("Insufficient rights to work on ticket")
	}
	return nil
}
