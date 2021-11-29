package docs

/**
 *
 * @api {POST} /helpdesk/support/close_shift 09. Закрытие смены сотруднику ТП
 * @apiName CloseShift
 * @apiGroup 05. Сотрудник ТП
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam  {Uint64} support_id ID сотрудника которому нужно закрыть смену
 *
 * @apiSuccess (Success 200) {String} 		status 			Статус ответа
 *
 * @apiParamExample  {json} Request-Example:
 * {
 *     "support_id" : 6
 * }
 *
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "status": "ok"
 * }
 *
 *
 */
