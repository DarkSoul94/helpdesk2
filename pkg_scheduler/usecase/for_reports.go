package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_consts"
	"github.com/DarkSoul94/helpdesk2/pkg_scheduler"
	"github.com/DarkSoul94/helpdesk2/pkg_scheduler/internal_models"
)

type ShedulerForReports struct {
	repo   pkg_scheduler.ISchedulerRepo
	consts pkg_consts.IConstsUsecase
}

func NewShedulerForReports(repo pkg_scheduler.ISchedulerRepo, consts pkg_consts.IConstsUsecase) *ShedulerForReports {
	return &ShedulerForReports{
		repo:   repo,
		consts: consts,
	}
}

func (sr *ShedulerForReports) SupportsShiftsMotivation(startDate, endDate string) ([]*internal_models.Motivation, models.Err) {
	return nil, nil
}
