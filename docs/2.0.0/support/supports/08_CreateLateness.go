package docs

/**
 *
 * @api {POST} /helpdesk/support/create_lateness 08. Отправка причины опоздания
 * @apiName CreateLateness
 * @apiGroup 05. Сотрудник ТП
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 *
 * @apiHeader {String}	BearerToken 	Авторизационный токен
 *
 * @apiParam  {Uint64} 	support_id	ID сотрудника которому нужно открыть смену
 * @apiParam  {String} 	cause				Причина опоздания
 *
 * @apiSuccess (Success 200) {String} 		status 			Статус ответа
 *
 * @apiParamExample  {json} Request-Example:
 * {
 *     "support_id" : 4,
 *     "cause": "test"
 * }
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "status": "ok"
 * }
 *
 *
 */
