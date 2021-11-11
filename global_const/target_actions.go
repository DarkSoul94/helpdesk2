package global_const

//группа целевых действий по администрировании системы
const (
	AdminTA             string = "administration"              //для проверки доступа на урл
	AdminTA_UserUpdate  string = "administration.user.update"  //изменение пользователя
	AdminTA_GroupCreate string = "administration.group.create" //создание группы пользователей
	AdminTA_GroupUpdate string = "administration.group.update" //изменение группы пользователей
	AdminTA_GroupGet    string = "administration.group.get"    //получение групп пользователей

	TicketTA_FullSearch string = "ticket.full_search" //расширенный поиск
	TicketTa_Create     string = "ticket.create"      // создание
)
