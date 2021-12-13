package permissions

import (
	"github.com/DarkSoul94/helpdesk2/models"
)

func checkUserUpdate(group *models.Group) bool {
	return group.ChangeSettings
}

func checkGroupChange(group *models.Group) bool {
	return group.ChangeSettings
}

func checkGroupGet(group *models.Group) bool {
	return group.ChangeSettings
}

func checkIsAdmin(group *models.Group) bool {
	return group.ChangeSettings
}

func checkFullSearch(group *models.Group) bool {
	return group.FullSearch
}

func checkWorkOnTicket(group *models.Group) bool {
	return group.WorkOnTickets
}

func checkResolveTicket(group *models.Group) bool {
	return group.CanResolveTicket
}

func checkSeeAdditionalInfo(group *models.Group) bool {
	return group.SeeAdditionalInfo
}

func checkCreateTicket(group *models.Group) bool {
	return group.CreateTicket
}

func checkReportGet(group *models.Group) bool {
	return group.CanReports
}
