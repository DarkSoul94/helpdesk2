package pkg_scheduler

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_scheduler/internal_models"
	supp_models "github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
)

type ISchedulerUsecase interface {
	UpdateOfficesList(offices []*internal_models.Office) models.Err
	GetOfficesList() (actual, deleted []*internal_models.Office, err models.Err)
	UpdateShiftsSchedule(schedule []*internal_models.Cell) models.Err
	GetSchedule(date string) ([]*internal_models.Cell, []*internal_models.Office, models.Err)

	GetSupportGroups() (seniors, regulars []*supp_models.Card, err models.Err)

	GetLateness(date string) ([]*internal_models.Lateness, models.Err)
	UpdateLateness(latenessID, decisionID uint64) models.Err
	CheckNewLateness() bool
}

type ISuppSchedulerUsecase interface {
	CreateLateness(supportID uint64, cause string, time time.Time)
	IsSupportLate(supportID uint64) bool
	CheckShiftInScheduler(supportID uint64) models.Err
}
