package usecase

import (
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
