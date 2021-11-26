package docs

/**
 *
 * @api {GET} /support/status_list 02. Получение списка возможных статусов для работников ТП
 * @apiName GetStatusesForSupport
 * @apiGroup 05. Сотрудник ТП
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {SupportStatus[]} 	support_status 						Массив объектов "статус сотрудника ТП"
 * @apiSuccess (Success 200) {Uint64} 			support_status.support_status_id 	ID статуса
 * @apiSuccess (Success 200) {String} 			support_status.support_status_name 	Описание статуса
 *
 * @apiSuccessExample {json} Success-Response:
 * [
 *   {
 *       "support_status_id": 1,
 *       "support_status_name": "Принимаю запросы"
 *   },
 *   {
 *       "support_status_id": 2,
 *       "support_status_name": "Перерыв"
 *   },
 *   {
 *       "support_status_id": 3,
 *       "support_status_name": "Работа в офисе"
 *   },
 *   {
 *       "support_status_id": 4,
 *       "support_status_name": "Не принимаю запросы"
 *   }
 * ]
 *
 *
 */
