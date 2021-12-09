package users

/**
 *
 * @api {POST} /helpdesk/group/update 05. Обновление данных группы прав
 * @apiName UpdateGroup
 * @apiGroup 04. Пользователи
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam {Uint64} 		group_id 				ID группы
 * @apiParam {String} 		group_name 				Имя группы
 * @apiParam {Bool} 		create_ticket 			Разрешение принимать запросы
 * @apiParam {Bool} 		get_all_tickets 		Разрешение получать список запросов
 * @apiParam {Bool} 		see_additional_info 	Разрешение смотреть расширенную информацию по запросу
 * @apiParam {Bool} 		can_resolve_ticket 		Разрешение согласовывать запросы
 * @apiParam {Bool} 		work_on_tickets 		Разрешение работать над запросом
 * @apiParam {Bool} 		change_settings 		Разрешение изменять данные в базе
 *
 * @apiSuccess (Success 200) {string} status Статус ответа на запрос
 *
 * @apiParamExample  {json} Request-Example:
 * {
 *   "group_id": 2,
 *   "group_name": "support",
 *   "create_ticket": false,
 *   "get_all_tickets": false,
 *   "see_additional_info": false,
 *   "can_resolve_ticket": false,
 *   "work_on_tickets": false,
 *   "change_settings": false
 * }
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "status": "ok"
 * }
 *
 * @apiError GroupErr_NotExist Такой группы не существует
 *
 */
