package dto

import "github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"

type OutSupportStatus struct {
	ID   uint64 `json:"support_status_id"`
	Name string `json:"support_status_name"`
}

func ToOutSupportStatus(status *internal_models.Status) OutSupportStatus {
	return OutSupportStatus{
		ID:   status.ID,
		Name: status.Name,
	}
}
