package dto

import "github.com/DarkSoul94/helpdesk2/pkg_scheduler/internal_models"

type OutOffice struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Deleted bool   `json:"deleted"`
}

func ToModelOffices(shifts []OutOffice) []*internal_models.Office {
	var mShifts = make([]*internal_models.Office, 0)

	for _, shift := range shifts {
		mShift := &internal_models.Office{
			ID:      shift.ID,
			Name:    shift.Name,
			Color:   shift.Color,
			Deleted: shift.Deleted,
		}

		mShifts = append(mShifts, mShift)
	}

	return mShifts
}

func ToOutOffices(mShift []*internal_models.Office) []OutOffice {
	var oShifts = make([]OutOffice, 0)

	for _, shift := range mShift {

		oShifts = append(oShifts, OutOffice{
			ID:      shift.ID,
			Name:    shift.Name,
			Color:   shift.Color,
			Deleted: shift.Deleted,
		})
	}
	return oShifts
}
