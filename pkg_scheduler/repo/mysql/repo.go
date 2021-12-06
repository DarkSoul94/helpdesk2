package mysql

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_scheduler/internal_models"
	"github.com/jmoiron/sqlx"
)

type SchedulerRepo struct {
	db *sqlx.DB
}

func NewShedulerRepo(db *sql.DB) *SchedulerRepo {
	return &SchedulerRepo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *SchedulerRepo) AddOffice(office *internal_models.Office) models.Err {
	var (
		query string
		err   error
	)

	query = `INSERT INTO offices SET
	name = :name,
	color = :color,
	deleted = :deleted`

	if _, err = r.db.NamedExec(query, r.toDbOffice(office)); err != nil {
		logger.LogError("Failed add office to db", "pkg_scheduler/repo/mysql", "office name: "+office.Name, err)
		return models.InternalError("Не удалось добавить новый офис в базу")
	}
	return nil
}

func (r *SchedulerRepo) UpdateOffice(office *internal_models.Office) models.Err {
	var (
		query string
		err   error
	)

	query = `UPDATE offices SET
	name = :name,
	color = :color,
	deleted = :deleted
	WHERE id = :id`

	if _, err = r.db.NamedExec(query, r.toDbOffice(office)); err != nil {
		logger.LogError("Failed update office to db", "pkg_scheduler/repo/mysql", "Shift name: "+office.Name, err)
		return models.InternalError("Не удалось изменить указанный офис")
	}

	return nil

}

func (r *SchedulerRepo) GetOfficeByID(officeID uint64) (*internal_models.Office, models.Err) {
	var (
		office dbOffice
		query  string
		err    error
	)

	query = `SELECT * FROM offices
				WHERE id = ?`
	if err = r.db.Get(&office, query, officeID); err != nil {
		logger.LogError("Failed read office from db", "pkg_scheduler/repo/mysql", "Office id: "+strconv.FormatUint(officeID, 10), err)
		return nil, models.InternalError("Не удалось получить указанный офис")
	}

	return r.toModelOffice(office), nil
}

func (r *SchedulerRepo) GetOfficesList(deleted bool, dates ...string) ([]*internal_models.Office, models.Err) {
	var (
		offices  []dbOffice
		mOffices []*internal_models.Office
		query    string
		err      error
	)

	query = `SELECT * FROM offices WHERE deleted = ?`
	if len(dates) != 0 {
		query = query + `
		AND id IN (SELECT office_id FROM shifts_schedule
			WHERE EXTRACT(MONTH FROM date) = EXTRACT(MONTH FROM ?)
			AND EXTRACT(YEAR FROM date) = EXTRACT(YEAR FROM ?))`
		err = r.db.Select(&offices, query, deleted, dates[0], dates[0])
	} else {
		err = r.db.Select(&offices, query, deleted)
	}
	if err != nil {
		logger.LogError("Failed read office from db", "scheduleManager/repo/mysql", "", err)
		return nil, models.InternalError("Не удалось получить список офисов")
	}

	for _, office := range offices {
		mOffices = append(mOffices, r.toModelOffice(office))
	}

	return mOffices, nil
}

func (r *SchedulerRepo) UpdateCell(cell *internal_models.Cell) models.Err {
	var (
		query string
		err   error
	)
	query = `INSERT INTO shifts_schedule SET
		support_id = :support_id,
		office_id = :office_id,
		start_time = :start_time,
		end_time = :end_time,
		date = :date,
		vacation = :vacation,
		sick_leave = :sick_leave
	ON DUPLICATE KEY UPDATE 
		office_id = :office_id,
		start_time = :start_time,
		end_time = :end_time,
		date = :date,
		vacation = :vacation,
		sick_leave = :sick_leave`

	res, err := r.db.NamedExec(query, r.toDbShiftsScheduleCell(cell))
	if err != nil {
		logger.LogError(
			"Failed add shifts schedule cell to db",
			"pkg_scheduler/repo/mysql",
			fmt.Sprintf("support_id: %d; date: %s", cell.SupportID, cell.Date),
			err,
		)
		return models.InternalError("Не удалось обновить график смен")
	}
	if id, _ := res.LastInsertId(); id != 0 {
		cell.ID = uint64(id)
	}

	return nil
}

func (r *SchedulerRepo) DeleteCells(actualCellsIDs map[string][]uint64) models.Err {
	var (
		query string
		err   error
	)
	for date, val := range actualCellsIDs {
		args := []interface{}{
			date,
			date,
		}
		for _, id := range val {
			args = append(args, id)
		}
		query = `DELETE FROM shifts_schedule 
		WHERE EXTRACT(MONTH FROM date) = EXTRACT(MONTH FROM ?)
		AND EXTRACT(YEAR FROM date) = EXTRACT(YEAR FROM ?)
	  AND id NOT IN(?` + strings.Repeat(`,?`, len(val)-1) + `)`
		if _, err = r.db.Exec(query, args...); err != nil {
			logger.LogError(
				"Failed drop shifts schedule cells",
				"pkg_scheduler/repo/mysql",
				"",
				err,
			)
			return models.InternalError("Не удалось обновить график смен")
		}
	}
	return nil
}

func (r *SchedulerRepo) GetSchedule(date string) ([]*internal_models.Cell, models.Err) {
	var (
		dbShiftsSchedule []dbCell
		mShiftsSchedule  []*internal_models.Cell
		query            string
		err              error
	)

	query = `SELECT * FROM shifts_schedule
				WHERE EXTRACT(MONTH FROM date) = EXTRACT(MONTH FROM ?)
				AND EXTRACT(YEAR FROM date) = EXTRACT(YEAR FROM ?)`

	err = r.db.Select(&dbShiftsSchedule, query, date, date)
	if err != nil {
		logger.LogError("Failed read shifts schedule from db", "pkg_scheduler/repo/mysql", date, err)
		return nil, models.InternalError("Не удалось получить график смен")
	}

	for _, cell := range dbShiftsSchedule {
		mShiftsSchedule = append(mShiftsSchedule, r.toModelShiftsScheduleCell(cell))
	}

	return mShiftsSchedule, nil
}

func (r *SchedulerRepo) CreateLateness(lateness *internal_models.Lateness) models.Err {
	var (
		err   error
		query string
	)
	query = `
		INSERT INTO support_lateness SET
			date = :date,
			support_id = :support_id,
			cause = :cause,
			difference = :difference`
	dbLate := r.toDbLateness(lateness)
	_, err = r.db.NamedExec(query, &dbLate)
	if err != nil {
		logger.LogError("Failed create support lateness",
			"pkg_scheduler/repo/mysql",
			fmt.Sprintf("support id: %d", lateness.SupportID),
			err)
		return models.InternalError("Не удалось создать опоздание")
	}
	return nil
}

func (r *SchedulerRepo) GetLateness(date string) ([]*internal_models.Lateness, models.Err) {
	var (
		err          error
		query        string
		latenessList = make([]dbLateness, 0)
		mList        = make([]*internal_models.Lateness, 0)
	)
	query = `
		SELECT Late.*, User.user_name AS support_name FROM support_lateness AS Late
		LEFT JOIN users AS User ON user_id = support_id
		WHERE EXTRACT(MONTH FROM date) = EXTRACT(MONTH FROM ?)
					AND EXTRACT(YEAR FROM date) = EXTRACT(YEAR FROM ?)`
	err = r.db.Select(&latenessList, query, date, date)
	if err != nil {
		logger.LogError("Failed get support lateness", "pkg_scheduler/repo/mysql", fmt.Sprintf("date: %s", date), err)
		return nil, models.InternalError("Не удалось получить опоздания за месяц")
	}
	for _, record := range latenessList {
		mList = append(mList, r.toModelSupportLateness(&record))
	}
	return mList, nil
}

func (r *SchedulerRepo) GetLatenessByID(latenessID uint64) (*internal_models.Lateness, models.Err) {
	var (
		err        error
		query      string
		dbLateness dbLateness
	)
	query = `
		SELECT 
			support_lateness.*, 
			users.user_name AS support_name 
		FROM support_lateness
		LEFT JOIN users ON support_id = user_id
		WHERE id = ?`
	err = r.db.Get(&dbLateness, query, latenessID)
	if err != nil {
		logger.LogError("Failed get support lateness", "pkg_scheduler/repo/mysql",
			fmt.Sprintf("lateness ID: %d", latenessID),
			err)
		return nil, models.InternalError("Не удалось получить опоздание")
	}
	return r.toModelSupportLateness(&dbLateness), nil
}

func (r *SchedulerRepo) UpdateLateness(lateness *internal_models.Lateness) models.Err {
	var (
		err   error
		query string
	)
	dbLateness := r.toDbLateness(lateness)
	query = `
		UPDATE support_lateness SET
			decision = :decision
		WHERE id = :id`
	_, err = r.db.NamedExec(query, dbLateness)
	if err != nil {
		logger.LogError("Failed set decision to suppote lateness", "pkg_scheduler/repo/mysql",
			fmt.Sprintf("lateness ID: %d", lateness.ID),
			err)
		return models.InternalError("Не удалось обновить решение по указанному опозданию")
	}
	return nil
}

func (r *SchedulerRepo) CheckNewLateness() bool {
	var (
		result bool
		query  string
	)
	query = `
	SELECT EXISTS (
		SELECT id FROM support_lateness
		WHERE decision IS NULL
	)`
	err := r.db.Get(&result, query)
	if err != nil {
		logger.LogError("Failed get lateness without decision", "pkg_scheduler/repo/mysql", "", err)
	}
	return result
}

func (r *SchedulerRepo) GetCurrentLatenessConf() (*internal_models.Config, models.Err) {
	var (
		conf                   internal_models.Config
		LatePenalty, GraceTime dbConst
		query                  string
		err                    error
	)

	query = `SELECT * FROM consts WHERE name = ?`
	if err = r.db.Get(&LatePenalty, query, internal_models.ConfKey_Penalty); err != nil {
		logger.LogError("Failed read const from db", "pkg_scheduler/repo/mysql", internal_models.ConfKey_Penalty, err)
	}
	if err = r.db.Get(&GraceTime, query, internal_models.ConfKey_GraceTime); err != nil {
		logger.LogError("Failed read const from db", "pkg_scheduler/repo/mysql", internal_models.ConfKey_GraceTime, err)
	}
	LatePenalty.FromConst(conf.Penalty)
	GraceTime.FromConst(conf.GraceTime)
	return &conf, nil
}

func (r *SchedulerRepo) ConstHistory(name string, data interface{}) models.Err {
	history := createConstHistory(name, data)
	query := `
	INSERT INTO const_change_history SET
		date = :date,
		name = :name,
		val = :val,
		val_type = :val_type
	ON DUPLICATE KEY UPDATE
		val = :val,
		val_type := val_type`

	if _, err := r.db.NamedExec(query, history); err != nil {
		logger.LogError("Failed write const change history to db", "pkg_scheduler/repo/mysql", name, err)
		return models.InternalError("Не удалось записать историю изменения настроек графика")
	}
	return nil
}

func (r *SchedulerRepo) SetLatenessConf(conf *internal_models.Config) models.Err {

	return nil
}

func (r *SchedulerRepo) Close() {
	r.db.Close()
}
