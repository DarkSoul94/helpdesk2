package docs

/**
 *
 * @api {POST} /helpdesk/resolve_ticket/resolve Согласование запроса
 * @apiName ResolveTicket
 * @apiGroup 03. Запросы в тех. поддержку
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam {Uint64}	ticket_id	ID запроса который согласовывается.
 *
 * @apiSuccess (Success 200) {String} status Статус ответа на запрос
 *
 * @apiParamExample  {json} Request-Example:
 * {
 *   "ticket_id":5
 * }
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *     "status": "ok"
 * }
 *
 * @apiError ErrTicketDoesNotExist Запроса с таким ID не существует
 * @apiError ErrDoesNotNeedApproval Запрос не нуждается в согласовании
 *
 */
