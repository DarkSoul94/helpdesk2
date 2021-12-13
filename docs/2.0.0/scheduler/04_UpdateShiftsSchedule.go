package scheduler

/**
 *
 * @api {POST} /helpdesk/table/update_schedule 04. Обновить график смен
 * @apiName UpdateShiftsSchedule
 * @apiGroup 8. Табель
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam  {[]ShiftsScheduleCell} 	shiftsSchedule 								Массив объектов "ячейка графика смен"
 * @apiParam  {Uint64} 								shiftsSchedule.id 						ИД ячейки. При добалении новой ячейки ид = 0
 * @apiParam  {Uint64} 								shiftsSchedule.support_id 		ИД сотрудника ТП
 * @apiParam  {Uint64} 								shiftsSchedule.office_id 			ИД смены
 * @apiParam  {String} 								shiftsSchedule.date 					Дата в графике
 * @apiParam  {String} 								shiftsSchedule.start_time 		Время начала смены
 * @apiParam  {String} 								shiftsSchedule.end_time 			Время конца смены
 * @apiParam  {Bool} 									shiftsSchedule.vacation 			Признак отпуска
 * @apiParam  {Bool} 									shiftsSchedule.sick_leave 		Признак больничного
 *
 * @apiSuccess (Success 200) {String} status Статус ответа
 *
 *
 * @apiParamExample  {Json} Обновить смену в графике:
 * [
 *   {
 *     "id": 8,
 *     "support_id": 4,
 *     "office_id": 10,
 *     "date": "2021-10-15"
 *     "start_time": "08:00",
 *     "end_time": "20:00",
 *		 "vacation": false,
 *     "sick_leave": false
 *   }
 * ]
 *
 *
 *
 * @apiSuccessExample {Json} Success-Response:
 * {
 *     "status": "ok"
 * }
 *
 *
 */
