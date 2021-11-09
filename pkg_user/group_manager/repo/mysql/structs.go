package mysql

import "github.com/jmoiron/sqlx"

type Repo struct {
	db *sqlx.DB
}

type dbGroup struct {
	GroupID           uint64 `db:"group_id"`
	GroupName         string `db:"group_name"`
	CreateTicket      bool   `db:"create_ticket"`
	GetAllTickets     bool   `db:"get_all_tickets"`
	SeeAdditionalInfo bool   `db:"see_additional_info"`
	CanResolveTicket  bool   `db:"can_resolve_ticket"`
	WorkOnTickets     bool   `db:"work_on_tickets"`
	ChangeSettings    bool   `db:"change_settings"`
	CanReports        bool   `db:"can_reports"`
	FullSearch        bool   `db:"full_search"`
}
