package docs

/**
 *
 * @api {POST} /helpdesk/support/open_shift 07. Открытие смены сотруднику ТП
 * @apiName OpenShift
 * @apiGroup 05. Сотрудник ТП
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 *
 * @apiHeader {String}	BearerToken 	Авторизационный токен
 *
 * @apiParam  {Uint64} 	support_id	ID сотрудника которому нужно открыть смену
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
