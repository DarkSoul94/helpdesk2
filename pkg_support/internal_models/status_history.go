package internal_models

import "time"

type StatusHistory struct {
	ID         uint64
	Support    *Support
	SelectTime time.Time
	Duration   time.Duration
	ShiftID    uint64
}

func (sh *StatusHistory) New(support *Support, shiftID uint64) {
	*sh = StatusHistory{
		Support:    support,
		SelectTime: time.Now().Truncate(time.Second),
		Duration:   0,
		ShiftID:    shiftID,
	}
}

func (sh *StatusHistory) SetDuration() {
	sh.Duration = time.Duration(time.Since(sh.SelectTime).Seconds())
}
