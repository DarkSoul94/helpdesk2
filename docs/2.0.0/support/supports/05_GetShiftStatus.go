package docs

/**
 *
 * @api {GET} /helpdesk/support/shift_status 05. Получение текущего статуса смены
 * @apiName GetSupportShiftStatus
 * @apiGroup 05. Сотрудник ТП
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {Bool} 		shift_status 	Статус смены. True - открыта
 * @apiSuccess (Success 200) {String} 		status 			Статус ответа
 *
 * @apiSuccessExample {json} Смена закрыта:
 * {
 *   "shift_status": false,
 *   "status": "ok"
 * }
 *
 * @apiSuccessExample {json} Смена открыта:
 * {
 *   "shift_status": true,
 *   "status": "ok"
 * }
 *
 */
