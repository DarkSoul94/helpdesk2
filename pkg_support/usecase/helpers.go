package usecase

import (
	"github.com/DarkSoul94/helpdesk2/global_const/literal_keys"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
)

//priorityHelper модифицирует переданного саппорта в зависимости от нового статуса и текущего приоритетного саппорта.
//Также по необходимости находит саппорта на которого необходимо переместить признак приоритетности.
//Возвращает массив саппортов готовых к обновлению.
func (u *SupportUsecase) priorityHelper(support *internal_models.Support) []*internal_models.Support {
	forUpdate := make([]*internal_models.Support, 0)
	prioritized := u.repo.GetPrioritizedSupportID()

	if support.Status.AcceptTicket && prioritized == 0 {
		support.Priority = true
	}

	if !support.Status.AcceptTicket && prioritized == support.ID {
		support.Priority = false
		activeSupp, _ := u.repo.GetActiveSupports()
		if nextSupp := nextPrioritizedSupport(support, activeSupp); nextSupp != nil {
			forUpdate = append(forUpdate, nextSupp)
		}
	}

	forUpdate = append(forUpdate, support)
	return forUpdate
}

//nextPrioritizedSupport анализирует список активных саппортов.
//Возвращает саппорта на которого необходимо переместить приоритет распределения запросов
func nextPrioritizedSupport(support *internal_models.Support, active []*internal_models.Support) *internal_models.Support {
	for id, val := range active {
		if val.ID == support.ID {
			if id+1 <= len(active) {
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
	statusHistory, err := u.repo.GetLastStatusHistory(support.ID, shiftID)
	if err == nil && statusHistory != nil {
		//расчет и установка длительности нахождения в статусе
		statusHistory.SetDuration()
		if err := u.repo.UpdateHistoryRecord(statusHistory); err != nil {
			return err
		}
	}
	statusHistory.New(support, shiftID)
	return u.repo.CreateHistoryRecord(statusHistory)
}

func (u *SupportUsecase) getInfoHelper(support *internal_models.Support) (*internal_models.SupportInfo, models.Err) {
	info := new(internal_models.SupportInfo)
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

	return info, nil
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
