package users

/**
 *
 * @api {GET} /helpdesk/group/for_resolve 06. Получение списка груп пользователей с правами согласовывать запросы
 * @apiName GetGroupsListForResolve
 * @apiGroup 04. Пользователи
 * @apiVersion  2.0.0
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {String} 	status 							Статус запроса
 * @apiSuccess (Success 200) {[]Groups} groups 							Массив объектов групп
 * @apiSuccess (Success 200) {Uint64} 	groups.group_id 		ID группы
 * @apiSuccess (Success 200) {String} 	groups.group_name		Название группы
 *
 *
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "groups": [
 *     {
 *       "group_id": 4,
 *       "group_name": "Сотрудник бэк-офиса"
 *     }
 *   ],
 *   "status": "ok"
 * }
 *
 *
 */
