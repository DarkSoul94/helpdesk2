package scheduler

/**
 *
 * @api {POST} /helpdesk/table/lateness_conf 08. Изменить настройки опозданий
 * @apiName SetLatenessConf
 * @apiGroup 8. Табель
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam  {Uint64} late_penalty 	Штраф за минуту опоздания
 * @apiParam  {Uint64} grace_time 		Количество льготных минут
 *
 * @apiSuccess (200) {String} status 	Cтатус выполнения
 *
 * @apiParamExample  {Json} Request-Example:
 * {
 *     "late_penalty":5,
 *     "grace_time":20
 * }
 *
 *
 * @apiSuccessExample {Json} Success-Response:
 * {
 *     "status": "ok"
 * }
 *
 *
 */
