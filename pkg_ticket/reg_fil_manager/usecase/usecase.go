package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/reg_fil_manager"
)

type RegFilUsecase struct {
	repo reg_fil_manager.IRegFilRepo
}

func NewRegFilUsecase(repo reg_fil_manager.IRegFilRepo) *RegFilUsecase {
	return &RegFilUsecase{
		repo: repo,
	}
}

func (u *RegFilUsecase) CreateRegion(reg *internal_models.Region) (uint64, models.Err) {
	if u.repo.CheckRegionExist(0, reg.Name) {
		return 0, ErrRegionExist
	}

	id, err := u.repo.CreateRegion(reg)
	if err != nil {
		return 0, models.InternalError(err.Error())
	}

	return id, nil
}

func (u *RegFilUsecase) UpdateRegion(reg *internal_models.Region) models.Err {
	if !u.repo.CheckRegionExist(reg.ID, reg.Name) {
		return ErrRegionExist
	}

	err := u.repo.UpdateRegion(reg)
	if err != nil {
		return models.InternalError(err.Error())
	}

	return nil
}

func (u *RegFilUsecase) DeleteRegion(id uint64) models.Err {
	err := u.repo.DeleteRegion(id)
	if err != nil {
		return models.InternalError(err.Error())
	}

	return nil
}

func (u *RegFilUsecase) CreateFilial(fil *internal_models.Filial) (uint64, models.Err) {
	if u.repo.CheckFilialExist(0, fil.RegionID, fil.Name) {
		return 0, ErrFilialExist
	}

	id, err := u.repo.CreateFilial(fil)
	if err != nil {
		return 0, models.InternalError(err.Error())
	}

	return id, nil
}

func (u *RegFilUsecase) UpdateFilial(fil *internal_models.Filial) models.Err {
	if !u.repo.CheckFilialExist(fil.ID, fil.RegionID, fil.Name) {
		return ErrFilialNotExist
	}

	err := u.repo.UpdateFilial(fil)
	if err != nil {
		return models.InternalError(err.Error())
	}

	return nil
}

func (u *RegFilUsecase) DeleteFilial(id uint64) models.Err {
	err := u.repo.DeleteFilial(id)
	if err != nil {
		return models.InternalError(err.Error())
	}

	return nil
}

func (u *RegFilUsecase) GetFilialByIp(ip string) (*internal_models.Filial, *internal_models.Region, models.Err) {
	fil, reg, err := u.repo.GetFilialByIp(ip)
	if err != nil {
		return nil, nil, models.InternalError(err.Error())
	}

	return fil, reg, nil
}

func (u *RegFilUsecase) GetRegionsWithFilials() ([]*internal_models.RegionWithFilials, models.Err) {
	list, err := u.repo.GetRegionsWithFilials()
	if err != nil {
		return nil, models.InternalError(err.Error())
	}

	return list, nil
}
