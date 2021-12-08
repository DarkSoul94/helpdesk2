package usecase

import (
	"strconv"
	"time"

	"github.com/DarkSoul94/helpdesk2/cachettl"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_scheduler"
	"github.com/DarkSoul94/helpdesk2/pkg_scheduler/internal_models"
	"github.com/spf13/viper"
)

type ShedulerForSupport struct {
	repo  pkg_scheduler.ISchedulerRepo
	store cachettl.ObjectStore
}

//NewScheduleManager ...
func NewScheduleForSupport(repo pkg_scheduler.ISchedulerRepo) *ShedulerForSupport {
	return &ShedulerForSupport{
		repo:  repo,
		store: *cachettl.NewObjectStore(viper.GetDuration("app.ttl_cache.clear_period") * time.Second),
	}
}

func (ss *ShedulerForSupport) CheckShiftInScheduler(supportID uint64) models.Err {
	var late internal_models.Lateness
	shift, err := ss.repo.GetTodayShift(supportID)
	if err != nil {
		return err
	}
	if shift != nil {
		if shift.Vacation {
			return models.BadRequest("Невозможно открыть смену так как в графике отмечен отпуск")
		}
		if shift.SickLeave {
			return models.BadRequest("Невозможно открыть смену так как в графике отмечен больничный")
		}
		if shift.Late() {
			late.New(supportID, shift.StartTime)
			ss.store.Add(strconv.Itoa(int(supportID)), late, viper.GetInt64("app.ttl_cache.late_cause_life"))
			return models.ErrorWithSuccess("Укажите причину опоздания")
		}
	}
	return nil
}

func (ss *ShedulerForSupport) CreateLateness(supportID uint64, cause string) (time.Time, models.Err) {
	var lateness = new(internal_models.Lateness)
	if err := ss.store.Get(strconv.Itoa(int(supportID)), lateness); err != nil {
		return time.Time{}, models.BadRequest("По саппорту нет опозданий") //TODO вынести в отдельный error
	}
	lateness.Cause = cause
	if err := ss.repo.CreateLateness(lateness); err != nil {
		return time.Time{}, err
	}
	ss.store.Delete(strconv.Itoa(int(supportID)))

	return lateness.Date, nil
}
