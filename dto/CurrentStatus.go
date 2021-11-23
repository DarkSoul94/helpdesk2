package dto

import (
	"github.com/DarkSoul94/helpdesk2/global_const/literal_keys"
	suppModels "github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
)

type CurrentStatus struct {
	SupportsStatus []SupportCurrentStatus `json:"support_current_status"`
	Total          map[string]int         `json:"total"`
	Wait           int                    `json:"wait_ticket_count"`
}

type SupportCurrentStatus struct {
	SupportID   uint64 `json:"support_id"`
	Support     string `json:"support"`
	Status      string `json:"status"`
	ShiftStatus bool   `json:"shift_status"`
	InWork      int    `json:"in_work"`
	Postproned  int    `json:"postproned"`
	Revision    int    `json:"revision"`
	Complete    int    `json:"complete"`
}

func ToOutCurrentStatus(supportsInfo []*suppModels.SupportInfo, total map[string]int) CurrentStatus {
	statuses := make([]SupportCurrentStatus, 0)
	for _, val := range supportsInfo {
		statuses = append(statuses, toOutSupportCurrentStatus(val))
	}
	outTotal := make(map[string]int)
	for key, val := range total {
		if key != literal_keys.TS_Wait {
			outKey := "total_"
			outTotal[outKey+key] = val
		}
	}
	return CurrentStatus{
		SupportsStatus: statuses,
		Total:          outTotal,
		Wait:           total[literal_keys.TS_Wait],
	}
}

func toOutSupportCurrentStatus(supp *suppModels.SupportInfo) SupportCurrentStatus {
	return SupportCurrentStatus{
		SupportID:   supp.Shift.Support.ID,
		Support:     supp.Shift.Support.Name,
		Status:      supp.Shift.Support.Status.Name,
		ShiftStatus: !supp.Shift.ClosingStatus,
		InWork:      supp.Tickets[literal_keys.TS_InWork],
		Postproned:  supp.Tickets[literal_keys.TS_Postponed],
		Revision:    supp.Tickets[literal_keys.TS_Revision],
		Complete:    supp.Tickets[literal_keys.TS_Completed],
	}
}
