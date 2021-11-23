package mysql

import (
	"database/sql"
	"fmt"

	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/jmoiron/sqlx"
)

func NewTicketRepo(db *sql.DB) *TicketRepo {
	return &TicketRepo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *TicketRepo) GetTicketStatuses() ([]*internal_models.TicketStatus, error) {
	var (
		dbList []dbTicketStatus
		mList  []*internal_models.TicketStatus
		query  string
		err    error
	)

	query = `SELECT * FROM ticket_status`

	err = r.db.Select(&dbList, query)
	if err != nil {
		logger.LogError(
			"Failed read ticket status list",
			"pkg_ticket/repo/mysql",
			"",
			err,
		)
		return nil, err
	}

	for _, stat := range dbList {
		mList = append(mList, r.toModelTicketStatus(stat))
	}

	return mList, nil
}

func (r *TicketRepo) GetTicketStatusesSortPriority(isSupport bool) map[uint]uint {
	type statusPriority struct {
		StatusID uint `db:"ticket_status_id"`
		Priority uint `db:"sort_priority"`
	}

	var (
		list     []statusPriority
		priority map[uint]uint = make(map[uint]uint)
		query    string
		err      error
	)

	if isSupport {
		query = `SELECT ticket_status_id, sort_priority_supp AS sort_priority FROM ticket_status`
	} else {
		query = `SELECT ticket_status_id, sort_priority_user AS sort_priority FROM ticket_status`
	}

	err = r.db.Select(&list, query)
	if err != nil {
		logger.LogError(
			"Failed read ticket status sort priority",
			"pkg_ticket/repo/mysql",
			"",
			err,
		)
	}

	for _, val := range list {
		priority[val.StatusID] = val.Priority
	}

	return priority
}

func (r *TicketRepo) CreateTicket(ticket *internal_models.Ticket) (uint64, error) {
	var (
		res   sql.Result
		query string
		err   error
	)

	query = `INSERT INTO tickets SET
	ticket_date = :ticket_date,
	section_id = :section_id,
	ticket_text = :ticket_text,
	ticket_status_id = :ticket_status_id,
	ticket_author_id = :ticket_author_id,
	filial = :filial,
	ip = :ip`

	res, err = r.db.NamedExec(query, r.toDbTicket(ticket))
	if err != nil {
		logger.LogError(
			"Failed create ticket",
			"pkg_ticket/repo/mysql",
			fmt.Sprintf("section id: %d; text: %s; status id: %d; author id: %d; filial: %s; ip: %s", ticket.CatSect.ID, ticket.Text, ticket.Status.ID, ticket.Author.ID, ticket.Filial, ticket.IP),
			err,
		)
		return 0, err
	}

	id, _ := res.LastInsertId()

	return uint64(id), nil
}

func (r *TicketRepo) GetTicketListForAdmin(limit, offset int) ([]*internal_models.Ticket, error) {
	var (
		dbList []dbTicket
		query  string
		err    error
	)

	query = `SELECT T.*, TS.ticket_status_id, TS.ticket_status_name FROM tickets AS T
				INNER JOIN ticket_status AS TS ON TS.ticket_status_id = T.ticket_status_id
				ORDER BY ticket_id
				DESC LIMIT ? OFFSET ?`

	err = r.db.Select(&dbList, query, limit, offset)
	if err != nil {
		logger.LogError(
			"Failed read ticket list for admin",
			"pkg_ticket/repo/mysql",
			"",
			err,
		)
		return nil, err
	}

	if len(dbList) == 0 {
		return []*internal_models.Ticket{}, nil
	}

	return r.convertTicketList(dbList), nil
}

func (r *TicketRepo) GetTicketListForSupport(supportID uint64, limit, offset int) ([]*internal_models.Ticket, error) {
	var (
		dbList []dbTicket
		query  string
		err    error
	)

	query = `SELECT T.*, TS.ticket_status_id, TS.ticket_status_name FROM tickets AS T
				INNER JOIN ticket_status AS TS ON TS.ticket_status_id = T.ticket_status_id 
				WHERE support_id = ? 
				ORDER BY TS.sort_priority_supp, T.ticket_id
				DESC LIMIT ? OFFSET ?`

	err = r.db.Select(&dbList, query, supportID, limit, offset)
	if err != nil {
		logger.LogError(
			"Failed read ticket list for support",
			"pkg_ticket/repo/mysql",
			fmt.Sprintf("user id: %d", supportID),
			err,
		)
		return nil, err
	}

	if len(dbList) == 0 {
		return []*internal_models.Ticket{}, nil
	}

	return r.convertTicketList(dbList), nil
}

func (r *TicketRepo) GetTicketListForUser(authorID uint64, limit, offset int) ([]*internal_models.Ticket, error) {
	var (
		dbList []dbTicket
		query  string
		err    error
	)

	query = `SELECT T.*, TS.ticket_status_id, TS.ticket_status_name FROM tickets AS T
				INNER JOIN ticket_status AS TS ON TS.ticket_status_id = T.ticket_status_id 
				WHERE T.ticket_author_id = ?
				AND (
					T.ticket_status_id NOT IN(8, 9)
					OR T.ticket_id IN (
						SELECT ticket_id FROM ticket_status_history
						WHERE curr_status_id IN (8, 9)
						AND CAST(curr_status_time AS DATE) = CURRENT_DATE
					)
				)
				ORDER BY TS.sort_priority_user, T.ticket_id
				DESC LIMIT ? OFFSET ?`

	err = r.db.Select(&dbList, query, authorID, limit, offset)
	if err != nil {
		logger.LogError(
			"Failed read ticket list for regular user",
			"pkg_ticket/repo/mysql",
			fmt.Sprintf("user id: %d", authorID),
			err,
		)
		return nil, err
	}

	if len(dbList) == 0 {
		return []*internal_models.Ticket{}, nil
	}

	return r.convertTicketList(dbList), nil
}

func (r *TicketRepo) GetTicketListForApproval(groupID uint64, limit, offset int, forResolver bool) ([]*internal_models.Ticket, error) {
	var (
		dbList []dbTicket
		query  string
		err    error
	)

	if forResolver {
		query = `SELECT T.*, TS.ticket_status_id, TS.ticket_status_name FROM tickets AS T
					INNER JOIN ticket_status AS TS ON TS.ticket_status_id = T.ticket_status_id
					WHERE resolved_user_id IS NULL
					AND EXISTS (
						SELECT * FROM approval_bindings
						WHERE group_id = ?
						AND T.section_id = approval_bindings.section_id
					) AND T.ticket_status_id != 8
					ORDER BY T.ticket_id
					DESC LIMIT ? OFFSET ?`

		err = r.db.Select(&dbList, query, groupID, limit, offset)
		if err != nil {
			logger.LogError(
				"Failed read ticket list for resolver",
				"pkg_ticket/repo/mysql",
				fmt.Sprintf("group id: %d", groupID),
				err,
			)
			return nil, err
		}
	} else {
		query = `SELECT T.*, TS.ticket_status_id, TS.ticket_status_name FROM tickets AS T
					INNER JOIN category_section AS CS ON CS.section_id = T.section_id
					INNER JOIN ticket_status AS TS ON TS.ticket_status_id = T.ticket_status_id
					WHERE resolved_user_id IS NULL
					AND CS.need_approval = true
					AND T.ticket_status_id NOT IN(8, 9)
					ORDER BY T.ticket_id
					DESC LIMIT ? OFFSET ?`

		err = r.db.Select(&dbList, query, limit, offset)
		if err != nil {
			logger.LogError(
				"Failed read ticket list for support/admin",
				"pkg_ticket/repo/mysql",
				fmt.Sprintf("group id: %d", groupID),
				err,
			)
			return nil, err
		}
	}

	if len(dbList) == 0 {
		return []*internal_models.Ticket{}, nil
	}

	return r.convertTicketList(dbList), nil
}

func (r *TicketRepo) convertTicketList(dbList []dbTicket) []*internal_models.Ticket {
	var (
		mList []*internal_models.Ticket
	)

	for _, dbTick := range dbList {
		mList = append(mList, r.toModelTicket(dbTick))
	}

	return mList
}

func (r *TicketRepo) CheckNeedApprovalTicketExist(groupID uint64, forResolver bool) (bool, error) {
	var (
		count int
		query string
		err   error
	)

	if forResolver {
		query = `SELECT COUNT(*) FROM tickets 
		WHERE resolved_user_id IS NULL
		AND EXISTS (
			SELECT * FROM approval_bindings
			WHERE group_id = ?
			AND tickets.section_id = approval_bindings.section_id
		) AND ticket_status_id != 8`

		err = r.db.Get(&count, query, groupID)
		if err != nil {
			logger.LogError(
				"Failed check need approval ticket exists for resolver",
				"pkg_ticket/repo/mysql",
				fmt.Sprintf("group id: %d", groupID),
				err,
			)
			return false, err
		}
	} else {
		query = `SELECT COUNT(*) FROM tickets AS T
					INNER JOIN category_section AS CS ON CS.section_id = T.section_id
					WHERE resolved_user_id IS NULL
					AND CS.need_approval = true
					AND T.ticket_status_id NOT IN(8, 9)`
		err = r.db.Get(&count, query)
		if err != nil {
			logger.LogError(
				"Failed check need approval ticket exists for support/admin",
				"pkg_ticket/repo/mysql",
				"",
				err,
			)
			return false, err
		}
	}

	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (r *TicketRepo) Close() error {
	r.db.Close()
	return nil
}
