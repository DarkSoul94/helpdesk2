package docs

/**
 *
 * @api {GET} /helpdesk/support/get_support_status 03. Получение текущего статуса сотрудника ТП
 * @apiName GetSupportStatus
 * @apiGroup 05. Сотрудник ТП
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {Uint64} 			support_status_id 		ID статуса
 * @apiSuccess (Success 200) {String} 			support_status_name 	Описание статуса
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "support_status_id": 4,
 *   "support_status_name": "Не принимаю запросы"
 * }
 *
 *
 */
