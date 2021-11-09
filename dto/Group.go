package dto

import "github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"

type OutGroup struct {
	ID                uint64 `json:"group_id,omitempty"`
	Name              string `json:"group_name"`
	CreateTicket      bool   `json:"create_ticket,omitempty"`
	GetAllTickets     bool   `json:"get_all_tickets,omitempty"`
	SeeAdditionalInfo bool   `json:"see_additional_info,omitempty"`
	CanResolveTicket  bool   `json:"can_resolve_ticket,omitempty"`
	WorkOnTickets     bool   `json:"work_on_tickets,omitempty"`
	ChangeSettings    bool   `json:"change_settings,omitempty"`
	CanReports        bool   `json:"can_reports,omitempty"`
	FullSearch        bool   `json:"full_search,omitempty"`
}

func ToModelGroup(g OutGroup) *group_manager.Group {
	return &group_manager.Group{
		ID:                g.ID,
		Name:              g.Name,
		CreateTicket:      g.CreateTicket,
		GetAllTickets:     g.GetAllTickets,
		SeeAdditionalInfo: g.SeeAdditionalInfo,
		CanResolveTicket:  g.CanResolveTicket,
		WorkOnTickets:     g.WorkOnTickets,
		ChangeSettings:    g.ChangeSettings,
		CanReports:        g.CanReports,
		FullSearch:        g.FullSearch,
	}
}

func ToOutGroup(group *group_manager.Group) OutGroup {
	if group != nil {
		return OutGroup{
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
	return OutGroup{}
}

func ToOutGroupList(groups []*group_manager.Group) []OutGroup {
	var outGroups []OutGroup = make([]OutGroup, 0)

	for _, group := range groups {
		outGroups = append(outGroups, OutGroup{ID: group.ID, Name: group.Name})
	}

	return outGroups
}
