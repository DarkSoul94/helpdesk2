package internal_models

import "time"

type StatusHistory struct {
	ID         uint64
	Support    *Support
	SelectTime time.Time
	Duration   time.Duration
	Shift      *Shift
}

func (sh *StatusHistory) New(support *Support, shiftID uint64) {
	*sh = StatusHistory{
		Support:    support,
		SelectTime: time.Now().Truncate(time.Second),
		Duration:   0,
		Shift: &Shift{
			ID: shiftID,
		},
	}
}

func (sh *StatusHistory) SetDuration() {
	sh.Duration = time.Duration(time.Since(sh.SelectTime).Seconds())
}
