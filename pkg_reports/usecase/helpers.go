package usecase

import (
	"fmt"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"
	ticket_models "github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

func (u *ReportsUsecase) calculateSupportMotivation(
	period internal_models.Period,
	suppMotiv *internal_models.Motivation,
	categories []*ticket_models.Category,
) models.Err {

	counts, err := u.repo.GetSupportTicketCountByCategory(period.StartDate, period.EndDate, suppMotiv.Support.ID)
	if err != nil {
		text := fmt.Sprintf("Не удалось получить количество запросов в разрезе категорий по саппорту: %s", suppMotiv.Support.Name)
		return models.InternalError(text)
	}

	for _, category := range categories {
		count, ok := counts[category.ID]
		if !ok {
			if category.Old {
				continue
			}
			count = 0
		}

		catMotiv := &internal_models.MotivCategory{
			ID:    category.ID,
			Name:  category.Name,
			Count: count,
		}

		suppMotiv.TotalTicketsCount += count
		suppMotiv.TotalMotivation = suppMotiv.TotalMotivation.Add(catMotiv.CalcMotiv(category.Price))
		suppMotiv.ByCategory = append(suppMotiv.ByCategory, catMotiv)
	}

	suppMotiv.Total = suppMotiv.TotalMotivation.Add(suppMotiv.TotalByShifts)

	return nil
}
