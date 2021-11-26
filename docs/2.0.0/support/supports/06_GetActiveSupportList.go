package docs

/**
 *
 * @api {GET} /helpdesk/support/active_support_list 06. Получение списка активных сотрудников ТП
 * @apiName GetActiveSupportList
 * @apiGroup 05. Сотрудник ТП
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {ActiveSupport[]} 	active_support 				Массив объектов "Активный сотрудник ТП"
 * @apiSuccess (Success 200) {Uint64} 			active_support.user_id 		ID сотрудника
 * @apiSuccess (Success 200) {String} 			active_support.user_name 	ФИО сотрудника
 *
 * @apiSuccessExample {json} Success-Response:
 * [
 *     {
 *         "user_id": 5,
 *         "user_name": "Евгений Николаевич Табаков"
 *     },
 *     {
 *         "user_id": 4,
 *         "user_name": "Вячеслав Викторович Тищенко"
 *     }
 * ]
 *
 *
 */
