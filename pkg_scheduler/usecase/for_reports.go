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
	penaltyMap, err := sr.latenessHelper(startDate)
	if err != nil {
		return nil, err
	}

	shifts, err := sr.repo.GetShiftsCount(startDate, endDate)
	if err != nil {
		return nil, err
	}

	for _, support := range regulars {
		motiv := support.Wager.Mul(decimal.New(shifts[support.Support.ID], 0))
		if penalty, ok := penaltyMap[support.Support.ID]; ok {
			motiv = motiv.Sub(penalty)
		}
		if support.Senior != nil {
			regularsMap[support.Senior.ID] = append(regularsMap[support.Senior.ID], internal_models.Motivation{
				SupportID:   support.Support.ID,
				SupportName: support.Support.Name,
				Color:       support.Color,
				Motivation:  motiv,
			})
		} else {
			regularsMap[0] = append(regularsMap[0], internal_models.Motivation{
				SupportID:   support.Support.ID,
				SupportName: support.Support.Name,
				Color:       support.Color,
				Motivation:  motiv,
			})
		}
	}
	for _, support := range seniors {
		motiv := support.Wager.Mul(decimal.New(shifts[support.Support.ID], 0))
		if penalty, ok := penaltyMap[support.Support.ID]; ok {
			motiv = motiv.Sub(penalty)
		}

		motivations = append(motivations, internal_models.Motivation{
			SupportID:   support.Support.ID,
			SupportName: support.Support.Name,
			Color:       support.Color,
			Motivation:  motiv,
		})
		motivations = append(motivations, regularsMap[support.Support.ID]...)
	}
	if val, ok := regularsMap[0]; ok {
		motivations = append(motivations, val...)
	}
	return motivations, nil
}
