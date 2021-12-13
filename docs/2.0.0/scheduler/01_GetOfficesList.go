package scheduler

/**
 *
 * @api {GET} /helpdesk/table/offices_list 01. Получить список офисов
 * @apiName GetOfficesList
 * @apiGroup 8. Табель
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (200) {[]Office} 	actual 							Список актуальных офисов
 * @apiSuccess (200) {Uint64} 		actual.id 					ИД смены
 * @apiSuccess (200) {String} 		actual.name 				Название смены
 * @apiSuccess (200) {String} 		actual.color 				Цвет отображения в графике
 * @apiSuccess (200) {Bool} 			actual.deleted 			Признак удалена смена или нет.
 * @apiSuccess (200) {[]Office} 	deleted 						Список актуальных офисов
 * @apiSuccess (200) {Uint64} 		deleted.id 					ИД смены
 * @apiSuccess (200) {String} 		deleted.name 				Название смены
 * @apiSuccess (200) {String} 		deleted.color 			Цвет отображения в графике
 * @apiSuccess (200) {Bool} 			deleted.deleted 		Признак удалена смена или нет.
 *
 * @apiSuccessExample {Json} Success-Response:
 * {
 *   "actual": [
 *     {
 *       "id": 9,
 *       "name": "вторая",
 *       "color": "#BEFF2E",
 *       "deleted": false
 *     },
 *     {
 *       "id": 10,
 *       "name": "артилерийская",
 *       "color": "#1FC91EF2",
 *       "deleted": false
 *     }
 *   ],
 *   "deleted": [
 *     {
 *       "id": 8,
 *       "name": "первая",
 *       "color": "#487C7CFF",
 *       "deleted": true
 *     }
 *   ],
 *   "status": "ok"
 * }
 *
 *
 */
