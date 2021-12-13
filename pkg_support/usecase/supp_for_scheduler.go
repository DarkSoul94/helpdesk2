package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_support"
	"github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
)

type SuppForScheduler struct {
	repo pkg_support.ISupportRepo
}

func NewSuppForSchedulerUsecase(repo pkg_support.ISupportRepo) *SuppForScheduler {
	return &SuppForScheduler{
		repo: repo,
	}
}

func (us *SuppForScheduler) GetSupportGroups() (seniors, regulars []*internal_models.Card, err models.Err) {
	var list []*internal_models.Card
	list, err = us.repo.GetCardsList()
	if err != nil {
		return
	}
	for _, val := range list {
		if val.IsSenior {
			seniors = append(seniors, val)
			continue
		}
		regulars = append(regulars, val)
	}
	return
}
