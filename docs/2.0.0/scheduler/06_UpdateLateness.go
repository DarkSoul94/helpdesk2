package scheduler

/**
 *
 * @api {POST} /helpdesk/table/lateness/update 06. Установить решение по опозданию
 * @apiName UpdateLateness
 * @apiGroup 8. Табель
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam  {Uint64} lateness_id ID записи об опоздании
 * @apiParam  {Uint64} decision_id ID решения по опозданию
 *
 * @apiSuccess (Success 200) {String} status	Статус выполнения запроса
 *
 * @apiParamExample  {json} Request-Example:
 * {
 *     "lateness_id": 1,
 *     "decision_id": 1
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
