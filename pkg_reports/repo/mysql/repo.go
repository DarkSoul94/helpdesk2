package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_reports"
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

func (r *ReportsRepo) GetSupportsShifts(startDate, endDate time.Time) ([]internal_models.SupportsShifts, error) {
	var (
		dbShifts     []dbSupportsShifts
		openingTimes map[string][]internal_models.OpeningDayTime = make(map[string][]internal_models.OpeningDayTime)
		mShifts      []internal_models.SupportsShifts
		query        string
		err          error
	)

	query = `SELECT user_name AS support, opening_time, closing_time, difference AS count_of_minutes_late FROM supports_shifts AS shift
	LEFT JOIN support_lateness AS late ON (
		shift.support_id = late.support_id 
		AND EXTRACT(DAY FROM shift.opening_time) = EXTRACT(DAY FROM late.date) 
		AND EXTRACT(MONTH FROM shift.opening_time) = EXTRACT(MONTH FROM late.date)
	)
	INNER JOIN users ON shift.support_id = user_id
	WHERE opening_time BETWEEN ? AND ?`

	err = r.db.Select(&dbShifts, query, startDate, endDate)
	if err != nil {
		logger.LogError(
			"Failed read supports shifts",
			"pkg_reports/repo/mysql",
			fmt.Sprintf("start date: %s; end date: %s;", startDate, endDate),
			err,
		)

		return nil, err
	}

	for _, shift := range dbShifts {
		openingTime := internal_models.OpeningDayTime{
			OpeningDate: shift.OpeningDate.Local().Format("2006-01-02 15:04:05"),
		}

		if shift.ClosingDate.Valid {
			openingTime.ClosingDate = shift.ClosingDate.Time.Local().Format("2006-01-02 15:04:05")
		} else {
			openingTime.ClosingDate = " "
		}

		if shift.CountOfMinutesLate.Valid {
			openingTime.CountOfMinutesLate = uint64(shift.CountOfMinutesLate.Int64)
		}

		openingTimes[shift.Support] = append(openingTimes[shift.Support], openingTime)
	}

	graceTime := r.GetConstVal(startDate, pkg_reports.GraceTime)

	for support, val := range openingTimes {
		shift := internal_models.SupportsShifts{
			Support:          support,
			WithOutGraceTime: 0,
		}

		for _, dayTime := range val {
			shift.WithOutGraceTime += dayTime.CountOfMinutesLate
			shift.DayTime = append(shift.DayTime, dayTime)
		}

		if shift.WithOutGraceTime > graceTime {
			shift.WithOutGraceTime -= graceTime
		} else {
			shift.WithOutGraceTime = 0
		}
		mShifts = append(mShifts, shift)
	}
	return mShifts, nil
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

func (r *ReportsRepo) GetConstVal(date time.Time, name string) uint64 {
	var (
		val   uint64
		query string
	)

	query = `SELECT val FROM const_change_history
	WHERE name = ? 
	AND EXTRACT(YEAR FROM date) >= EXTRACT(YEAR FROM ?)
	AND EXTRACT(MONTH FROM date) > EXTRACT(MONTH FROM ?)
	LIMIT 1`
	r.db.Get(&val, query, name, date, date)

	if val == 0 {
		query = `SELECT data FROM consts
		WHERE name = ?
		LIMIT 1`
		r.db.Get(&val, query, name)
	}

	return val
}

func (r *ReportsRepo) Close() error {
	return r.db.Close()
}
