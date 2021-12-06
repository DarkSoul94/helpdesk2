package dto

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg_scheduler/internal_models"
)

type OutSupportLateness struct {
	ID         uint64 `json:"id"`
	Date       string `json:"date"`
	Name       string `json:"name"`
	Cause      string `json:"cause"`
	DecisionID uint64 `json:"decision_id"`
	Difference uint64 `json:"difference"`
}

type OutLatenessDecision struct {
	ID   uint64 `json:"id"`
	Text string `json:"text"`
}

func ToOutSupportLateness(mLate *internal_models.Lateness) *OutSupportLateness {
	return &OutSupportLateness{
		ID:         mLate.ID,
		Date:       mLate.Date.Format("2006-01-02 15:04:05"),
		Name:       mLate.SupportName,
		Cause:      mLate.Cause,
		DecisionID: mLate.Decision.ID,
		Difference: uint64(time.Duration(mLate.Difference * uint64(time.Minute)).Minutes()),
	}
}

func ToOutDecisionsList(mList []internal_models.Decision) []OutLatenessDecision {
	outList := make([]OutLatenessDecision, 0)
	for _, val := range mList {
		outList = append(outList, OutLatenessDecision{
			ID:   val.ID,
			Text: val.Text,
		})
	}
	return outList
}
