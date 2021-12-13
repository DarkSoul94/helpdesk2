package reg_fil_manager

import "github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"

type IRegFilRepo interface {
	CheckRegionExist(id uint64, name string) bool
	CreateRegion(reg *internal_models.Region) (uint64, error)
	UpdateRegion(reg *internal_models.Region) error
	DeleteRegion(id uint64) error

	CheckFilialExist(id, reg_id uint64, name string) bool
	CreateFilial(fil *internal_models.Filial) (uint64, error)
	UpdateFilial(fil *internal_models.Filial) error
	DeleteFilial(id uint64) error
	GetFilialByIp(ip string) (*internal_models.Filial, *internal_models.Region, error)
	GetRegionsWithFilials() ([]*internal_models.RegionWithFilials, error)

	Close() error
}
