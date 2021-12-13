package users

/**
 *
 * @api {GET} /helpdesk/user/ 01. Получение списка пользователей
 * @apiName GetAllUsers
 * @apiGroup 04. Пользователи
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 * @apiDescription Возвращает массив объектов
 *
 * @apiHeader {String} BearerToken Авторизационный токен
 *
 * @apiSuccess (Success 200) {User[]} 	user 					Массив объектов пользователь
 * @apiSuccess (Success 200) {Uint64} 	user.user_id 			ID пользователя
 * @apiSuccess (Success 200) {String} 	user.user_name			Имя пользователя
 * @apiSuccess (Success 200) {String} 	user.email				Электронная почта пользователя
 * @apiSuccess (Success 200) {Group} 	user.group				Группа в которой состоит пользователь
 * @apiSuccess (Success 200) {Uint64} 	user.group.group_id		ID группы
 * @apiSuccess (Success 200) {String} 	user.group.group_name	Имя группы
 *
 * @apiSuccessExample {json} Success-Response:
 *[
 *   {
 *       "user_id": 1,
 *       "user_name": "Евгений Николаевич Табаков",
 *       "email": "tabakov.e.n@limefin.com",
 *       "group": {
 *           "group_id": 2,
 *           "group_name": "admin"
 *       }
 *   },
 *   {
 *       "user_id": 2,
 *       "user_name": "Вячеслав Викторович Тищенко",
 *       "email": "tishchenko.v.v@limefin.com",
 *       "group": {
 *           "group_id": 2,
 *           "group_name": "admin"
 *       }
 *   }
 *]
 *
 *
 */
