package docs

/**
 *
 * @api {GET} /helpdesk/resolve_ticket/resolve_tickets_list 06. Получение списка запросов в тех. поддержку ожидающих согласования
 * @apiName GetResolveTicketsList
 * @apiGroup 03. Запросы в тех. поддержку
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiParam {Int}	count	Количество запросов которые сервер должен вернуть
 * @apiParam {Int}	offset	Сдвиг по списку запросов (общее колчество запросов которые клиент уже получил)
 *
 * @apiExample  Example usage:
 * http://localhost:8888/helpdesk/resolve_ticket/resolve_tickets_list?count=5&offset=0
 *
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {String[]}         fields                              Список полей которые должны отрисоваться на стороне фронта
 * @apiSuccess (Success 200) {Ticket[]}         tickets                             Массив запросов в тех. поддержку
 * @apiSuccess (Success 200) {Uint64} 	       	tickets.ticket_id 					ID запроса
 * @apiSuccess (Success 200) {String} 	       	tickets.ticket_date 				Дата создания запроса
 * @apiSuccess (Success 200) {String}        	tickets.category   					Категория
 * @apiSuccess (Success 200) {String}  			tickets.section            			Раздел категории
 * @apiSuccess (Success 200) {String} 	        tickets.ticket_text 				Текст запроса
 * @apiSuccess (Success 200) {String}     		tickets.status 						Статус запроса
 * @apiSuccess (Success 200) {String}     		tickets.ticket_author               Автор запроса
 * @apiSuccess (Success 200) {String}     		tickets.filial               		Отделение за которым закрелен автор запроса
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *     "fields": [
 *         	"ticket_id",
 *         	"ticket_date",
 *         	"category",
 *         	"section",
 *         	"ticket_text",
 *         	"status",
 *			"filial",
 *         	"ticket_author",
 *     ],
 *     "tickets": [
 *        {
 *            "ticket_id": 51,
 *            "ticket_date": "2021-05-26T11:40:41+03:00",
 *            "category": "Оборудование",
 *            "section": "Настройка интернета",
 *            "ticket_text": "adasdasdads",
 *            "status": "В ожидании согласования",
 *            "filial": "not found",
 *            "ticket_author": "Артем Владимирович Шелкопляс"
 *        },
 *        {
 *            "ticket_id": 49,
 *            "ticket_date": "2021-05-26T10:33:09+03:00",
 *            "category": "1С",
 *            "section": "Удаление кассовых",
 *            "ticket_text": "gdfgdfbdfbdfb",
 *            "status": "В ожидании согласования",
 *            "filial": "not found",
 *            "ticket_author": "Вячеслав Викторович Тищенко"
 *        },
 *     ]
 * }
 */
