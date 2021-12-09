package users

/**
 *
 * @api {POST} /helpdesk/group/create 03. Создание группы прав пользователей
 * @apiName CreateGroup
 * @apiGroup 04. Пользователи
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam {String} 		group_name 				Имя группы
 * @apiParam {Bool} 		create_ticket 			Разрешение принимать запросы
 * @apiParam {Bool} 		get_all_tickets 		Разрешение получать список запросов
 * @apiParam {Bool} 		see_additional_info 	Разрешение смотреть расширенную информацию по запросу
 * @apiParam {Bool} 		can_resolve_ticket 		Разрешение согласовывать запросы
 * @apiParam {Bool} 		work_on_tickets 		Разрешение работать над запросом
 * @apiParam {Bool} 		change_settings 		Разрешение изменять данные в базе
 *
 * @apiSuccess (Success 200) {String} 	status 		Статус ответа на запрос
 * @apiSuccess (Success 200) {String}	group_id	ID созданной группы
 *
 * @apiParamExample  {json} Request-Example:
 * {
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
 *     "group_id": 3,
 *     "status": "ok"
 * }
 *
 * @apiError ErrGroupAlreadyExist Такая группа уже существует
 *
 */
