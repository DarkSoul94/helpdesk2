package internal_models

type UserForGenerate struct {
	UserID uint64
	Count  int
}

type TicketGenerate struct {
	Text      string
	SectionID uint64
	Users     []UserForGenerate
}
