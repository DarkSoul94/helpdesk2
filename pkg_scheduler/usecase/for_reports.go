package usecase

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_consts"
	"github.com/DarkSoul94/helpdesk2/pkg_scheduler"
	"github.com/DarkSoul94/helpdesk2/pkg_scheduler/internal_models"
	"github.com/DarkSoul94/helpdesk2/pkg_support"
	"github.com/shopspring/decimal"
)

type ShedulerForReports struct {
	repo    pkg_scheduler.ISchedulerRepo
	consts  pkg_consts.IConstsUsecase
	support pkg_support.ISuppForScheduler
}

func NewShedulerForReports(repo pkg_scheduler.ISchedulerRepo, consts pkg_consts.IConstsUsecase, support pkg_support.ISuppForScheduler) *ShedulerForReports {
	return &ShedulerForReports{
		repo:    repo,
		consts:  consts,
		support: support,
	}
}

func (sr *ShedulerForReports) SupportsShiftsMotivation(startDate, endDate time.Time) ([]internal_models.Motivation, models.Err) {
	motivations := make([]internal_models.Motivation, 0)
	regularsMap := make(map[uint64][]internal_models.Motivation)

	seniors, regulars, _ := sr.support.GetSupportGroups()
	shifts, err := sr.repo.GetShiftsCount(startDate, endDate)
	if err != nil {
		return nil, err
	}
	for _, val := range regulars {
		regularsMap[val.Senior.ID] = append(regularsMap[val.Senior.ID], internal_models.Motivation{
			SupportID:   val.Support.ID,
			SupportName: val.Support.Name,
			Color:       val.Color,
			Motivation:  val.Wager.Mul(decimal.New(shifts[val.Support.ID], 0)),
		})
	}
	for _, support := range seniors {
		motivations = append(motivations, internal_models.Motivation{
			SupportID:   support.Support.ID,
			SupportName: support.Support.Name,
			Color:       support.Color,
			Motivation:  support.Wager.Mul(decimal.New(shifts[support.Support.ID], 0)),
		})
		motivations = append(motivations, regularsMap[support.Support.ID]...)
	}
	return motivations, nil
}
