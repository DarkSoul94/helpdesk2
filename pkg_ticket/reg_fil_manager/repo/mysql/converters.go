package mysql

import "github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"

func (r *RegFilRepo) toDbRegion(reg *internal_models.Region) dbRegion {
	return dbRegion{
		ID:   reg.ID,
		Name: reg.Name,
	}
}

func (r *RegFilRepo) toModelRegion(reg dbRegion) *internal_models.Region {
	return &internal_models.Region{
		ID:   reg.ID,
		Name: reg.Name,
	}
}

func (r *RegFilRepo) toDbFilial(fil *internal_models.Filial) dbFilial {
	return dbFilial{
		ID:       fil.ID,
		RegionID: fil.RegionID,
		Name:     fil.Name,
		IP:       fil.IP,
	}
}

func (r *RegFilRepo) toModelFilial(fil dbFilial) *internal_models.Filial {
	return &internal_models.Filial{
		ID:       fil.ID,
		RegionID: fil.RegionID,
		Name:     fil.Name,
		IP:       fil.IP,
	}
}
