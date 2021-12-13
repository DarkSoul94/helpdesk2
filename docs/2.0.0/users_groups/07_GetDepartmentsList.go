package users

/**
 *
 * @api {GET} /helpdesk/user/departments_list 07. Получение списка отделов сотрудников
 * @apiName GetDepartmentsList
 * @apiGroup 04. Пользователи
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken Авторизационный токен
 *
 * @apiSuccess (200) {Departmen[]} 	departments Список отделов
 * @apiSuccess (200) {String} 		status 		Статус запроса
 *
 * @apiSuccessExample {type} Success-Response:
 * {
 *   "departments": [
 *     "Техническая поддержка",
 *     "Разработчики"
 *   ],
 *   "status": "ok"
 * }
 *
 *
 */
