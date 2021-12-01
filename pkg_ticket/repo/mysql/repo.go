package mysql

import (
	"database/sql"
	"errors"
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
	need_resolve = :need_resolve,
	ticket_text = :ticket_text,
	ticket_status_id = :ticket_status_id,
	ticket_author_id = :ticket_author_id,
	support_id = :support_id,
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

func (r *TicketRepo) UpdateTicket(ticket *internal_models.Ticket) error {
	var (
		query string
		err   error
	)

	query = `UPDATE tickets SET
				section_id = :section_id,
				need_resolve = :need_resolve,
				ticket_status_id = :ticket_status_id,
				support_id = :support_id,
				resolved_user_id = :resolved_user_id,
				service_comment = :service_comment
				WHERE ticket_id = :ticket_id`

	_, err = r.db.NamedExec(query, r.toDbTicket(ticket))
	if err != nil {
		logger.LogError(
			"Failed update ticket",
			"pkg_ticket/repo/mysql",
			fmt.Sprintf("ticket id: %d; section id: %d; status id: %d; support id: %d; service comment: %s", ticket.ID, ticket.CatSect.ID, ticket.Status.ID, ticket.Support.ID, ticket.ServiceComment),
			err,
		)

		return err
	}

	return nil
}

func (r *TicketRepo) GetLastTicketStatusHistory(ticketID uint64) (*internal_models.TicketStatusHistory, error) {
	var (
		history dbTicketStatusHistory
		query   string
		err     error
	)

	query = `SELECT TH.*, TS.ticket_status_id, TS.ticket_status_name FROM ticket_status_history AS TH
				INNER JOIN ticket_status AS TS ON TS.ticket_status_id = TH.ticket_status_id 
				WHERE TH.ticket_id = ?
				ORDER BY TH.id DESC LIMIT 1`

	err = r.db.Get(&history, query, ticketID)
	if err != nil {
		logger.LogError(
			"Failed read ticket status history",
			"pkg_ticket/repo/mysql",
			fmt.Sprintf("ticket id: %d;", ticketID),
			err,
		)
		return nil, err
	}

	return r.toModelTicketStatusHistory(history), nil
}

func (r *TicketRepo) CreateTicketStatusHistory(history *internal_models.TicketStatusHistory) error {
	var (
		query string
		err   error
	)

	query = `INSERT INTO ticket_status_history SET
		ticket_id = :ticket_id,
		changed_user_id = :changed_user_id,
		select_time = :select_time,
		ticket_status_id = :ticket_status_id,
		duration = :duration`

	_, err = r.db.NamedExec(query, r.toDbTicketStatusHistory(history))
	if err != nil {
		logger.LogError(
			"Failed create ticket status history",
			"pkg_ticket/repo/mysql",
			fmt.Sprintf("ticket id: %d; changed user id: %d; select time: %s; status id: %d; duration: %d;", history.TicketId, history.ChangedUser.ID, history.SelectTime, history.Status.ID, history.Duration),
			err,
		)
		return err
	}

	return nil
}

func (r *TicketRepo) UpdateTicketStatusHistory(history *internal_models.TicketStatusHistory) error {
	var (
		query string
		err   error
	)

	query = `UPDATE ticket_status_history SET
				duration = :duration
				WHERE id = :id`

	_, err = r.db.NamedExec(query, r.toDbTicketStatusHistory(history))
	if err != nil {
		logger.LogError(
			"Failed update ticket status history",
			"pkg_ticket/repo/mysql",
			fmt.Sprintf("id: %d;", history.ID),
			err,
		)
		return err
	}

	return nil
}

func (r *TicketRepo) GetAllTicketStatusHistory(ticketID uint64) ([]*internal_models.TicketStatusHistory, error) {
	var (
		historyList  []dbTicketStatusHistory
		mHistoryList []*internal_models.TicketStatusHistory
		query        string
		err          error
	)

	query = `SELECT TH.*, TS.ticket_status_id, TS.ticket_status_name FROM ticket_status_history AS TH
				INNER JOIN ticket_status AS TS ON TS.ticket_status_id = TH.ticket_status_id 
				WHERE TH.ticket_id = ?
				ORDER BY TH.id`

	err = r.db.Select(&historyList, query, ticketID)
	if err != nil {
		logger.LogError(
			"Failed read all ticket status history",
			"pkg_ticket/repo/mysql",
			fmt.Sprintf("ticket id: %d;", ticketID),
			err,
		)

		return nil, err
	}

	for _, history := range historyList {
		mHistoryList = append(mHistoryList, r.toModelTicketStatusHistory(history))
	}

	return mHistoryList, nil
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
						WHERE ticket_status_id IN (8, 9)
						AND CAST(select_time AS DATE) = CURRENT_DATE
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
					INNER JOIN ticket_status AS TS ON TS.ticket_status_id = T.ticket_status_id
					WHERE T.need_resolve = true
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

func (r *TicketRepo) GetTicketListForDistribute() ([]*internal_models.Ticket, error) {
	var (
		dbList []dbTicket
		query  string
		err    error
	)

	query = `SELECT T.* FROM tickets AS T
				JOIN category_section USING(section_id)
				JOIN category USING(category_id)
				WHERE ticket_status_id = ?
				ORDER BY support_id DESC, significant_category DESC, significant_category_section DESC, ticket_date`

	err = r.db.Select(&dbList, query, internal_models.TSWaitID)
	if err != nil {
		return nil, err
	}

	return r.convertTicketList(dbList), nil
}

func (r *TicketRepo) GetTicketListForReturnToDistribute() ([]*internal_models.Ticket, error) {
	var (
		dbList []dbTicket
		query  string
		err    error
	)

	query = `SELECT T.* FROM tickets AS T
				WHERE EXISTS (SELECT * FROM supports_activity AS SA WHERE T.ticket_id = SA.ticket_id AND SA.reassignment = true)`

	err = r.db.Select(&dbList, query)
	if err != nil {
		return nil, err
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

func (r *TicketRepo) GetTicket(ticketID uint64) (*internal_models.Ticket, error) {
	var (
		ticket dbTicket
		query  string
		err    error
	)

	query = `SELECT T.*, TS.ticket_status_id, TS.ticket_status_name FROM tickets AS T
				INNER JOIN ticket_status AS TS ON TS.ticket_status_id = T.ticket_status_id
				WHERE ticket_id = ?`

	err = r.db.Get(&ticket, query, ticketID)
	if err != nil {
		logger.LogError(
			"Failed get ticket",
			"pkg_ticket/repo/mysql",
			fmt.Sprintf("ticket id: %d;", ticketID),
			err,
		)
		return nil, err
	}

	return r.toModelTicket(ticket), nil
}

func (r *TicketRepo) StealTicket(ticketID, supportID uint64, toWork bool) error {
	var (
		query string
		err   error
	)

	if toWork {
		query = `UPDATE tickets SET
					support_id = ?,
					ticket_status_id = 4
					WHERE ticket_id = ?`
	} else {
		query = `UPDATE tickets SET
					support_id = ?
					WHERE ticket_id = ?`
	}

	_, err = r.db.Exec(query, supportID, ticketID)
	if err != nil {
		logger.LogError(
			"Failed steal ticket",
			"pkg_ticket/repo/mysql",
			fmt.Sprintf("ticket id: %d; support id: %d;", ticketID, supportID),
			err,
		)
		return err
	}

	return nil
}

func (r *TicketRepo) TicketGrade(ticketID, userID uint64, grade uint) error {
	var (
		res   sql.Result
		query string
		err   error
	)

	query = `UPDATE tickets SET 
				ticket_grade = ?
				WHERE ticket_id = ? 
				AND ticket_status_id = 9 
				AND support_id IS NOT NULL
				AND ticket_author_id = ?`

	res, err = r.db.Exec(query, grade, ticketID, userID)
	if err != nil {
		logger.LogError(
			"Failed ticket grade",
			"pkg_ticket/repo/mysql",
			fmt.Sprintf("ticket id: %d; grade: %d; user id: %d;", ticketID, grade, userID),
			err,
		)
		return err
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		return errors.New("Ошибка оценки запроса")
	}

	return nil
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
					WHERE T.need_resolve = true
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

func (r *TicketRepo) GetTicketsCount(supportID, statusID uint64) int {
	var count int
	if supportID != 0 {
		query := getCountWithSupp()
		if err := r.db.Get(&count, query, supportID, statusID); err != nil {
			logger.LogError(
				"Failed read tickets count",
				"pkg_ticket/repo/mysql",
				fmt.Sprintf("support id: %d, status id: %d", supportID, statusID),
				err)
		}
		return count
	}
	query := getCountWithoutSupp()
	if err := r.db.Get(&count, query, statusID); err != nil {
		logger.LogError(
			"Failed read tickets count",
			"pkg_ticket/repo/mysql",
			fmt.Sprintf("support id: %d, status id: %d", supportID, statusID),
			err)
	}
	return count
}

func (r *TicketRepo) GetTodayTicketsCount(supportID, statusID uint64) int {
	var count int
	if supportID != 0 {
		query := getTodaysCountWithSupp()
		if err := r.db.Get(&count, query, supportID, statusID); err != nil {
			logger.LogError(
				"Failed read tickets count",
				"pkg_ticket/repo/mysql",
				fmt.Sprintf("support id: %d, status id: %d", supportID, statusID),
				err)
		}
		return count
	}

	query := getTodaysCountWithoutSupp()
	if err := r.db.Get(&count, query, statusID); err != nil {
		logger.LogError(
			"Failed read tickets count",
			"pkg_ticket/repo/mysql",
			fmt.Sprintf("support id: %d, status id: %d", supportID, statusID),
			err)
	}
	return count
}

func (r *TicketRepo) Close() error {
	r.db.Close()
	return nil
}
