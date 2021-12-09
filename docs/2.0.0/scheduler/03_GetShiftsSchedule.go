package scheduler

/**
 *
 * @api {GET} /helpdesk/table/schedule 03. Получение графика смен
 * @apiName GetShiftsSchedule
 * @apiGroup 8. Табель
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiExample  Example usage:
 * http://localhost:8888/helpdesk/table/schedule?date=2021-10-01
 *
 * @apiSuccess (200) {[]Office} 							legend 														Список оффисов для легенды графика
 * 																																							(список офисов который включает в себя актуальные офисы,
 *																																							а также те удаленные которые уже были отмечены в графике
 *																																							за указанный период)
 * @apiSuccess (200) {Uint64} 								legend.id 												ИД смены
 * @apiSuccess (200) {String} 								legend.name 											Название смены
 * @apiSuccess (200) {String} 								legend.color 											Цвет для отображения на графике
 * @apiSuccess (200) {Bool} 									legend.deleted 										Признак являеться ли смена удаленной
 * @apiSuccess (200) {[]ShiftsScheduleCell} 	shifts_schedule 									Массив объектов "ячейка графика смен"
 * @apiSuccess (200) {Uint64} 								shifts_schedule.id 								ИД ячейки
 * @apiSuccess (200) {Uint64} 								shifts_schedule.support_id 				ИД сотрудника тп
 * @apiSuccess (200) {Uint64} 								shifts_schedule.shift_id 					ИД смены
 * @apiSuccess (200) {String} 								shifts_schedule.date 							Дата смены
 * @apiSuccess (200) {String} 								shifts_schedule.start_time 				Время начала смены
 * @apiSuccess (200) {String} 								shifts_schedule.end_time 					Время конца смены
 * @apiSuccess (200) {Bool} 									shifts_schedule.sick_leave				Признак больничного
 * @apiSuccess (200) {Bool} 									shifts_schedule.vacation 					Признак отпуска
 * @apiSuccess (200) {Bool} 									shifts_schedule.late 							Признак было ли опоздание в эту смену
 * @apiSuccess (200) {[]Support} 							regular_supports									Список обычных суппортов
 * @apiSuccess (200) {Uint64} 								regular_supports.id								ID суппорта
 * @apiSuccess (200) {String} 								regular_supports.name							Имя суппорта
 * @apiSuccess (200) {String} 								regular_supports.color						Цвет отображения
 * @apiSuccess (200) {Uint64} 								regular_supports.senior_id				ID старшего суппорта (если не назначен прийдет `0`)
 * @apiSuccess (200) {[]Support} 							senior_supports										Список старших суппортов
 * @apiSuccess (200) {Uint64} 								senior_supports.id								ID суппорта
 * @apiSuccess (200) {String} 								senior_supports.name							Имя суппорта
 * @apiSuccess (200) {String} 								senior_supports.color						  Цвет отображения
 * @apiSuccess (200) {Uint64} 								senior_supports.senior_id					ID старшего суппорта (если не назначен прийдет `0`)
 *
 *
 * @apiSuccessExample {Json} Success-Response:
 * {
 *   "legend": [
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
 *     },
 *     {
 *       "id": 11,
 *       "name": "Новая",
 *       "color": "#D92B48",
 *       "deleted": false
 *     },
 *     {
 *       "id": 12,
 *       "name": "*новая",
 *       "color": "#677FD9FF",
 *       "deleted": false
 *     }
 *   ],
 *   "shifts_schedule": [
 *     {
 *       "id": 62,
 *       "support_id": 4,
 *       "office_id": 9,
 *       "start_time": "08:00",
 *       "end_time": "20:00",
 *       "date": "2021-12-08",
 *       "vacation": false,
 *       "sick_leave": true,
 *       "late": false
 *     },
 *     {
 *       "id": 63,
 *       "support_id": 4,
 *       "office_id": 9,
 *       "start_time": "08:00",
 *       "end_time": "20:00",
 *       "date": "2021-12-07",
 *       "vacation": false,
 *       "sick_leave": false,
 *       "late": false
 *     }
 *   ],
 *   "status": "ok",
 *	 "regular_supports": [
 *	   {
 *	     "id": 4,
 *	     "name": "Вячеслав Викторович Тищенко",
 *	     "color": "0xFFFFFF",
 *	     "senior_id": 5
 *	   },
 *	   {
 *	     "id": 6,
 *	     "name": "Евгений Николаевич Табаков",
 *	     "color": "0xFFFFFF",
 *	     "senior_id": 5
 *	   },
 *	   {
 *	     "id": 7,
 *	     "name": "Владислав Сергеевич Маспанов",
 *	     "color": "0xFFFFFF",
 *	     "senior_id": 0
 *	   }
 *	 ],
 *   "senior_supports": [
 *    {
 *      "id": 5,
 *      "name": "Артем Владимирович Шелкопляс",
 *      "color": "0xFFFFFF",
 *      "senior_id": 0
 *    }
 *  ],
 * }
 *
 *
 */
