package file_manager

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type IFileUsecase interface {
	CreateFiles(files []*internal_models.File, ticketID uint64) models.Err
	GetFile(fileID uint64) (*internal_models.File, models.Err)
	GetTicketFiles(ticketID uint64) ([]*internal_models.File, models.Err)
}
