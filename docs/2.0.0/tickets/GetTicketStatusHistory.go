package docs

/**
 *
 * @api {get} /helpdesk/ticket_status/history 09. Получение истории изменения статусов запроса
 * @apiName GetTicketStatusHistory
 * @apiGroup 03. Запросы в тех. поддержку
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam  {Uint64} ticket_id ID запроса по которому нужно посмотреть историю статусов
 *
 * @apiSuccess (Success 200) {StatusHistory[]} 	status_history 						Массив объектов "история изменения статуса"
 * @apiSuccess (Success 200) {String} 			status_history.curr_status_time 	Время присвоения текущего статуса
 * @apiSuccess (Success 200) {String} 			status_history.curr_status        	Название текущего статуса
 * @apiSuccess (Success 200) {String} 			status_history.changed_user       	Пользователь сменивший статус
 * @apiSuccess (Success 200) {Uint64} 			status_history.difference       	Время нахождения в статусе
 *
 * @apiExample  Example usage:
 * http://localhost:8888/helpdesk/ticket_status/history?ticket_id=2
 *
 *
 * @apiSuccessExample {json} Success-Response:
 * [
 *   {
 *     "curr_status_time": "2021-11-24T12:43:53Z",
 *     "curr_status": "В ожидании",
 *     "changed_user": "Вячеслав Викторович Тищенко",
 *     "difference": 1784
 *   }
 * ]
 *
 * @apiError ErrStatusHistoryNotExist В базе нет записей истории изменения статусов по данному запросу
 *
 */
