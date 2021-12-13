package mysql

import "github.com/DarkSoul94/helpdesk2/models"

// toModelsGroup
func (r *GroupRepo) toModelsGroup(group *dbGroup) *models.Group {
	return &models.Group{
		ID:                group.ID,
		Name:              group.Name,
		CreateTicket:      group.CreateTicket,
		GetAllTickets:     group.GetAllTickets,
		SeeAdditionalInfo: group.SeeAdditionalInfo,
		CanResolveTicket:  group.CanResolveTicket,
		WorkOnTickets:     group.WorkOnTickets,
		ChangeSettings:    group.ChangeSettings,
		CanReports:        group.CanReports,
		FullSearch:        group.FullSearch,
	}
}

//toDbGroup ...
func (r *GroupRepo) toDbGroup(group *models.Group) *dbGroup {
	return &dbGroup{
		ID:                group.ID,
		Name:              group.Name,
		CreateTicket:      group.CreateTicket,
		GetAllTickets:     group.GetAllTickets,
		SeeAdditionalInfo: group.SeeAdditionalInfo,
		CanResolveTicket:  group.CanResolveTicket,
		WorkOnTickets:     group.WorkOnTickets,
		ChangeSettings:    group.ChangeSettings,
		CanReports:        group.CanReports,
		FullSearch:        group.FullSearch,
	}
}
