package usecase

import (
	"github.com/DarkSoul94/helpdesk2/global_const/literal_keys"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
)

func (u *SupportUsecase) changePriority(support *internal_models.Support) models.Err {
	prior := u.repo.GetPrioritizedSupportID()
	support.Status, _ = u.repo.GetStatus(support.Status.ID)
	switch prior {
	case 0:
		if support.Status.AcceptTicket {
			support.Priority = true
		}
	case support.ID:
		activeSupp, _ := u.repo.GetActiveSupports()
		if !(len(activeSupp) <= 1 && support.Status.AcceptTicket) {
			if nextPrior := nextPrioritizedSupport(support, activeSupp); nextPrior != nil {
				u.repo.UpdateSupport(nextPrior)
			}
			support.Priority = false
			return u.repo.UpdateSupport(support)
		}

	}

	return u.repo.UpdateSupport(support)
}

//nextPrioritizedSupport анализирует список активных саппортов.
//Возвращает саппорта на которого необходимо переместить приоритет распределения запросов
func nextPrioritizedSupport(support *internal_models.Support, active []*internal_models.Support) *internal_models.Support {
	for id, val := range active {
		if val.ID == support.ID {
			if id+1 < len(active) {
				active[id+1].Priority = true
				return active[id+1]
			}
			if active[0].ID != support.ID {
				active[0].Priority = true
				return active[0]
			}
		}
	}
	return nil
}

//statusHistoryHelper предназначена для работы с историей статусов суппорта
func (u *SupportUsecase) statusHistoryHelper(support *internal_models.Support, shiftID uint64) models.Err {
	statusHistory, _ := u.repo.GetLastStatusHistory(support.ID, shiftID)
	if statusHistory != nil {
		//расчет и установка длительности нахождения в статусе
		statusHistory.SetDuration()
		if err := u.repo.UpdateHistoryRecord(statusHistory); err != nil {
			return err
		}
	}
	newHistory := internal_models.StatusHistory{}
	newHistory.New(support, shiftID)
	return u.repo.CreateHistoryRecord(&newHistory)
}

func (u *SupportUsecase) getInfoHelper(support *internal_models.Support) (*internal_models.SupportInfo, models.Err) {
	var (
		info = internal_models.SupportInfo{
			Tickets: map[string]int{},
		}
	)

	shift, err := u.repo.GetLastShift(support.ID)
	if err != nil {
		return nil, err
	}

	shift.Support = support
	if shift.Support.Status, err = u.repo.GetStatus(support.Status.ID); err != nil {
		return nil, err
	}

	countArgs := []string{
		literal_keys.TS_InWork,
		literal_keys.TS_Implementation,
		literal_keys.TS_Postponed,
		literal_keys.TS_Revision,
	}

	todayCountArgs := []string{
		literal_keys.TS_Completed,
	}
	for key, val := range u.ticket.GetTicketsCounts(support.ID, countArgs...) {
		if key == literal_keys.TS_InWork ||
			key == literal_keys.TS_Implementation {
			info.Tickets[literal_keys.TS_InWork] += val
			continue
		}
		info.Tickets[key] = val
	}

	for key, val := range u.ticket.GetTodayTicketsCounts(support.ID, todayCountArgs...) {
		info.Tickets[key] = val
	}
	info.Shift = shift
	return &info, nil
}

func (u *SupportUsecase) totalInfoHelper() map[string]int {
	info := make(map[string]int)
	countArgs := []string{
		literal_keys.TS_InWork,
		literal_keys.TS_Implementation,
		literal_keys.TS_Postponed,
		literal_keys.TS_Revision,
		literal_keys.TS_Wait,
	}

	todayCountArgs := []string{
		literal_keys.TS_Completed,
	}

	for key, val := range u.ticket.GetTicketsCounts(0, countArgs...) {
		if key == literal_keys.TS_InWork ||
			key == literal_keys.TS_Implementation {
			info[literal_keys.TS_InWork] += val
			continue
		}

		info[key] = val
	}

	for key, val := range u.ticket.GetTodayTicketsCounts(0, todayCountArgs...) {
		info[key] = val
	}

	return info
}
