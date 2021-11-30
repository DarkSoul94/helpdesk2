package pkg_ticket

import (
	"context"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type ITicketUsecase interface {
	CreateCategory(cat *internal_models.Category) (uint64, models.Err)
	UpdateCategory(cat *internal_models.Category) models.Err

	CreateCategorySection(sec *internal_models.CategorySection) (uint64, models.Err)
	UpdateCategorySection(sec *internal_models.CategorySection) models.Err
	GetCategorySection(forSearch bool) ([]*internal_models.SectionWithCategory, models.Err)
	GetCategorySectionList() ([]internal_models.CategorySectionList, models.Err)

	CreateRegion(reg *internal_models.Region) (uint64, models.Err)
	UpdateRegion(reg *internal_models.Region) models.Err
	DeleteRegion(id uint64) models.Err

	CreateFilial(fil *internal_models.Filial) (uint64, models.Err)
	UpdateFilial(fil *internal_models.Filial) models.Err
	DeleteFilial(id uint64) models.Err
	GetRegionsWithFilials() ([]*internal_models.RegionWithFilials, models.Err)

	GetTicketStatuses(group_id uint64, all bool) ([]*internal_models.TicketStatus, models.Err)
	GetAllTicketStatusHistory(ticketID uint64, user *models.User) ([]*internal_models.TicketStatusHistory, models.Err)
	CreateTicket(ticket *internal_models.Ticket) (uint64, models.Err)
	GenerateTicket(ticket internal_models.TicketGenerate, author *models.User) models.Err
	UpdateTicket(ticket *internal_models.Ticket, user *models.User, checkControls bool) models.Err
	GetTicketList(user *models.User, limit, offset int) ([]*internal_models.Ticket, []string, map[uint]uint, models.Err)
	GetTicket(ticketID uint64, user *models.User) (*internal_models.Ticket, models.Err)
	StealTicket(id uint64, newSupport *models.User) models.Err
	CheckNeedApprovalTicketExist(groupID uint64) bool
	GetApprovalTicketList(groupID uint64, limit, offset int) ([]*internal_models.Ticket, []string, models.Err)

	DistributeTicket(ctx context.Context)

	CreateComment(comment *internal_models.Comment) (uint64, models.Err)
}
