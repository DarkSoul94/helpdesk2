package usecase

import "github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"

func (u *ReportsUsecase) summaryMotivation(motivation map[string][]internal_models.Motivation) map[string][]internal_models.Motivation {
	const total_motiv string = "Общая мотивация за период"

	resultMap := make(map[string][]internal_models.Motivation)
	resultMap[total_motiv] = make([]internal_models.Motivation, 0)

	result := make(map[uint64]internal_models.Motivation)

	for period, motiv := range motivation {
		for _, supp := range motiv {
			resMotiv, ok := result[supp.Support.ID]
			if !ok {
				result[supp.Support.ID] = supp
				continue
			}
			for index, categoryMotiv := range supp.ByCategory {
				resMotiv.ByCategory[index].Count += categoryMotiv.Count
			}

			resMotiv.TotalTicketsCount += supp.TotalTicketsCount
			resMotiv.TotalMotivation = supp.TotalMotivation.Add(resMotiv.TotalMotivation)
			resMotiv.TotalByShifts = supp.TotalByShifts.Add(resMotiv.TotalByShifts)
			resMotiv.Total = supp.Total.Add(resMotiv.Total)
		}
		resultMap[period] = motiv
	}
	for _, val := range result {
		resultMap[total_motiv] = append(resultMap[total_motiv], val)
	}
	return resultMap
}
