package docs

/**
*
* @api {GET} /helpdesk/ticket/tickets_list Получение списка запросов в тех. поддержку
* @apiName GetTicketsList
* @apiGroup 03. Запросы в тех. поддержку
* @apiVersion  2.0.0
* @apiSampleRequest off
* @apiDescription В зависимости от прав закрепленных за группой пользователей будет отличатся результирующий список запросов.
* В случае если у пользователя есть разрешение `get_all_tickets` - получит весь список запросов,
* при `can_resolve_ticket` - получит написанные собой запросы, а также запросы на согласование,
* при `work_on_tickets` - получит распределенные на него запросы. В случае если нет ни одного из этих разрешений пользователь получит список только написанных им запросов.
* Поле `support` отображается только для пользователя с правами `admin`.
*
* @apiParam {Int}	count	Количество запросов которые сервер должен вернуть
* @apiParam {Int}	offset	Сдвиг по списку запросов (общее колчество запросов которые клиент уже получил)
*
* @apiExample  Example usage:
* http://localhost:8888/helpdesk/ticket/tickets_list?count=5&offset=0
*
*
* @apiHeader {String} BearerToken 	Авторизационный токен
*
* @apiSuccess (Success 200) {String[]}         	fields                              Список полей которые должны отрисоваться на стороне фронта
* @apiSuccess (Success 200) {Ticket[]}         	tickets                             Массив запросов в тех. поддержку
* @apiSuccess (Success 200) {Uint64} 	       	tickets.ticket_id 					ID запроса
* @apiSuccess (Success 200) {String} 	       	tickets.ticket_date 				Дата создания запроса
* @apiSuccess (Success 200) {Bool} 	       		tickets.significant 				Важность запроса
* @apiSuccess (Success 200) {String}        	tickets.category   					Категория
* @apiSuccess (Success 200) {String}  			tickets.section            			Раздел категории
* @apiSuccess (Success 200) {String} 	        tickets.ticket_text 				Текст запроса
* @apiSuccess (Success 200) {String}     		tickets.status 						Статус запроса
* @apiSuccess (Success 200) {String}     		tickets.ticket_author               Автор запроса
* @apiSuccess (Success 200) {String}     		tickets.support               		Сотрудник ТП работающий над запросом
* @apiSuccess (Success 200) {String}     		tickets.filial               		Отделение за которым закрелен автор запроса
* @apiSuccess (Success 200) {Uint}     			tickets.grade               		Оценка запроса
* @apiSuccess (Success 200) {Uint}     			tickets.sort_priority               Приоритет отображения запроса (1 - наивысший приоритет). В рамках одного приоритета запросы сортируются по своим ИД по убыванию.
*
* @apiSuccessExample {json} Ответ при запросе списка админом:
* {
*     "fields": [
*			"ticket_id",
*			"ticket_date",
*			"category",
*			"section",
*			"ticket_text",
*			"status",
*			"filial",
*			"ticket_author",
*			"support",
*			"grade"
*     ],
*     "tickets": [
*		{
*			"ticket_id": 68,
*			"ticket_date": "2021-07-28T06:55:17Z",
*			"significant": false,
*			"category": "Оборудование",
*			"section": "Проблема с принтером/печатью",
*			"ticket_text": "afasdasdqds",
*			"status": "Отклонен",
*			"ticket_status_id": 8,
*			"filial": "not found",
*			"ticket_author": "Евгений Николаевич Табаков",
*			"support": "Вячеслав Викторович Тищенко",
*			"grade": 0,
*			"sort_priority": 1
*		}
*     ]
* }
*
* @apiSuccessExample {json} Ответ при запросе списка остальными пользователями:
* {
*     "fields": [
*			"ticket_id",
*			"ticket_date",
*			"category",
*			"section",
*			"ticket_text",
*			"status",
*			"filial",
*			"ticket_author",
*			"grade"
*     ],
*     "tickets": [
*         {
*			"ticket_id": 68,
*			"ticket_date": "2021-07-28T06:55:17Z",
*			"significant": false,
*			"category": "Оборудование",
*			"section": "Проблема с принтером/печатью",
*			"ticket_text": "afasdasdqds",
*			"status": "Отклонен",
*			"ticket_status_id": 8,
*			"filial": "not found",
*			"ticket_author": "Евгений Николаевич Табаков",
*			"support": "",
*			"grade": 0,
*			"sort_priority": 4
*         }
*     ]
* }
*
*
 */
