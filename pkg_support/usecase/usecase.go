package usecase

import "github.com/DarkSoul94/helpdesk2/pkg_support"

func NewSupportUsecase(repo pkg_support.ISupportRepo) *SupportUsecase {
	return &SupportUsecase{
		repo: repo,
	}
}
