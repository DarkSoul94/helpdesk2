package dto

import "github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"

type InpFilial struct {
	ID       uint64 `json:"filial_id"`
	RegionID uint64 `json:"region_id"`
	Name     string `json:"filial"`
	IP       string `json:"ip"`
}

func ToModelFilial(fil InpFilial) *internal_models.Filial {
	return &internal_models.Filial{
		ID:       fil.ID,
		RegionID: fil.RegionID,
		Name:     fil.Name,
		IP:       fil.IP,
	}
}

func ToOutFilial(fil *internal_models.Filial) InpFilial {
	return InpFilial{
		ID:       fil.ID,
		RegionID: fil.RegionID,
		Name:     fil.Name,
		IP:       fil.IP,
	}
}
