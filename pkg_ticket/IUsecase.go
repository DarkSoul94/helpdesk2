package pkg_ticket

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/cat_sec_manager"
)

type ITicketUsecase interface {
	CreateCategory(cat *cat_sec_manager.Category) (uint64, models.Err)
}
