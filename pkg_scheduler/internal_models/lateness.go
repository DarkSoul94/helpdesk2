package internal_models

import "time"

type Lateness struct {
	ID          uint64
	Date        time.Time
	SupportID   uint64
	SupportName string
	Cause       string
	Decision    *Decision
	Difference  uint64
}

func (l *Lateness) New(supportID uint64, start time.Time) {
	date := time.Now().Truncate(time.Minute)
	time, _ := time.Parse("15:04:05", date.Format("15:04:05"))
	dif := time.Sub(start).Minutes()
	*l = Lateness{
		Date:       date,
		SupportID:  supportID,
		Difference: uint64(dif),
		Decision:   &Decision{},
	}
}
