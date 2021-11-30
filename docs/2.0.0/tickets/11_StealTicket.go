package docs

/**
 *
 * @api {POST} /helpdesk/ticket/steal 11. Взять чужой запрос себе в работу.
 * @apiName StealTicket
 * @apiGroup 03. Запросы в тех. поддержку
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam {Uint64}	ticket_id	ID запроса
 *
 * @apiSuccess (Success 200) {String} status Статус ответа на запрос
 *
 * @apiParamExample  {json} Request-Example:
 * {
 *   "ticket_id":5
 * }
 *
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *     "status": "ok"
 * }
 *
 * @apiSuccessExample {json} Error-Response:
 * {
 *      "status": "error",
 * 		"error":"Ticket is complete"
 * }
 *
 * @apiError ErrTicketIsComplete Ticket is complete
 * @apiError ErrTicketDoesNotExist Ticket with this id not exist
 *
 */
