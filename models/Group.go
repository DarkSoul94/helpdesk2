package models

type Group struct {
	ID                uint64
	Name              string
	CreateTicket      bool
	GetAllTickets     bool
	SeeAdditionalInfo bool
	CanResolveTicket  bool
	WorkOnTickets     bool
	ChangeSettings    bool
	CanReports        bool
	FullSearch        bool
	//tree              treeLayer
}
