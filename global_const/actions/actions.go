package actions

const (
	AdminTA string = "administration" //для проверки доступа на урл

)

const (
	AdminTA_UserUpdate string = "administration.user.update" //изменение пользователя
	AdminTA_UserGet    string = "administration.user.get"    //получение пользователя

	AdminTA_GroupCreate string = "administration.group.create" //создание группы пользователей
	AdminTA_GroupUpdate string = "administration.group.update" //изменение группы пользователей
	AdminTA_GroupGet    string = "administration.group.get"    //получение групп пользователей

)

const (
	TicketTA_Create            string = "ticket.create"              //создание запроса
	TicketTA_Update            string = "ticket.update"              //обновление запроса
	TicketTA_Get               string = "ticket.get"                 //получить запрос
	TicketTA_Filtered          string = "ticket.get.filtered"        //получить отфильтрованный список запросов
	TicketTA_Work              string = "ticket.work"                //обрабатывать запрос
	TicketTA_FullSearch        string = "ticket.full_search"         //расширенный поиск
	TicketTA_Resolve           string = "ticket.resolve"             //согласование запроса
	TicketTA_SeeAdditionalInfo string = "ticket.see_additional_info" //смотреть доп инфо по запросу
)

const (
	ReportTA_Get string = "report.get"
)
