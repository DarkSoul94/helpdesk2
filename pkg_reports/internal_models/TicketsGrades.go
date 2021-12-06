package internal_models

type TicketGrade struct {
	TicketID    uint64
	TicketGrade uint
}

type UserTicketGrades struct {
	UserName         string
	TicketsGrades    []TicketGrade
	AverageUserGrade float64
}

type DepartmentTicketGrade struct {
	Department  string
	UsersGrades []UserTicketGrades
}
