package usecase

import "github.com/DarkSoul94/helpdesk2/pkg_user/support_manager"

func NewSupportUsecase(repo support_manager.SupportRepo) *SupportUsecase {
	return &SupportUsecase{
		repo: repo,
	}
}
