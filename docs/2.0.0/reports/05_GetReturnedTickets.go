package docs

/**
 *
 * @api {GET} /helpdesk/reports/returned_tickets 05. Список запросов возвращенных обратно в работу после статус выполнено или отклонено
 * @apiName GetReturnedTickets
 * @apiGroup 06. Отчёты
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken Авторизационный токен
 *
 * @apiParam  {String} start_date 	Дата начала выборки, включительно
 * @apiParam  {String} end_date		Дата конца выборки, данные за этот день не учитываются
 *
 * @apiExample  Example usage:
 * http://localhost:8888/helpdesk/reports/returned_tickets?start_date=2021-05-01&end_date=2021-05-07
 *
 * @apiSuccess (Success 200) {Ticket[]} ticket 					Массив объектов "запрос"
 * @apiSuccess (Success 200) {Uint64} 	ticket.ticket_id 		Номер запроса
 * @apiSuccess (Success 200) {String} 	ticket.ticket_date 		Дата создания запроса
 * @apiSuccess (Success 200) {String} 	ticket.category 		Категория запроса
 * @apiSuccess (Success 200) {String} 	ticket.section 			Раздел категории запроса
 * @apiSuccess (Success 200) {String} 	ticket.ticket_text 		Текст запроса
 * @apiSuccess (Success 200) {String} 	ticket.status 			Текущий статус запроса
 * @apiSuccess (Success 200) {String} 	ticket.author 			ФИО автора запроса
 * @apiSuccess (Success 200) {String} 	ticket.support 			ФИО сотрудника тех.поддержки
 * @apiSuccess (Success 200) {Uint64} 	ticket.ticket_grade 	Оценка запроса
 *
 *
 * @apiSuccessExample {json} Success-Response:
 * [
 *   {
 *     "ticket_id": 23,
 *     "ticket_date": "2021-06-11 15:27:37",
 *     "category": "Оборудование",
 *     "section": "Проблема с принтером/печатью",
 *     "ticket_text": "Не печатает принтер",
 *     "status": "В работе",
 *     "author": "Евгений Николаевич Табаков",
 *     "support": "Артем Владимирович Шелкопляс",
 *     "ticket_grade": 0
 *   },
 *   {
 *     "ticket_id": 19,
 *     "ticket_date": "2021-06-09 16:39:54",
 *     "category": "1С",
 *     "section": "Изменение/удаление кассовых ордеров",
 *     "ticket_text": "удалить кассовый ордер",
 *     "status": "Выполнен",
 *     "author": "Евгений Николаевич Табаков",
 *     "support": "Артем Владимирович Шелкопляс",
 *     "ticket_grade": 4
 *   }
 * ]
 *
 */
