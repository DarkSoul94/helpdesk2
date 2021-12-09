package scheduler

/**
 *
 * @api {GET} /helpdesk/table/lateness_conf 07. Получить настройки опозданий
 * @apiName GetLatenessConf
 * @apiGroup 8. Табель
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (200) {LatenessConf} lateness_conf 								Объект "настройки опозданий"
 * @apiSuccess (200) {Uint64} 			lateness_conf.late_penalty 		Штраф за минуту опоздания
 * @apiSuccess (200) {Uint64} 			lateness_conf.grace_time 			Количество льготных минут
 * @apiSuccess (200) {String} 			status 												Cтатус выполнения
 *
 * @apiSuccessExample {type} Success-Response:
 * {
 *   "lateness_conf": {
 *     "late_penalty": 5,
 *     "grace_time": 20
 *   },
 *   "status": "ok"
 * }
 *
 *
 */
