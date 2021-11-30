package file_manager

import "github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"

type IFileRepo interface {
	CreateFile(file *internal_models.File) error
	GetFile(fileID uint64) (*internal_models.File, error)
	GetTicketFiles(ticketID uint64) ([]*internal_models.File, error)
	Close() error
}
