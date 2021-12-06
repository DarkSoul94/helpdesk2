package docs

/**
*
* @api {GET} /helpdesk/reports/tickets_status_difference 02. Время нахождения запроса в разных статусах
* @apiName GetTicketStatusDifference
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
* http://localhost:8888/helpdesk/reports/tickets_status_difference?start_date=2021-05-01&end_date=2021-06-01
*
* @apiSuccess (Success 200) {TicketStatuDifferenceTime[]} 	ticket_status_difference_time									Массив обектов "время нахождения запроса в разных статусах"
* @apiSuccess (Success 200) {Uint} 							ticket_status_difference_time.ticket_id 						Ид запроса
* @apiSuccess (Success 200) {String} 						ticket_status_difference_time.support_name 						Имя сотрудника тех.-поддержки
* @apiSuccess (Success 200) {String} 						ticket_status_difference_time.section 							Имя раздела категории
* @apiSuccess (Success 200) {StatusDifference[]} 			ticket_status_difference_time.status_difference 				Масиив обектов "время нахождения в статусе"
* @apiSuccess (Success 200) {String} 						ticket_status_difference_time.status_difference.status 			Имя статуса
* @apiSuccess (Success 200) {String} 						ticket_status_difference_time.status_difference.diff_time 		Время нахождения в этом статусе
*
* @apiSuccessExample {json} Success-Response:
* [
*     {
*         "ticket_id": 3,
*         "support_name": "Вячеслав Викторович Тищенко",
*         "section": "Прочее",
*         "status_difference": [
*             {
*                 "status": "В ожидании",
*                 "diff_time": "12m35s"
*             },
*             {
*                 "status": "Отправлен на доработку",
*                 "diff_time": "1m26s"
*             }
*         ]
*     },
*     {
*         "ticket_id": 4,
*         "support_name": "Вячеслав Викторович Тищенко",
*         "section": "Прочее",
*         "status_difference": [
*             {
*                 "status": "В ожидании",
*                 "diff_time": "5s"
*             },
*             {
*                 "status": "В работе",
*                 "diff_time": "2m42s"
*             },
*             {
*                 "status": "В процессе реализации",
*                 "diff_time": "17s"
*             },
*             {
*                 "status": "Отправлен на доработку",
*                 "diff_time": "18s"
*             },
*             {
*                 "status": "Отложен",
*                 "diff_time": "30s"
*             }
*         ]
*     }
* ]
*
*
 */
