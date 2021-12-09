package users

/**
 *
 * @api {POST} /helpdesk/user/update 02. Обновление данных пользователя
 * @apiName UpdateUser
 * @apiGroup 04. Пользователи
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam {Uint64}	user_id 			ID пользователя1
 * @apiParam {Uint64} 	group_id			ID группы
 *
 * @apiSuccess (Success 200) {String}    status  Статус ответа на запрос
 *
 * @apiParamExample  {json} Request-Example:
 *{
 *   "user_id": 2,
 *   "group_id": 1
 *}
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *  	"status": "ok"
 * }
 *
 *
 */
