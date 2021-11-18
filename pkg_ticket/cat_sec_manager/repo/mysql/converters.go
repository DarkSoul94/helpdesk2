package mysql

import "github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"

func (r *CatSecRepo) toDbCategory(mCat *internal_models.Category) dbCategory {
	return dbCategory{
		ID:          mCat.ID,
		Name:        mCat.Name,
		Significant: mCat.Significant,
		Old:         mCat.Old,
		Price:       mCat.Price,
	}
}

func (r *CatSecRepo) toModelsCategory(dbCat dbCategory) *internal_models.Category {
	return &internal_models.Category{
		ID:          dbCat.ID,
		Name:        dbCat.Name,
		Significant: dbCat.Significant,
		Old:         dbCat.Old,
		Price:       dbCat.Price,
	}
}

func (r *CatSecRepo) toDbCategorySection(sec *internal_models.CategorySection) dbCategorySection {
	dbSec := dbCategorySection{
		ID:           sec.ID,
		CategoryID:   sec.CategoryID,
		Name:         sec.Name,
		Significant:  sec.Significant,
		Old:          sec.Old,
		NeedApproval: sec.NeedApproval,
	}

	if len(sec.Template) > 0 {
		dbSec.Template.String = sec.Template
		dbSec.Template.Valid = true
	} else {
		dbSec.Template.Valid = false
	}

	return dbSec
}

func (r *CatSecRepo) toModelCategorySection(dbSec dbCategorySection) *internal_models.CategorySection {
	sec := &internal_models.CategorySection{
		ID:             dbSec.ID,
		CategoryID:     dbSec.CategoryID,
		Name:           dbSec.Name,
		Significant:    dbSec.Significant,
		Old:            dbSec.Old,
		NeedApproval:   dbSec.NeedApproval,
		ApprovalGroups: dbSec.ApprovalGroups,
	}

	if dbSec.Template.Valid {
		sec.Template = dbSec.Template.String
	}

	return sec
}

func (r *CatSecRepo) toModelSectionWithCategory(dbSec dbSectionWithCategory) *internal_models.SectionWithCategory {
	sec := &internal_models.SectionWithCategory{
		ID:           dbSec.ID,
		Name:         dbSec.Name,
		Significant:  dbSec.Significant,
		Old:          dbSec.Old,
		NeedApproval: dbSec.NeedApproval,
		Cat:          r.toModelsCategory(*dbSec.Category),
	}

	if dbSec.Template.Valid {
		sec.Template = dbSec.Template.String
	}

	return sec
}
