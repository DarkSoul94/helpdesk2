package mysql

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
	"github.com/jmoiron/sqlx"
)

func NewSupportRepo(db *sql.DB) *Repo {
	return &Repo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

//CreateSupport предназначена для создания новой записи сотрудника ТП
func (r *Repo) CreateSupport(support *internal_models.Support) models.Err {
	dbSupp := r.toDbSupport(support)
	query := `
		INSERT INTO support SET 
			support_id = :support_id,
			status_id = :status_id,
			priority = :priority`
	if _, err := r.db.NamedExec(query, dbSupp); err != nil {
		logger.LogError("Failed create support", "pkg_support/repo/mysql", fmt.Sprintf("support id: %d", dbSupp.SupportID), err)
		return errSupportCreate
	}
	return nil
}

//DeleteSupport удаление запси суппорта из БД по ID пользователя
func (r *Repo) DeleteSupport(userID uint64) models.Err {
	query := `
		DELETE FROM support
		WHERE support_id = ?`
	if _, err := r.db.Exec(query, userID); err != nil {
		logger.LogError("Failed delete support", "pkg_support/repo/mysql", fmt.Sprintf("support id: %d", userID), err)
		return errSupportDelete
	}
	return nil
}

//UpdateSupport обновление статуса и приоритета распределения саппорта
func (r *Repo) UpdateSupport(support *internal_models.Support) models.Err {
	dbSupp := r.toDbSupport(support)
	query := `
	UPDATE support SET
		status_id = :status_id,
		priority = :priority
	WHERE support_id = :support_id`
	if _, err := r.db.NamedExec(query, dbSupp); err != nil {
		logger.LogError("Failed update support", "pkg_support/repo/mysql", fmt.Sprintf("support id: %d", dbSupp.SupportID), err)
		return errSupportUpdate
	}
	return nil
}

//GetSupport получение объекта суппорта по ID пользователя. Возвращает ошибку если такой саппорт не найден
func (r *Repo) GetSupport(userID uint64) (*internal_models.Support, models.Err) {
	dbSupp := new(dbSupport)
	query := `
		SELECT 
			support.*, 
			users.user_name AS support_name 
		FROM support 
		LEFT JOIN users ON user_id = support_id
		WHERE support_id = ?`
	if err := r.db.Get(dbSupp, query, userID); err != nil {
		logger.LogError("Failed get support", "pkg_support/repo/mysql", fmt.Sprintf("support id: %d", userID), err)
		return nil, errSupportGet
	}
	return r.toModelSupport(dbSupp), nil
}

//GetSupportList получить полный список саппортов
func (r *Repo) GetSupportList() ([]*internal_models.Support, models.Err) {
	dbSupp := make([]dbSupport, 0)
	mSupports := make([]*internal_models.Support, 0)
	query := `
		SELECT 
			support.*, 
			users.user_name AS support_name 
		FROM support 
		LEFT JOIN users ON user_id = support_id`
	if err := r.db.Select(&dbSupp, query); err != nil {
		logger.LogError("Failed get support list", "pkg_support/repo/mysql", "", err)
		return nil, errSupportGet
	}

	for _, support := range dbSupp {
		mSupports = append(mSupports, r.toModelSupport(&support))
	}
	return mSupports, nil
}

//GetActiveSupports получить список активных саппортов
func (r *Repo) GetActiveSupports() ([]*internal_models.Support, models.Err) {
	dbSupp := make([]dbSupport, 0)
	mSupports := make([]*internal_models.Support, 0)
	query := `
		SELECT 
			support.*, 
			users.user_name AS support_name 
		FROM support 
		LEFT JOIN users 
			ON user_id = support_id
		WHERE EXISTS (
			SELECT * FROM support_status
			WHERE support_status_id = status_id
				AND accept_ticket = true
		)
		ORDER BY support_id`
	if err := r.db.Select(&dbSupp, query); err != nil {
		logger.LogError("Failed get support list", "pkg_support/repo/mysql", "", err)
		return nil, errSupportGetActive
	}
	for _, support := range dbSupp {
		mSupports = append(mSupports, r.toModelSupport(&support))
	}
	return mSupports, nil
}

func (r *Repo) GetRandomFreeSupport() (*internal_models.Support, models.Err) {
	dbSupp := new(dbSupport)
	query := `
		SELECT * FROM support
		WHERE EXISTS (
			SELECT * FROM support_status
			WHERE support_status_id = status_id
				AND accept_ticket = true
		) AND NOT EXIST (
			SELECT * FROM supports_activity
			WHERE supports_activity.support_id = support.support_id
		) ORDER BY RAND() LIMIT 1`
	if err := r.db.Get(dbSupp, query); err != nil {
		logger.LogError("Failed get support status", "pkg_support/repo/mysql", "", err)
		return nil, errSupportGetRandom
	}
	return r.toModelSupport(dbSupp), nil
}

//GetPrioritizedSupportID возвращает ID саппорта у которого установлен приоритет распределения
func (r *Repo) GetPrioritizedSupportID() uint64 {
	var id uint64
	query := `
		SELECT support_id
		FROM support
		WHERE priority = true
		LIMIT 1`
	if err := r.db.Get(&id, query); err != nil {
		return 0
	}
	return id
}

func (r *Repo) CheckForActivity(supportID uint64) bool {
	var active bool
	query := `
		SELECT EXISTS (
			SELECT * FROM support
			RIGHT JOIN support_status AS Stat 
				ON support.status_id = Stat.support_status_id 
			WHERE support.support_id = ?
				AND Stat.accept_ticket = true
		)`
	r.db.Get(&active, query, supportID)
	return active
}

func (r *Repo) CheckForBusy(supportID uint64) bool {
	var active bool
	query := `
		SELECT EXISTS (
			SELECT * FROM supports_activity
			WHERE support_id = ?
		)`
	r.db.Get(&active, query, supportID)
	return active
}

func (r *Repo) CreateSupportActivity(supportID, ticketID uint64) models.Err {
	query := `
	INSERT INTO supports_activity SET
		support_id = ?,
		ticket_id = ?`
	if _, err := r.db.Exec(query, supportID, ticketID); err != nil {
		logger.LogError("Failed create support activity", "pkg_support/repo/mysql", fmt.Sprintf("support id: %d, ticket id: %d", supportID, ticketID), err)
		return errSupportModifyActivity
	}
	return nil
}

func (r *Repo) RemoveSupportActivity(ticketID uint64) models.Err {
	query := `
	DELETE FROM support_activity
		WHERE ticket_id = ?`
	if _, err := r.db.Exec(query, ticketID); err != nil {
		logger.LogError("Failed remove support activity", "pkg_support/repo/mysql", fmt.Sprintf("ticket id: %d", ticketID), err)
		return errSupportModifyActivity
	}
	return nil
}

func (r *Repo) UpdateSupportActivity(supportID, ticketID uint64) models.Err {
	query := `
	UPDATE supports_activity SET
		support_id = ?
	WHERE ticket_id = ?`
	if _, err := r.db.Exec(query, supportID, ticketID); err != nil {
		logger.LogError("Failed update support activity", "pkg_support/repo/mysql", fmt.Sprintf("support id: %d, ticket id: %d", supportID, ticketID), err)
		return errSupportModifyActivity
	}
	return nil
}

//GetStatus получить статус саппорта по ID статуса
func (r *Repo) GetStatus(statusID uint64) (*internal_models.Status, models.Err) {
	status := new(dbStatus)
	query := `SELECT * FROM support_status WHERE support_status_id = ?`
	if err := r.db.Get(status, query, statusID); err != nil {
		logger.LogError("Failed get support status", "pkg_support/repo/mysql", fmt.Sprintf("status id: %d", statusID), err)
		return nil, errStatusGet
	}
	return r.toModelsStatus(status), nil
}

//GetStatusesList получить список возможных статусов саппортов
func (r *Repo) GetStatusesList() ([]*internal_models.Status, models.Err) {
	dbStat := make([]dbStatus, 0)
	mStat := make([]*internal_models.Status, 0)
	query := `SELECT * FROM support_status`
	if err := r.db.Select(&dbStat, query); err != nil {
		logger.LogError("Failed get support statuses list", "pkg_support/repo/mysql", "", err)
		return nil, errSupportGet
	}
	for _, stat := range dbStat {
		mStat = append(mStat, r.toModelsStatus(&stat))
	}
	return mStat, nil
}

func (r *Repo) UpdateShift(shift *internal_models.Shift) (uint64, models.Err) {
	var query string
	dbShift := r.toDbShift(shift)

	if dbShift.ID == 0 {
		query = `INSERT INTO supports_shifts SET
		support_id = :support_id,
		opening_time = :opening_time`
	} else {
		query = `UPDATE supports_shifts SET
		support_id = :support_id,
		closing_time = :closing_time,
		closing_status = :closing_status
		WHERE id = :id`
	}
	res, err := r.db.NamedExec(query, dbShift)
	if err != nil {
		logger.LogError("Failed insert shift changes to db", "helpdesk/repo/mysql", strconv.FormatUint(shift.Support.ID, 10), err)
		return 0, errShiftUpdateShift
	}
	id, _ := res.LastInsertId()
	return uint64(id), nil
}

//GetLastShift получить последнюю смену саппорта
func (r *Repo) GetLastShift(supportID uint64) (*internal_models.Shift, models.Err) {
	shift := new(dbShift)
	query := `
		SELECT * 
		FROM supports_shifts
		WHERE support_id = ?
		ORDER BY id DESC LIMIT 1`
	if err := r.db.Get(shift, query, supportID); err != nil {
		logger.LogError("Failed get support last shift", "pkg_support/repo/mysql", fmt.Sprintf("support id: %d", supportID), err)
		return nil, errShiftGet
	}
	return r.toModelShift(shift), nil
}

//CreateHistoryRecord создает запись в истории изменения статусов сотрудников
func (r *Repo) CreateHistoryRecord(statHistory *internal_models.StatusHistory) models.Err {
	dbHistory := r.toDbStatusHistory(statHistory)
	query := `
		INSERT INTO support_status_history SET
			support_id = :support_id,
			select_time = :select_time,
			status_id = :status_id,
			duration = :duration,
			shift_id = :shift_id`
	if _, err := r.db.NamedExec(query, dbHistory); err != nil {
		logger.LogError("Failed create record to support status history", "pkg_support/repo/mysql",
			fmt.Sprintf("support id: %d, status id: %d, shift id: %d", dbHistory.SupportID, dbHistory.StatusID, dbHistory.ShiftID), err)
		return errHistoryCreate
	}
	return nil
}

//UpdateHistoryRecord внесение длительности нахождения саппорта в указанном статусе
func (r *Repo) UpdateHistoryRecord(statHistory *internal_models.StatusHistory) models.Err {
	dbHistory := r.toDbStatusHistory(statHistory)
	query := `
		UPDATE support_status_history SET
			duration = :duration
		WHERE id = :id`
	if _, err := r.db.NamedExec(query, dbHistory); err != nil {
		logger.LogError("Failed update support status history record", "pkg_support/repo/mysql",
			fmt.Sprintf("support id: %d, record id: %d", dbHistory.SupportID, dbHistory.ID), err)
		return errCardCreate
	}
	return nil
}

//GetLastStatusHistory получить последнюю запись из истории статусов саппортов
func (r *Repo) GetLastStatusHistory(supportID, shiftID uint64) (*internal_models.StatusHistory, models.Err) {
	dbHistory := new(dbStatusHistory)
	query := `
		SELECT * FROM support_status_history
		WHERE support_id = ?
			AND shift_id = ?
		ORDER BY id DESC
		LIMIT 1`
	if err := r.db.Get(dbHistory, query, supportID, shiftID); err != nil {
		logger.LogError("Failed get support status history record", "pkg_support/repo/mysql",
			fmt.Sprintf("support id: %d, shift id: %d", supportID, shiftID), err)
		return nil, errHistoryGet
	}
	return r.toModelsStatusHistory(dbHistory), nil
}

//CreateSupportCard создает новую запись карточки суппорта.
func (r *Repo) CreateCard(card *internal_models.Card) models.Err {
	dbCard := r.toDbSupportCard(card)
	query := `
		INSERT INTO supports_cards SET 
			support_id = :support_id,
			internal_number = :internal_number,
			mobile_number = :mobile_number,
			birth_date = :birth_date,
			is_senior = :is_senior,
			senior_id = :senior_id,
			wager = :wager,
			comment = :comment,
			color = :color`
	if _, err := r.db.NamedExec(query, &dbCard); err != nil {
		logger.LogError("Failed create support card", "pkg_support/repo/mysql", fmt.Sprintf("support id: %d", dbCard.SupportID), err)
		return errCardCreate
	}
	return nil
}

//DeleteCard удаляет карточку саппорта из БД
func (r *Repo) DeleteCard(supportID uint64) models.Err {
	query := `
	DELETE FROM supports_cards
	WHERE support_id = ?`
	if _, err := r.db.Exec(query, supportID); err != nil {
		logger.LogError("Failed delete support card", "pkg_support/repo/mysql", fmt.Sprintf("support id: %d", supportID), err)
		return errCardDelete
	}
	return nil
}

//GetCardBySupportID получение карточки суппорта по ID саппорта
func (r *Repo) GetCardBySupportID(supportID uint64) (*internal_models.Card, models.Err) {
	card := new(dbCard)
	query := `
	SELECT * FROM supports_cards 
	WHERE support_id = ?`
	if err := r.db.Get(card, query, supportID); err != nil {
		logger.LogError("Failed delete support card", "pkg_support/repo/mysql", fmt.Sprintf("support id: %d", supportID), err)
		return nil, errCardGet
	}
	return r.toModelSupportCard(card), nil
}

//ResetSenior во всех записях саппортов, где ID старшего равно переданному, очищает поле с ID старшего и устанавливает цвет на белый
func (r *Repo) ResetSenior(seniorID uint64) models.Err {
	query := `
	UPDATE supports_cards SET
		support_id = NULL,
		color = "#FFFFFF"
	WHERE senior_id = ?`
	if _, err := r.db.Exec(query, seniorID); err != nil {
		logger.LogError("Failed reset support cards", "pkg_support/repo/mysql", fmt.Sprintf("senior id: %d", seniorID), err)
		return errCardDelete
	}
	return nil
}

func (r *Repo) GetSupportListForToday() ([]*internal_models.Support, models.Err) {
	list := make([]dbSupport, 0)
	mList := make([]*internal_models.Support, 0)
	query := `
		SELECT support.*, users.user_name AS support_name FROM support
		RIGHT JOIN users ON users.user_id = support.support_id
		WHERE EXISTS (
			SELECT * FROM supports_shifts
			WHERE supports_shifts.support_id = support.support_id
			AND (
				CAST(opening_time AS DATE) = CURRENT_DATE
				OR closing_status = false
				) 
		)`
	if err := r.db.Select(&list, query); err != nil {
		logger.LogError("Failed get support list", "pkg_support/repo/mysql", "", err)
		return nil, errSupportGetList
	}
	for _, val := range list {
		mList = append(mList, r.toModelSupport(&val))
	}
	return mList, nil
}

func (r *Repo) Close() error {
	r.db.Close()
	return nil
}
