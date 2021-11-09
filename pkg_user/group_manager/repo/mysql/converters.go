package mysql

import "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"

func (r *Repo) toModelGroup(group dbGroup) *group_manager.Group {
	return &group_manager.Group{
		ID:                group.GroupID,
		Name:              group.GroupName,
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

func (r *Repo) toDbGroup(group *group_manager.Group) dbGroup {
	return dbGroup{
		GroupID:           group.ID,
		GroupName:         group.Name,
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
