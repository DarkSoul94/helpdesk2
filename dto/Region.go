package dto

import "github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"

type InpRegion struct {
	ID   uint64 `json:"region_id"`
	Name string `json:"region"`
}

type OutRegionWithFilials struct {
	ID      uint64      `json:"region_id"`
	Name    string      `json:"region"`
	Filials []InpFilial `json:"filials"`
}

func ToModelRegion(reg InpRegion) *internal_models.Region {
	return &internal_models.Region{
		ID:   reg.ID,
		Name: reg.Name,
	}
}

func ToOutRegion(reg *internal_models.Region) InpRegion {
	return InpRegion{
		ID:   reg.ID,
		Name: reg.Name,
	}
}

func ToOutRegionWithFilials(reg *internal_models.RegionWithFilials) OutRegionWithFilials {
	outReg := OutRegionWithFilials{
		ID:   reg.Region.ID,
		Name: reg.Region.Name,
	}

	for _, fil := range reg.Filials {
		outReg.Filials = append(outReg.Filials, ToOutFilial(fil))
	}

	return outReg
}
