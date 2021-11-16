package mysql

import (
	"database/sql"
	"fmt"

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

//GetSupport получение объекта суппорта по ID пользователя. Возвращает ошибку если такой саппорт не найден
func (r *Repo) GetSupport(userID uint64) (*internal_models.Support, models.Err) {
	dbSupp := new(dbSupport)
	query := `
	SELECT * FROM support
	WHERE support_id = ?`
	if err := r.db.Get(dbSupp, query, userID); err != nil {
		logger.LogError("Failed get support", "pkg_support/repo/mysql", fmt.Sprintf("support id: %d", userID), err)
		return nil, errSupportGet
	}
	return r.toModelSupport(dbSupp), nil
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
	if _, err := r.db.NamedExec(query, dbCard); err != nil {
		logger.LogError("Failed create support card", "pkg_support/repo/mysql", fmt.Sprintf("support id: %d", dbCard.SupportID), err)
		return errCardCreate
	}
	return nil
}

//DeleteCard удаляет карточку саппорта из БД
func (r *Repo) DeleteCard(supportID uint64) models.Err {
	query := `
	DELETE FROM support
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

func (r *Repo) Close() error {
	r.db.Close()
	return nil
}
