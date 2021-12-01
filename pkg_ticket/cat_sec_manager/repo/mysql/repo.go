package mysql

import (
	"database/sql"
	"fmt"

	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/jmoiron/sqlx"
)

func NewCatSecRepo(db *sql.DB) *CatSecRepo {
	return &CatSecRepo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *CatSecRepo) CheckCategoryExist(id uint64, name string) bool {
	var (
		db_id uint64
		query string
		err   error
	)

	if id != 0 {
		query = `SELECT category_id FROM category
			WHERE category_id = ?`

		err = r.db.Get(&db_id, query, id)
		if err != nil {
			return false
		}
	} else {
		query = `SELECT category_id FROM category
			WHERE category_name = ?`
		err = r.db.Get(&db_id, query, name)
		if err != nil {
			return false
		}
	}

	if db_id == 0 {
		return false
	}

	return true
}

func (r *CatSecRepo) CreateCategory(cat *internal_models.Category) (uint64, error) {
	var (
		res   sql.Result
		query string
		err   error
	)

	query = `INSERT INTO category SET
				category_name = :category_name,
				significant_category = :significant_category,
				old_category = :old_category,
				price = :price`

	res, err = r.db.NamedExec(query, r.toDbCategory(cat))
	if err != nil {
		logger.LogError(
			"Failed create category",
			"pkg_ticket/cat_sec_manager/repo/mysql",
			fmt.Sprintf("name: %s; significant: %t; old: %t; price: %d", cat.Name, cat.Significant, cat.Significant, cat.Price),
			err,
		)
		return 0, err
	}

	id, _ := res.LastInsertId()

	return uint64(id), nil
}

func (r *CatSecRepo) UpdateCategory(cat *internal_models.Category) error {
	var (
		query string
		err   error
	)

	query = `UPDATE category SET
				category_name = :category_name,
				significant_category = :significant_category,
				old_category = :old_category,
				price = :price
				WHERE category_id = :category_id`

	_, err = r.db.NamedExec(query, r.toDbCategory(cat))
	if err != nil {
		logger.LogError(
			"Failed update category",
			"pkg_ticket/cat_sec_manager/repo/mysql",
			fmt.Sprintf("id: %d; name: %s; significant: %t; old: %t; price: %d", cat.ID, cat.Name, cat.Significant, cat.Significant, cat.Price),
			err,
		)
		return err
	}

	return nil
}

func (r *CatSecRepo) CheckCategorySectionExist(id, cat_id uint64, name string) bool {
	var (
		db_id uint64
		query string
		err   error
	)

	if id != 0 {
		query = `SELECT category_id FROM category_section
			WHERE section_id = ?`

		err = r.db.Get(&db_id, query, id)
		if err != nil {
			return false
		}
	} else {
		query = `SELECT category_id FROM category_section
			WHERE category_section_name = ? AND category_id = ?`
		err = r.db.Get(&db_id, query, name, cat_id)
		if err != nil {
			return false
		}
	}

	if db_id == 0 {
		return false
	}

	return true
}

func (r *CatSecRepo) updateApprovalBindings(sec_id uint64, groups_id []uint64) {
	var (
		query string
	)

	query = `INSERT INTO approval_bindings SET
				section_id = ?,
				group_id = ?`

	for _, id := range groups_id {
		r.db.Exec(query, sec_id, id)
	}

	query = `DELETE FROM approval_bindings 
				WHERE section_id = ? 
				AND group_id NOT IN(?)`
	q, arg, _ := sqlx.In(query, sec_id, groups_id)
	r.db.Exec(q, arg...)
}

func (r *CatSecRepo) CreateCategorySection(sec *internal_models.CategorySection) (uint64, error) {
	var (
		res   sql.Result
		query string
		err   error
	)

	query = `INSERT INTO category_section SET
				category_id = :category_id,
				category_section_name = :category_section_name,
				significant_category_section = :significant_category_section,
				old_category_section = :old_category_section,
				need_approval = :need_approval,
				template = :template`

	res, err = r.db.NamedExec(query, r.toDbCategorySection(sec))
	if err != nil {
		logger.LogError(
			"Failed create category section",
			"pkg_ticket/cat_sec_manager/repo/mysql",
			fmt.Sprintf("category_id: %d; name: %s; significant: %t; old: %t; need_approval: %t; template: %s", sec.CategoryID, sec.Name, sec.Significant, sec.Old, sec.NeedApproval, sec.Template),
			err,
		)
		return 0, err
	}

	id, _ := res.LastInsertId()

	r.updateApprovalBindings(uint64(id), sec.ApprovalGroups)

	return uint64(id), nil
}

func (r *CatSecRepo) UpdateCategorySection(sec *internal_models.CategorySection) error {
	var (
		query string
		err   error
	)

	query = `UPDATE category_section SET
				category_id = :category_id,
				category_section_name = :category_section_name,
				significant_category_section = :significant_category_section,
				old_category_section = :old_category_section,
				need_approval = :need_approval,
				template = :template
				WHERE section_id = :section_id`

	_, err = r.db.NamedExec(query, r.toDbCategorySection(sec))
	if err != nil {
		logger.LogError(
			"Failed update category section",
			"pkg_ticket/cat_sec_manager/repo/mysql",
			fmt.Sprintf("id: %d; category_id: %d; name: %s; significant: %t; old: %t; need_approval: %t; template: %s", sec.ID, sec.CategoryID, sec.Name, sec.Significant, sec.Old, sec.NeedApproval, sec.Template),
			err,
		)
		return err
	}

	r.updateApprovalBindings(sec.ID, sec.ApprovalGroups)

	return nil
}

func (r *CatSecRepo) GetCategorySection(forSearch bool) ([]*internal_models.SectionWithCategory, error) {
	var (
		sectionList   []*internal_models.SectionWithCategory
		dbSectionList []dbSectionWithCategory
		query         string
		err           error
	)

	if forSearch {
		query = `
				SELECT * FROM category AS A 
				INNER JOIN category_section AS B 
				ON B.category_id = A.category_id 
				WHERE old_category_section = false AND old_category = false
				ORDER BY A.category_id`
	} else {
		query = `
				SELECT * FROM category AS A 
				INNER JOIN category_section AS B 
				ON B.category_id = A.category_id 
				WHERE old_category_section = false AND old_category = false
				AND B.service = false
				AND A.service = false
				ORDER BY A.category_id`
	}

	err = r.db.Select(&dbSectionList, query)

	if err != nil {
		logger.LogError("Failed read category section from db", "helpdesk/repo/mysql/repo", "", err)
		return nil, err
	}

	for _, sec := range dbSectionList {
		sectionList = append(sectionList, r.toModelSectionWithCategory(sec))
	}

	return sectionList, nil
}

func (r *CatSecRepo) GetCategorySectionList() ([]internal_models.CategorySectionList, error) {
	var (
		categories []dbCategory
		catWithSec []internal_models.CategorySectionList
		query      string
		err        error
	)

	query = `SELECT * FROM category 
	WHERE service = false
	ORDER BY category_id`

	err = r.db.Select(&categories, query)
	if err != nil {
		logger.LogError("Failed read category list", "pkg_ticket/cat_sec_manager/repo/mysql", "", err)
		return nil, err
	}

	for _, category := range categories {
		var (
			dbSections []dbCategorySection
			sections   []*internal_models.CategorySection
		)

		query = `SELECT * FROM category_section
					WHERE category_id = ?`
		err = r.db.Select(&dbSections, query, category.ID)
		if err != nil {
			logger.LogError("Failed read category sections", "pkg_ticket/cat_sec_manager/repo/mysql", fmt.Sprintf("category id: %d", category.ID), err)
			continue
		}

		for _, section := range dbSections {
			section.ApprovalGroups = make([]uint64, 0)

			if section.NeedApproval {
				query = `SELECT group_id FROM approval_bindings WHERE section_id = ?`

				r.db.Select(&section.ApprovalGroups, query, section.ID)
			}

			sections = append(sections, r.toModelCategorySection(section))
		}
		secList := internal_models.CategorySectionList{
			Category: r.toModelsCategory(category),
			Sections: sections,
		}

		catWithSec = append(catWithSec, secList)
	}

	return catWithSec, nil
}

func (r *CatSecRepo) GetCategorySectionByID(id uint64) (*internal_models.CategorySection, error) {
	var (
		sect  dbCategorySection
		query string
		err   error
	)

	query = `SELECT * FROM category_section
				WHERE section_id = ?`
	err = r.db.Get(&sect, query, id)
	if err != nil {
		logger.LogError(
			"Failed read category section",
			"pkg_ticket/cat_sec_manager/repo/mysql",
			fmt.Sprintf("id: %d", id),
			err,
		)
		return nil, err
	}

	return r.toModelCategorySection(sect), nil
}

func (r *CatSecRepo) GetSectionWithCategoryByID(id uint64) (*internal_models.SectionWithCategory, error) {
	var (
		sect  dbSectionWithCategory
		query string
		err   error
	)

	query = `SELECT * FROM category_section AS CS
				INNER JOIN category AS C ON CS.category_id = C.category_id
				WHERE CS.section_id = ?`

	err = r.db.Get(&sect, query, id)
	if err != nil {
		logger.LogError(
			"Failed read section with category",
			"pkg_ticket/cat_sec_manager/repo/mysql",
			fmt.Sprintf("id: %d", id),
			err,
		)
		return nil, err
	}

	return r.toModelSectionWithCategory(sect), nil
}

func (r *CatSecRepo) CheckExistInResolveGroupList(sectionID, groupID uint64) bool {
	var (
		exist bool
		query string
		err   error
	)

	query = `SELECT EXISTS (SELECT * FROM approval_bindings
				WHERE section_id = ? AND group_id = ?)`

	err = r.db.Get(&exist, query, sectionID, groupID)
	if err != nil {
		logger.LogError(
			"Failed check exist in resolve group list",
			"pkg_ticket/cat_sec_manager/repo/mysql",
			fmt.Sprintf("section id: %d;", sectionID),
			err,
		)
	}

	return exist
}

func (r *CatSecRepo) Close() error {
	r.db.Close()
	return nil
}
