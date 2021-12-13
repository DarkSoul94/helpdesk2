package docs

/**
 *
 * @api {POST} /helpdesk/service/auto_create 03. Автоматическое создание запроса
 * @apiName AutoCreateTicket
 * @apiGroup 03. Запросы в тех. поддержку
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 * @apiDescription Автоматическое создание запроса из данных присланных по API сторонними сервисами.
 *
 * @apiParam  {String} 	text 			Текст запроса
 * @apiParam  {String} 	user_email 		Доменная почта пользователя
 * @apiParam  {String} 	[user_ip] 		IP-адресс компьютера с которого отправлялся запрос
 * @apiParam  {Bool} 	priority 		Признак являеться ли запрос приоритетным
 *
 * @apiSuccess (200) {String} 	status 		статус выполнения запроса
 * @apiSuccess (200) {Int} 		ticket_id 	номер созданного запроса(0 - если была ошибка)
 *
 * @apiParamExample  {json} Request-Example:
 * {
 *    "text":"api ticket",
 *    "user_email" : "tishchenko.v.v@limefin.com",
 *    "user_ip":"10.54.1.101",
 *	  "priority" : true
 * }
 *
 *
 * @apiSuccessExample {json} Запрос успешно создан:
 * {
 *   "status": "ok",
 *   "ticket_id": 49
 * }
 *
 * @apiSuccessExample {json} Запрос без email:
 * {
 * 	 "error": "Email is blank",
 * 	 "status": "error"
 * }
 *
 *
 */
