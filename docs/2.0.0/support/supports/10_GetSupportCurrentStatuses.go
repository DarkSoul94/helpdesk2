package docs

/**
 *
 * @api {GET} /support/get_current_statuses 10. Получение списка сотрудников ТП с их текущим рабочим статусом
 * @apiName GetSupportCurrentStatuses
 * @apiGroup 05. Сотрудник ТП
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {String}						status 									Статус выполнения запроса
 * @apiSuccess (Success 200) {support_current_status[]} 	support_current_status 					Массив статистика по запросам и текущих статусов сотрудников ТП работающих сегодня
 * @apiSuccess (Success 200) {String} 						support_current_status.support_id 		ID сотрудника ТП
 * @apiSuccess (Success 200) {String} 						support_current_status.support 			Имя сотрудника ТП
 * @apiSuccess (Success 200) {String} 						support_current_status.status 	    	Текущий статус сотрудника ТП
 * @apiSuccess (Success 200) {Bool} 						support_current_status.shift_status 	Текущий статус смены сотрудника ТП
 * @apiSuccess (Success 200) {Uint} 						support_current_status.in_work 			Количество запросов в работе у указанного сотрудника
 * @apiSuccess (Success 200) {Uint} 						support_current_status.postproned 		Количество отложеных запросов у указанного сотрудника
 * @apiSuccess (Success 200) {Uint} 						support_current_status.complete 		Количество выполненных за сегодня запросов у указанного сотрудника
 * @apiSuccess (Success 200) {Uint} 						support_current_status.revision 		Количество запросов за сегодня отправленных на доработку у указанного сотрудника
 * @apiSuccess (Success 200) {Bool} 						support_current_status.priority 		Признак приоритета по распределению запросов на саппорта
 * @apiSuccess (Success 200) {total} 						total							 		Суммарная статистика по кол-ву запросов в работе, отложеных, выполненых по сотрудникам за сегодня
 * @apiSuccess (Success 200) {Uint} 						total.total_in_work						Общее кол-во запросов в работе
 * @apiSuccess (Success 200) {Uint} 						total.total_postproned					Общее кол-во отложенных запросов
 * @apiSuccess (Success 200) {Uint} 						total.total_complete					Общее кол-во выполненных за сегодня запросов
 * @apiSuccess (Success 200) {Uint} 						total.total_revision					Общее кол-во запросов отправленных на доработку за сегодня
 * @apiSuccess (Success 200) {Uint} 						wait_ticket_count				    	Кол-во запросов в очереди на распределение
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "status": "ok",
 *   "support_current_status": [
 *     {
 *       "support_id": 5,
 *       "support": "Артем Владимирович Шелкопляс",
 *       "status": "Принимаю запросы",
 *       "shift_status": true,
 *       "in_work": 1,
 *       "postproned": 0,
 *       "complete": 0
 *     },
 *     {
 *       "support_id": 6,
 *       "support": "Вячеслав Викторович Тищенко",
 *       "status": "Не принимаю запросы",
 *       "shift_status": false,
 *       "in_work": 0,
 *       "postproned": 0,
 *       "complete": 0
 *     }
 *   ],
 *   "total": {
 *     "total_in_work": 1,
 *     "total_postproned": 0,
 *     "total_complete": 0
 *   },
 *   "wait_ticket_count": 0
 * }
 *
 *
 */
