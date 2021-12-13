package scheduler

/**
 *
 * @api {POST} /helpdesk/table/offices_list 02. Обновление списка оффисов
 * @apiName UpdateOfficesList
 * @apiGroup 8. Табель
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam  {[]Office} 	office 							Массив объектов "Офис"
 * @apiParam  {Uint64} 		office.id 					Ид смены, для новой смены 0 или без этого поля
 * @apiParam  {String} 		office.name 				Название смены
 * @apiParam  {String} 		office.start_time 	Время начала работы
 * @apiParam  {String} 		office.end_time 		Время конца работы
 * @apiParam  {String} 		office.color 				Цвет которым офис будет отображаться на графике
 * @apiParam  {Bool} 			office.deleted 			Признак удален офис или нет.
 *
 * @apiSuccess (200) {String} status Статус ответа
 *
 * @apiParamExample  {JSON} Request-Example:
 * [
 *   {
 *     "id": 1,
 *     "name": "первая",
 *     "start_time": "8:00",
 *     "end_time": "20:00",
 *     "color": "0xffffff",
 *     "deleted":false
 *   },
 *   {
 *     "id":2,
 *     "name": "вторая",
 *     "start_time": "9:00",
 *     "end_time": "21:00",
 *     "color": "0x3b6a32",
 *     "deleted":true
 *   },
 *     {
 *     "name": "артилерийская",
 *     "start_time": "7:30",
 *     "end_time": "19:30",
 *     "color": "0x5b90f3",
 *     "deleted":false
 *   }
 * ]
 *
 *
 * @apiSuccessExample {JSON} Success-Response:
 * {
 *     "status": "ok"
 * }
 *
 *
 */
