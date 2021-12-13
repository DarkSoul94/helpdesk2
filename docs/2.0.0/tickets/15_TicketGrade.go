package docs

/**
 *
 * @api {POST} /helpdesk/ticket/ticket_grade 15. Оценка запроса в тех.поддержку
 * @apiName CreateTicketGrade
 * @apiGroup 03. Запросы в тех. поддержку
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam {Uint64} 	ticket_id 		ID оцениваемого запроса
 * @apiParam {Uint64} 	ticket_grade 	Оценка
 *
 * @apiSuccess (Success 200) {String} status 			Статус ответа
 * @apiSuccess (Success 200) {String} ticket_grade_id 	ID созданного объекта
 *
 * @apiParamExample  {json} Request-Example:
 * {
 *   "ticket_id":5,
 *   "ticket_grade":5
 * }
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *    "status": "ok",
 *    "ticket_grade_id": "1"
 * }
 *
 *
 */
