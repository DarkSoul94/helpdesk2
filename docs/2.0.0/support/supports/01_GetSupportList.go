package docs

/**
 *
 * @api {GET} /helpdesk/support/support_list 01. Получение списка всех сотрудников ТП
 * @apiName GetSupportList
 * @apiGroup 05. Сотрудник ТП
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiExample  Example usage:
 * http://localhost:8888/helpdesk/support/support_list
 *
 * @apiHeader {String} BearerToken Авторизационный токен
 *
 * @apiSuccess (Success 200) {Support[]} 	support 						Массив объектов "Cотрудник ТП"
 * @apiSuccess (Success 200) {Uint64} 		support.user_id 		ID сотрудника
 * @apiSuccess (Success 200) {String} 		support.user_name 	ФИО сотрудника
 *
 *
 * @apiSuccessExample {type} Success-Response:
 * [
 *   {
 *     "user_id": 4,
 *     "user_name": "Евгений Николаевич Табаков"
 *   },
 *   {
 *     "user_id": 5,
 *     "user_name": "Артем Владимирович Шелкопляс"
 *   },
 *   {
 *     "user_id": 6,
 *     "user_name": "Вячеслав Викторович Тищенко"
 *   },
 *   {
 *     "user_id": 7,
 *     "user_name": "Александр Игоревич Кудряшов"
 *   }
 * ]
 *
 *
 */
