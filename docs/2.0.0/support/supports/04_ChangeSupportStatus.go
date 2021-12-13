package docs

/**
 *
 * @api {POST} /helpdesk/support/change_status 04. Сменить статус сотрудника ТП
 * @apiName ChangeSupportStatus
 * @apiGroup 05. Сотрудник ТП
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 * @apiDescription 	Смена рабочего статуса сотрудника ТП: "принимаю запросы", "не принимаю запросы" и т.д.
 * 									Используется как для смены статуса самому себе, так и для смены статуса другому сотруднику.
 * 									Сменить статус другому сотруднику может только админ.
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam {Uint64}	support_id 			ID сотрудника ТП. Если ID сотрудника не передается, то ID берется из авторизационного токена
 * @apiParam {Uint64}	support_status_id	ID нового статуса сотрудника ТП
 *
 * @apiParamExample  {json} Смена статуса другому суппорту:
 * {
 *		"support_id": 4,
 *		"support_status_id": 1
 * }
 *
 * @apiParamExample  {json} Смена статуса себе:
 * {
 *         "support_status_id": 1
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
