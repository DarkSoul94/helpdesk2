package usecase

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_consts"
	"github.com/shopspring/decimal"
)

func (sr *ShedulerForReports) latenessHelper(date time.Time) (map[uint64]decimal.Decimal, models.Err) {
	var (
		graceTime uint64
		penalty   decimal.Decimal
	)
	penaltyMap := make(map[uint64]decimal.Decimal)
	lateTime := make(map[uint64]uint64)
	sDate := date.Format("2006-01") + "-01"

	lateConfig, _ := sr.consts.GetConstForDate(date, pkg_consts.KeyConfig)
	graceTime = lateConfig[pkg_consts.KeyConfig_Grace].(uint64)
	penalty = lateConfig[pkg_consts.KeyConfig_Penalty].(decimal.Decimal)

	lateness, err := sr.repo.GetLateness(sDate)
	if err != nil {
		return nil, err
	}
	for _, late := range lateness {
		if _, val := late.Decision.GetDecisionValue(); !val {
			lateTime[late.SupportID] += late.Difference
		}
	}
	for suppportID, time := range lateTime {
		if time > graceTime {
			time -= graceTime
			penaltyMap[suppportID] = penalty.Mul(decimal.New(int64(time), 0))
		}
	}
	return penaltyMap, nil
}
