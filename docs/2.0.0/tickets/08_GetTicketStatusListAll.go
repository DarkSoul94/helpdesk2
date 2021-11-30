package docs

/**
 *
 * @api {get} /helpdesk/ticket_status/  08. Получение списка статусов запроса в ТП для выбора внутри запроса
 * @apiName GetTicketStatus
 * @apiGroup 03. Запросы в тех. поддержку
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 * @apiDescription Возвращает массив объектов
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {TicketStatus[]} 	ticket_status 						Массив объеквтов "статус запроса"
 * @apiSuccess (Success 200) {Uint64} 			ticket_status.ticket_status_id 		ID статуса запроса
 * @apiSuccess (Success 200) {String} 			ticket_status.ticket_status_name 	Имя статуса запроса
 *
 * @apiSuccessExample {json} Success-Response:
 * [
 *   {
 *       "ticket_status_id": 1,
 *       "ticket_status_name": "Новый"
 *   },
 *   {
 *       "ticket_status_id": 2,
 *       "ticket_status_name": "В ожидании"
 *   }
 * ]
 *
 *
 */
