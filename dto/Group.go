package dto

import (
	"github.com/DarkSoul94/helpdesk2/models"
)

type OutGroup struct {
	ID                uint64 `json:"group_id,omitempty"`
	Name              string `json:"group_name"`
	CreateTicket      bool   `json:"create_ticket"`
	GetAllTickets     bool   `json:"get_all_tickets"`
	SeeAdditionalInfo bool   `json:"see_additional_info"`
	CanResolveTicket  bool   `json:"can_resolve_ticket"`
	WorkOnTickets     bool   `json:"work_on_tickets"`
	ChangeSettings    bool   `json:"change_settings"`
	CanReports        bool   `json:"can_reports"`
	FullSearch        bool   `json:"full_search"`
}

func ToModelGroup(g OutGroup) *models.Group {
	return &models.Group{
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

func ToOutGroup(group *models.Group) OutGroup {
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

func ToOutGroupList(groups []*models.Group) []OutGroup {
	var outGroups []OutGroup = make([]OutGroup, 0)

	for _, group := range groups {
		outGroups = append(outGroups, ToOutGroup(group))
	}

	return outGroups
}
