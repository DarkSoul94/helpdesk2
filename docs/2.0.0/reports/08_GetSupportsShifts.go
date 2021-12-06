package docs

/**
 *
 * @api {GET} /helpdesk/reports/supports_shifts 08. Время открытия и закрытия смен супортов за период
 * @apiName GetSupportsShiftsOpeningTime
 * @apiGroup 06. Отчёты
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken Авторизационный токен
 *
 * @apiParam  {String} start_date 	Дата начала выборки
 * @apiParam  {String} end_date		Дата конца выборки
 *
 * @apiExample  Example usage:
 * http://localhost:8888/helpdesk/reports/supports_shifts?start_date=2021-08-04&end_date=2021-11-10
 *
 * @apiSuccess (200) {[]Period} 		period 												Результат за указанный период разделенный по месяцам
 * @apiSuccess (200) {[]SupportShifts} 	period.support_shifts 								Массив объектов "смены суппорта"
 * @apiSuccess (200) {String} 			period.support_shifts.support 						ФИО суппорта
 * @apiSuccess (200) {String} 			period.support_shifts.with_out_grace_time 			Время опоздания свыше льготного периода
 * @apiSuccess (200) {[]Shift} 			period.support_shifts.shifts 						Массив объектов "смена"
 * @apiSuccess (200) {String} 			period.support_shifts.shifts.opening_date 			Время открытия смены
 * @apiSuccess (200) {String} 			period.support_shifts.shifts.closing_date 			Время закрытия смены
 * @apiSuccess (200) {Uint64} 			period.support_shifts.shifts.count_of_minutes_late 	Опоздание за эту смены в минутах
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "2021-08-04 ~ 2021-08-31": [],
 *   "2021-09-01 ~ 2021-09-30": [],
 *   "2021-10-01 ~ 2021-10-31": [
 *     {
 *       "support": "Артем Владимирович Шелкопляс",
 *       "with_out_grace_time": "0s",
 *       "shifts": [
 *         {
 *           "opening_date": "2021-10-08 11:59:00",
 *           "closing_date": "2021-10-08 23:36:00",
 *           "count_of_minutes_late": 0
 *         }
 *       ]
 *     },
 *     {
 *       "support": "Вячеслав Викторович Тищенко",
 *       "with_out_grace_time": "9h52m0s",
 *       "shifts": [
 *         {
 *           "opening_date": "2021-10-08 17:35:00",
 *           "closing_date": "2021-10-08 22:35:00",
 *           "count_of_minutes_late": 605
 *         },
 *         {
 *           "opening_date": "2021-10-09 10:29:00",
 *           "closing_date": "2021-10-09 22:31:00",
 *           "count_of_minutes_late": 0
 *         }
 *       ]
 *     }
 *   ],
 *   "2021-11-01 ~ 2021-11-10": []
 * }
 *
 *
 */
