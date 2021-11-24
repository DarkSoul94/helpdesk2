package mysql

func getTodaysCountWithSupp() string {
	query := `
	SELECT COUNT(*) FROM tickets
	WHERE support_id = ?
		AND ticket_status_id = ?
		AND ticket_status_id = (
	    SELECT ticket_status_id FROM ticket_status_history 
	    WHERE id = (
				SELECT MAX(id) 
				FROM ticket_status_history 
				WHERE ticket_status_history.ticket_id = tickets.ticket_id
				) 
			AND CAST(ticket_status_history.select_time AS DATE) = CURRENT_DATE
	  )`
	return query
}

func getTodaysCountWithoutSupp() string {
	query := `
	SELECT COUNT(*) FROM tickets
	WHERE ticket_status_id = ?
		AND ticket_status_id = (
	    SELECT ticket_status_id FROM ticket_status_history 
	    WHERE id = (
	        SELECT MAX(id) 
	        FROM ticket_status_history 
	        WHERE ticket_status_history.ticket_id = tickets.ticket_id
					) 
	      AND CAST(ticket_status_history.select_time AS DATE) = CURRENT_DATE
	  )`
	return query
}

func getCountWithSupp() string {
	query := `
	SELECT COUNT(*) FROM tickets
	WHERE support_id = ?
	AND ticket_status_id = ?`
	return query
}

func getCountWithoutSupp() string {
	query := `
	SELECT COUNT(*) FROM tickets
	WHERE ticket_status_id = ?`
	return query
}
