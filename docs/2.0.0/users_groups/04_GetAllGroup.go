package users

/**
 *
 * @api {GET} /helpdesk/group/ 04. Получение списка групп прав
 * @apiName GetAllGroup
 * @apiGroup 04. Пользователи
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 * @apiDescription Возвращает массив объектов
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {Group[]} 		group 						Массив объектов "группа пользователя"
 * @apiSuccess (Success 200) {Uint64} 		group.group_id 				ID группы
 * @apiSuccess (Success 200) {String} 		group.group_name 			Имя группы
 * @apiSuccess (Success 200) {Bool} 		group.create_ticket 		Разрешение принимать запросы
 * @apiSuccess (Success 200) {Bool} 		group.get_all_tickets 		Разрешение получать список запросов
 * @apiSuccess (Success 200) {Bool} 		group.see_additional_info 	Разрешение смотреть расширенную информацию по запросу
 * @apiSuccess (Success 200) {Bool} 		group.can_resolve_ticket 	Разрешение согласовывать запросы
 * @apiSuccess (Success 200) {Bool} 		group.work_on_tickets 		Разрешение работать над запросом
 * @apiSuccess (Success 200) {Bool} 		group.change_settings 		Разрешение изменять данные в базе
 *
 * @apiSuccessExample {json} Success-Response:
 * [
 *   {
 *       "group_id": 1,
 *       "group_name": "regular_user",
 *       "create_ticket": true,
 *       "get_all_tickets": false,
 *       "see_additional_info": false,
 *       "can_resolve_ticket": false,
 *       "work_on_tickets": false,
 *       "change_settings": false
 *   },
 *   {
 *       "group_id": 2,
 *       "group_name": "support",
 *       "create_ticket": false,
 *       "get_all_tickets": false,
 *       "see_additional_info": false,
 *       "can_resolve_ticket": false,
 *       "work_on_tickets": false,
 *       "change_settings": false
 *   }
 * ]
 *
 */
