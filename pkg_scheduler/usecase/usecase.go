package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_scheduler"
	"github.com/DarkSoul94/helpdesk2/pkg_scheduler/internal_models"
)

type SchedulerUsecase struct {
	repo pkg_scheduler.ISchedulerRepo
}

func NewSchedulerUsecase(repo pkg_scheduler.ISchedulerRepo) *SchedulerUsecase {
	return &SchedulerUsecase{
		repo: repo,
	}
}

func (u *SchedulerUsecase) UpdateOfficesList(offices []*internal_models.Office) models.Err {
	for _, office := range offices {
		if office.ID == 0 {
			return u.repo.AddOffice(office)
		} else {
			currOffice, err := u.repo.GetOfficeByID(office.ID)
			if err != nil || office.Compare(currOffice) {
				continue
			}
			if err := u.repo.UpdateOffice(office); err != nil {
				return err
			}
		}
	}
	return nil
}

func (u *SchedulerUsecase) GetOfficesList() (actual, deleted []*internal_models.Office, err models.Err) {
	if actual, err = u.repo.GetOfficesList(false); err != nil {
		err = models.InternalError("Не удалось получить список актуальных офисов")
		return
	}

	if deleted, err = u.repo.GetOfficesList(true); err != nil {
		err = models.InternalError("Не удалось получить список удаленных офисов")
		return
	}
	return
}

func (u *SchedulerUsecase) UpdateShiftsSchedule(schedule []*internal_models.Cell) models.Err {
	actual := make(map[string][]uint64)
	for _, cell := range schedule {
		if err := u.repo.UpdateCell(cell); err != nil {
			return err
		}
		date := cell.Date.Format("2006-01") + "-01"
		actual[date] = append(actual[date], cell.ID)

	}
	return u.repo.DeleteCells(actual)
}

func (u *SchedulerUsecase) GetSchedule(date string) ([]*internal_models.Cell, []*internal_models.Office, models.Err) {
	var (
		schedule         []*internal_models.Cell
		offices, deleted []*internal_models.Office
		err              models.Err
	)

	if schedule, err = u.repo.GetSchedule(date); err != nil {
		return nil, nil, err
	}
	if len(schedule) == 0 {
		return nil, nil, nil
	}

	if offices, err = u.repo.GetOfficesList(false); err != nil {
		return nil, nil, models.InternalError("Не удалось получить список офисов")
	}
	if deleted, err = u.repo.GetOfficesList(true, date); err != nil {
		return nil, nil, models.InternalError("Не удалось получить список офисов")
	}
	offices = append(offices, deleted...)
	return schedule, offices, nil
}

func (u *SchedulerUsecase) GetLateness(date string) ([]*internal_models.Lateness, models.Err) {
	list, err := u.repo.GetLateness(date)
	if err != nil {
		return nil, models.InternalError("Не удалось получить список опозданий за месяц")
	}
	return list, nil
}

func (u *SchedulerUsecase) UpdateLateness(latenessID, decisionID uint64) models.Err {
	lateness, err := u.repo.GetLatenessByID(latenessID)
	if err != nil {
		return err
	}
	decision, err := internal_models.GetLateDecision(decisionID)
	if err != nil {
		return err
	}
	lateness.Decision = &decision
	return u.repo.UpdateLateness(lateness)
}

func (u *SchedulerUsecase) CheckNewLateness() bool {
	return u.repo.CheckNewLateness()
}
