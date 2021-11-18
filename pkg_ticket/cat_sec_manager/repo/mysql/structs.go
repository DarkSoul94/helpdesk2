package mysql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type CatSecRepo struct {
	db *sqlx.DB
}

type dbCategory struct {
	ID          uint64          `db:"category_id"`
	Name        string          `db:"category_name"`
	Significant bool            `db:"significant_category"`
	Old         bool            `db:"old_category"`
	Service     bool            `db:"service"`
	Price       decimal.Decimal `db:"price"`
}

type dbCategorySection struct {
	ID             uint64         `db:"section_id"`
	CategoryID     uint64         `db:"category_id"`
	Name           string         `db:"category_section_name"`
	Significant    bool           `db:"significant_category_section"`
	Old            bool           `db:"old_category_section"`
	NeedApproval   bool           `db:"need_approval"`
	Service        bool           `db:"service"`
	Template       sql.NullString `db:"template"`
	ApprovalGroups []uint64
}

type dbSectionWithCategory struct {
	ID           uint64         `db:"section_id"`
	Name         string         `db:"category_section_name"`
	Significant  bool           `db:"significant_category_section"`
	Old          bool           `db:"old_category_section"`
	NeedApproval bool           `db:"need_approval"`
	Service      bool           `db:"service"`
	Template     sql.NullString `db:"template"`
	Category     *dbCategory    `db:""`
}
