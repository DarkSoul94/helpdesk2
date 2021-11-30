package docs

/**
 *
 * @api {POST} /helpdesk/ticket/generate_tickets 02. Создание большого количества запросов
 * @apiName GenerateTickets
 * @apiName CreateTicket
 * @apiGroup 03. Запросы в тех. поддержку
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam  {String} 	text 			Текст запроса
 * @apiParam  {Int} 	section_id 		ИД раздела категории
 * @apiParam  {User[]} 	users 			Массив обектов "пользователь"
 * @apiParam  {Int} 	users.user_id 	ИД пользователя которому создать запрос
 * @apiParam  {Int} 	users.count 	Количество запросов которые необходимо создать
 *
 * @apiSuccess (Success 200) {String} status 	Статус ответа на запрос
 *
 * @apiParamExample  {json} Request-Example:
 * {
 *   "text":"тест много запросов3",
 *   "section_id":1,
 *   "users":[
 *       	{
 *           "user_id": 6,
 *           "count": 1
 *       	}
 *       ]
 * }
 *
 * @apiSuccessExample {json} Запросы успешно созданы:
 * {
 *     "status": "ok"
 * }
 *
 * @apiSuccessExample {json} Пустой текст запроса:
 * {
 *  "error": "Ticket text is blank",
 *  "status": "error"
 * }
 *
 * @apiSuccessExample {json} ИД не существующей категории:
 * {
 *  "error": "Such category section doesn't exist",
 *  "status": "error"
 * }
 *
 *
 */
