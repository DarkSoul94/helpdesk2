package mysql

import (
	"database/sql"
	"fmt"

	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/jmoiron/sqlx"
)

func NewRegFilRepo(db *sql.DB) *RegFilRepo {
	return &RegFilRepo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *RegFilRepo) CheckRegionExist(id uint64, name string) bool {
	var (
		db_id uint64
		query string
		err   error
	)

	if id != 0 {
		query = `SELECT region_id FROM regions
			WHERE region_id = ?`

		err = r.db.Get(&db_id, query, id)
		if err != nil {
			return false
		}
	} else {
		query = `SELECT region_id FROM regions
			WHERE region = ?`
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

func (r *RegFilRepo) CreateRegion(reg *internal_models.Region) (uint64, error) {
	var (
		res   sql.Result
		query string
		err   error
	)

	query = `INSERT INTO regions SET region = :region`
	res, err = r.db.NamedExec(query, r.toDbRegion(reg))
	if err != nil {
		logger.LogError(
			`Failed create region`,
			"pkg_ticket/reg_fil_manager/repo/mysql",
			fmt.Sprintf("name: %s", reg.Name),
			err,
		)
		return 0, err
	}

	id, _ := res.LastInsertId()

	return uint64(id), nil
}

func (r *RegFilRepo) UpdateRegion(reg *internal_models.Region) error {
	var (
		query string
		err   error
	)

	query = `UPDATE regions SET 
				region = :region
				WHERE region_id = :region_id`

	_, err = r.db.NamedExec(query, r.toDbRegion(reg))
	if err != nil {
		logger.LogError(
			"Failed update region",
			"pkg_ticket/reg_fil_manager/repo/mysql",
			fmt.Sprintf("id: %d; name: %s;", reg.ID, reg.Name),
			err,
		)
		return err
	}

	return nil
}

func (r *RegFilRepo) DeleteRegion(id uint64) error {
	_, err := r.db.Exec(`DELETE FROM regions WHERE region_id = ?`, id)
	if err != nil {
		logger.LogError(
			"Failed delete region",
			"pkg_ticket/reg_fil_manager/repo/mysql",
			fmt.Sprintf("id: %d", id),
			err,
		)
		return err
	}

	return nil
}

func (r *RegFilRepo) CheckFilialExist(id, reg_id uint64, name string) bool {
	var (
		db_id uint64
		query string
		err   error
	)

	if id != 0 {
		query = `SELECT filial_id FROM filials
			WHERE filial_id = ? AND region_id = ?`

		err = r.db.Get(&db_id, query, id, reg_id)
		if err != nil {
			return false
		}
	} else {
		query = `SELECT filial_id FROM filials
			WHERE filial = ? AND region_id = ?`
		err = r.db.Get(&db_id, query, name, reg_id)
		if err != nil {
			return false
		}
	}

	if db_id == 0 {
		return false
	}

	return true
}

func (r *RegFilRepo) CreateFilial(fil *internal_models.Filial) (uint64, error) {
	var (
		res   sql.Result
		query string
		err   error
	)

	query = `INSERT INTO filials SET 
				region_id = :region_id,
				filial = :filial,
				ip = :ip`

	res, err = r.db.NamedExec(query, r.toDbFilial(fil))
	if err != nil {
		logger.LogError(
			`Failed create filial`,
			"pkg_ticket/reg_fil_manager/repo/mysql",
			fmt.Sprintf("region id: %d; name: %s; ip: %s;", fil.RegionID, fil.Name, fil.IP),
			err,
		)
		return 0, err
	}

	id, _ := res.LastInsertId()

	return uint64(id), nil
}

func (r *RegFilRepo) UpdateFilial(fil *internal_models.Filial) error {
	var (
		query string
		err   error
	)

	query = `UPDATE filials SET
				filial = :filial,
				ip = :ip
				WHERE filial_id = :filial_id`
	_, err = r.db.NamedExec(query, r.toDbFilial(fil))
	if err != nil {
		logger.LogError(
			`Failed create filial`,
			"pkg_ticket/reg_fil_manager/repo/mysql",
			fmt.Sprintf("id: %d; region id: %d; name: %s; ip: %s;", fil.ID, fil.RegionID, fil.Name, fil.IP),
			err,
		)
		return err
	}

	return nil
}

func (r *RegFilRepo) DeleteFilial(id uint64) error {
	_, err := r.db.Exec(`DELETE FROM filials WHERE filial_id = ?`, id)
	if err != nil {
		logger.LogError(
			"Failed delete filial",
			"pkg_ticket/reg_fil_manager/repo/mysql",
			fmt.Sprintf("id: %d", id),
			err,
		)
		return err
	}

	return nil
}

func (r *RegFilRepo) GetFilialByIp(ip string) (*internal_models.Filial, *internal_models.Region, error) {
	var (
		fil   dbFilial
		reg   dbRegion
		query string
		err   error
	)

	query = `SELECT * FROM filials WHERE ip = ?`
	err = r.db.Get(&fil, query, ip)
	if err != nil {
		logger.LogError(
			"Failed read filial",
			"pkg_ticket/reg_fil_manager/repo/mysql",
			fmt.Sprintf("ip: %s;", ip),
			err,
		)
		return nil, nil, err
	}

	query = `SELECT * FROM regions WHERE region_id = ?`
	err = r.db.Get(&reg, query, fil.RegionID)
	if err != nil {
		logger.LogError(
			"Failed read region",
			"pkg_ticket/reg_fil_manager/repo/mysql",
			fmt.Sprintf("id: %d;", fil.RegionID),
			err,
		)
		return nil, nil, err
	}

	return r.toModelFilial(fil), r.toModelRegion(reg), nil
}

func (r *RegFilRepo) GetRegionsWithFilials() ([]*internal_models.RegionWithFilials, error) {
	var (
		outList   []*internal_models.RegionWithFilials
		dbRegions []dbRegion
		query     string
		err       error
	)

	query = `SELECT * FROM regions`
	err = r.db.Select(&dbRegions, query)
	if err != nil {
		logger.LogError(
			`Failed read regions`,
			"pkg_ticket/reg_fil_manager/repo/mysql",
			"",
			err,
		)
		return nil, err
	}

	for _, dbReg := range dbRegions {
		temp := &internal_models.RegionWithFilials{Region: r.toModelRegion(dbReg), Filials: make([]*internal_models.Filial, 0)}

		var dbFilials []dbFilial
		query = `SELECT * FROM filials 
					WHERE region_id = ?`
		r.db.Select(&dbFilials, query, dbReg.ID)

		for _, filial := range dbFilials {
			temp.Filials = append(temp.Filials, r.toModelFilial(filial))
		}

		outList = append(outList, temp)
	}

	return outList, nil
}

func (r *RegFilRepo) Close() error {
	r.db.Close()
	return nil
}
