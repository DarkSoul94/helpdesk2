package usecase

import (
	"fmt"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"
	report_models "github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"
	ticket_models "github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/shopspring/decimal"
)

func (u *ReportsUsecase) calculateSupportMotivation(
	period report_models.Period,
	suppMotiv report_models.Motivation,
	categories []*ticket_models.Category,
) (report_models.Motivation, models.Err) {

	counts, err := u.repo.GetSupportTicketCountByCategory(period.StartDate, period.EndDate, suppMotiv.Support.ID)
	if err != nil {
		text := fmt.Sprintf("Не удалось получить количество запросов в разрезе категорий по саппорту: %s", suppMotiv.Support.Name)
		return report_models.Motivation{}, models.InternalError(text)
	}

	for _, category := range categories {
		count, ok := counts[category.ID]
		if !ok {
			if category.Old {
				continue
			}
			count = 0
		}

		suppMotiv.ByCategory = append(suppMotiv.ByCategory, internal_models.MotivCategory{
			ID:    category.ID,
			Name:  category.Name,
			Count: count,
		})
		suppMotiv.TotalTicketsCount += count
		suppMotiv.TotalMotivation = suppMotiv.TotalMotivation.Add(category.Price.Mul(decimal.New(int64(count), 0)))
	}

	suppMotiv.Total = suppMotiv.TotalMotivation.Add(suppMotiv.TotalByShifts)

	return suppMotiv, nil
}
