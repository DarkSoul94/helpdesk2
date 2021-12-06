package internal_models

var (
	main Office = Office{
		ID:      1,
		Name:    "Главный офис",
		Deleted: false,
	}
)

type Office struct {
	ID      uint64
	Name    string
	Color   string
	Deleted bool
}

func (s *Office) Compare(shift *Office) bool {
	s.checkForMain()
	if s.Name != shift.Name {
		return false
	}
	if s.Color != shift.Color {
		return false
	}
	if s.Deleted != shift.Deleted {
		return false
	}

	return true
}

func (s *Office) checkForMain() {
	if s.ID == main.ID {
		s.Name = main.Name
		s.Deleted = main.Deleted
	}
}
