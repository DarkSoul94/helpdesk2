package docs

/**
 *
 * @api {POST} /helpdesk/reports/tickets_grades 04. Получение списка оценок запросов
 * @apiName GetReturnedTickets
 * @apiGroup 06. Отчёты
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 * @apiDescription Принимает в себя список ид пользователей и/или список отделов пользователей по которым нужно получить список оцененных запросов и их оценки
 *
 * @apiHeader {String} BearerToken Авторизационный токен
 *
 * @apiParam  {String} 		start_date 		Дата начала периода выборки
 * @apiParam  {String} 		end_date 		Дата конца периода выборки
 * @apiParam  {Uint64[]} 	users_id 		Массив ид пользователей по которым нужно сделать выборку
 * @apiParam  {String[]} 	departments 	Массив отделов по которым нужно сделать выборку
 *
 * @apiSuccess (200) {DepartmentsTicketsGrades[]} 	departments_tickets_grades 												Массив обектов "оценки запросов по пользователям по отделу"
 * @apiSuccess (200) {String} 						departments_tickets_grades.department 									Название отдела
 * @apiSuccess (200) {UsersTicketsGrades[]} 		departments_tickets_grades.users_grades 								Массив обектов "оценки запросов по пользователю"
 * @apiSuccess (200) {String} 						departments_tickets_grades.users_grades.user_name 						ФИО пользователя
 * @apiSuccess (200) {TicketsGrades[]} 				departments_tickets_grades.users_grades.tickets_grades 					Массив обектов "оценка запроса"
 * @apiSuccess (200) {Uint64} 						departments_tickets_grades.users_grades.tickets_grades.ticket_id 		ИД запроса
 * @apiSuccess (200) {Uint64} 						departments_tickets_grades.users_grades.tickets_grades.ticket_grade 	Оценка запроса
 * @apiSuccess (200) {Float64} 						departments_tickets_grades.users_grades.average_user_grade 				Средняя оценка по пользователю
 * @apiSuccess (200) {Float64} 						departments_tickets_grades.avarege_department_grade 					Средняя оценка по отделу
 *
 * @apiParamExample  {type} Request-Example:
 * {
 *     "start_date" : "2021-06-01",
 *     "end_date" : "2021-06-20",
 *     "users_id":
 *             [
 *                 4,
 *                 5,
 *                 6
 *             ],
 *		"departments": []
 * }
 *
 *
 * @apiSuccessExample {type} Success-Response:
 * [
 *   {
 *     "department": "Техническая поддержка",
 *     "users_grades": [
 *       {
 *         "user_name": "Вячеслав Викторович Тищенко",
 *         "tickets_grades": [
 *           {
 *             "ticket_id": 1,
 *             "ticket_grade": 5
 *           },
 *           {
 *             "ticket_id": 3,
 *             "ticket_grade": 5
 *           },
 *           {
 *             "ticket_id": 4,
 *             "ticket_grade": 2
 *           },
 *           {
 *             "ticket_id": 31,
 *             "ticket_grade": 3
 *           }
 *         ],
 *         "average_user_grade": 3.75
 *       },
 *       {
 *         "user_name": "Евгений Николаевич Табаков",
 *         "tickets_grades": [
 *           {
 *             "ticket_id": 19,
 *             "ticket_grade": 4
 *           },
 *           {
 *             "ticket_id": 22,
 *             "ticket_grade": 5
 *           }
 *         ],
 *         "average_user_grade": 4.5
 *       }
 *     ],
 *     "avarege_department_grade": 4.13
 *   },
 *   {
 *     "department": "Разработчики",
 *     "users_grades": [
 *       {
 *         "user_name": "Артем Владимирович Шелкопляс",
 *         "tickets_grades": [
 *           {
 *             "ticket_id": 2,
 *             "ticket_grade": 3
 *           },
 *           {
 *             "ticket_id": 5,
 *             "ticket_grade": 5
 *           },
 *           {
 *             "ticket_id": 15,
 *             "ticket_grade": 5
 *           }
 *         ],
 *         "average_user_grade": 4.33
 *       }
 *     ],
 *     "avarege_department_grade": 4.33
 *   }
 * ]
 *
 *
 */
