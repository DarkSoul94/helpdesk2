package internal_models

import "time"

type Shift struct {
	ID            uint64
	Support       *Support
	OpeningTime   time.Time
	ClosingTime   time.Time
	ClosingStatus bool
}
