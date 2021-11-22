package reg_fil_manager

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type IRegFilUsecase interface {
	CreateRegion(reg *internal_models.Region) (uint64, models.Err)
	UpdateRegion(reg *internal_models.Region) models.Err
	DeleteRegion(id uint64) models.Err

	CreateFilial(fil *internal_models.Filial) (uint64, models.Err)
	UpdateFilial(fil *internal_models.Filial) models.Err
	DeleteFilial(id uint64) models.Err
	GetFilialByIp(ip string) (*internal_models.Filial, *internal_models.Region, models.Err)
	GetRegionsWithFilials() ([]*internal_models.RegionWithFilials, models.Err)
}
