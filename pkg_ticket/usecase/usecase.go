package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/cat_sec_manager"
)

type TicketUsecase struct {
	ticketRepo pkg_ticket.ITicketRepo
	catSecUC   cat_sec_manager.ICatSecUsecase
}

func NewTicketUsecase(
	tRepo pkg_ticket.ITicketRepo,
	catSecUC cat_sec_manager.ICatSecUsecase,
) *TicketUsecase {
	return &TicketUsecase{
		ticketRepo: tRepo,
		catSecUC:   catSecUC,
	}
}

func (u *TicketUsecase) CreateCategory(cat *cat_sec_manager.Category) (uint64, models.Err) {
	return u.catSecUC.CreateCategory(*cat)
}
