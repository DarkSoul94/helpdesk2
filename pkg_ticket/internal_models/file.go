package internal_models

import "time"

type File struct {
	ID        uint64
	Name      string
	Date      time.Time
	Data      string
	Extension string
	Path      string
	TicketId  uint64
}
