package docs

/**
 *
 * @api {GET} /helpdesk/reports/tickets_count 06. Количество поступивших запросов в разрезе часов и дней
 * @apiName GetTicketsCountByDaysHours
 * @apiGroup 06. Отчёты
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken Авторизационный токен
 *
 * @apiParam  {String} start_date 	Дата начала выборки, включительно
 * @apiParam  {String} end_date		Дата конца выборки, данные за этот день не учитываются
 *
 * @apiExample  Example usage:
 * http://localhost:8888/helpdesk/reports/tickets_count?start_date=2021-04-01&end_date=2021-05-01
 *
 * @apiSuccess (Success 200) {CountByDay[]} 	count_by_day 						Массив обектов "Количество запросов за часы дня"
 * @apiSuccess (Success 200) {String} 			count_by_day.day 					Дата
 * @apiSuccess (Success 200) {CountByHour[]} 	count_by_day.count_by_hour 			Массив обектов "Количество запросов за час"
 * @apiSuccess (Success 200) {String} 			count_by_day.count_by_hour.hour 	Временной диапазон
 * @apiSuccess (Success 200) {Uint} 			count_by_day.count_by_hour.count 	Количество запросов
 *
 * @apiSuccessExample {json} Success-Response:
 * [
 *   {
 *     "date": "2021-06-01",
 *     "count_by_hour": [
 *       {
 *         "hour": "19:00:00 - 19:59:59",
 *         "count": 2
 *       },
 *       {
 *         "hour": "20:00:00 - 20:59:59",
 *         "count": 6
 *       }
 *     ]
 *   },
 *   {
 *     "date": "2021-06-02",
 *     "count_by_hour": [
 *       {
 *         "hour": "14:00:00 - 14:59:59",
 *         "count": 1
 *       },
 *       {
 *         "hour": "19:00:00 - 19:59:59",
 *         "count": 1
 *       }
 *     ]
 *   },
 * ]
 *
 *
 */
