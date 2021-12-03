package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"
	"github.com/jmoiron/sqlx"
)

func NewReportsRepo(db *sql.DB) *ReportsRepo {
	return &ReportsRepo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *ReportsRepo) GetAverageGradesBySupport(startDate, endDate time.Time) (map[string]float64, error) {
	var (
		dbAVG []dbAverageGrade
		mAVG  map[string]float64 = make(map[string]float64)
		query string
		err   error
	)

	query = `SELECT user_name, ROUND(AVG(ticket_grade), 2) AS average_grade FROM tickets
				INNER JOIN users ON user_id = support_id
				WHERE ticket_status_id = 9 
				AND	support_id IS NOT NULL 
				AND ticket_grade IS NOT NULL
				AND ticket_date BETWEEN ? AND ?
				GROUP BY support_id
				ORDER BY support_id`

	err = r.db.Select(&dbAVG, query, startDate, endDate)
	if err != nil {
		logger.LogError(
			"Failed read avarage grade",
			"pkg_reports/repo/mysql",
			fmt.Sprintf("start date: %s; end date: %s;", startDate, endDate),
			err,
		)
		return nil, err
	}

	for _, grade := range dbAVG {
		mAVG[grade.SupportName] = grade.AverageGrade
	}

	return mAVG, nil
}

func (r *ReportsRepo) GetSupportsStatusHistory(startDate, endDate time.Time) (map[string][]internal_models.SupportStatusHistory, error) {
	var (
		dbHistoryList []dbSupportStatusHistory
		mHistory      map[string][]internal_models.SupportStatusHistory = make(map[string][]internal_models.SupportStatusHistory)
		query         string
		err           error
	)

	query = `SELECT user_name AS support_name, support_status_name AS status_name, select_time, duration FROM support_status_history
				LEFT JOIN users ON user_id = support_id
				LEFT JOIN support_status ON support_status_id = status_id
				WHERE shift_id IS NOT NULL
				AND EXISTS (SELECT * FROM supports_shifts
								WHERE id = shift_id
								AND opening_time BETWEEN ? AND ?
							)
				ORDER BY support_id, select_time`

	err = r.db.Select(&dbHistoryList, query, startDate, endDate)
	if err != nil {
		logger.LogError(
			"Failed read supports status history",
			"pkg_reports/repo/mysql",
			fmt.Sprintf("start date: %s; end date: %s;", startDate, endDate),
			err,
		)
		return nil, err
	}

	for _, history := range dbHistoryList {
		mHistory[history.SupportName] = append(mHistory[history.SupportName], internal_models.SupportStatusHistory{
			StatusName: history.StatusName,
			SelectTime: history.SelectTime,
			Duration:   history.Duration,
		})
	}

	return mHistory, nil
}

func (r *ReportsRepo) Close() error {
	return r.db.Close()
}
