package global_const

//группа целевых действий связанные с задачами
const (
	ActionTaskGet          = "task.get"          //получение списка задач
	ActionTaskGet_author   = "task.get.author"   //получение списка задач по автору
	ActionTaskGet_dev      = "task.get.dev"      //получение списка задач по разработчику
	ActionTaskGet_customer = "task.get.customer" //получение списка задач по заказчику

	ActionTaskCreate = "task.create" //создание новой задачи

	ActionTaskUpdate = "task.update" //обновление задачи
)

//группа целевых действий связанных с изменениями настроек системы, групп, пользователей и т.д.
const (
	SettingsUserUpdate = "settings.user.update" //обновление пользователя

	SettingsGroupUpdate = "settings.group.update" //обновление группы пользователей
	SettingsGroupCreate = "settings.group.create" //создание группы пользователей
)

var ActionsForPerm []string = []string{
	ActionTaskGet_author,
	ActionTaskGet_dev,
	ActionTaskGet_customer,
	ActionTaskCreate,
	ActionTaskUpdate,
	SettingsUserUpdate,
	SettingsGroupUpdate,
	SettingsGroupCreate,
}
