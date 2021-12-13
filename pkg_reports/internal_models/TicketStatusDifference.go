package internal_models

type StatusDifference struct {
	StatusName string
	Duration   string
}

type TicketDifference struct {
	TicketID    uint64
	SupportName string
	Section     string
}
