package mysql

import "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"

// toModelsGroup
func (r *GroupRepo) toModelsGroup(group *dbGroup) *group_manager.Group {
	return &group_manager.Group{
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
func (r *GroupRepo) toDbGroup(group *group_manager.Group) *dbGroup {
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
