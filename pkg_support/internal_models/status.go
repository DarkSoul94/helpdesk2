package internal_models

type Status struct {
	ID           uint64
	Name         string
	AcceptTicket bool
}

func (s *Status) accept() {
	s.ID = 1
	s.AcceptTicket = true
}

func (s *Status) notAccept() {
	s.ID = 4
	s.AcceptTicket = false
}
