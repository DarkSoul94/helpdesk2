package docs

/**
 *
 * @api {GET} /helpdesk/reports/average_grades 03. Получение списка средних оценок за запросы
 * @apiName GetAverageGrades
 * @apiGroup 06. Отчёты
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 * @apiDescription Получение списка усредненных оценок запросов по каждому сотруднику тех. поддержки. Последняя запись - усредненная оценка по отделу.
 *
 * @apiHeader {String} BearerToken Авторизационный токен
 *
 * @apiParam  {String} start_date 	Дата начала выборки, включительно
 * @apiParam  {String} end_date		Дата конца выборки, данные за этот день не учитываются
 *
 * @apiExample  Example usage:
 * http://localhost:8888/helpdesk/reports/average_grades?start_date=2021-05-01&end_date=2021-06-01
 *
 * @apiSuccess (Success 200) {AverageGrade[]}  	average_grade                            	Массив объектов `средняя оценка за запросы по сотруднику`
 * @apiSuccess (Success 200) {String}  			average_grade.support                    	Сотрудник тех. поддержки
 * @apiSuccess (Success 200) {Uint64} 	        average_grade.average_grade_by_support 		Средняя оценка запросов
 *
 * @apiSuccessExample {json} Success-Response:
 * [
 *     {
 *         "support": "Вячеслав Викторович Тищенко",
 *         "average_grade_by_support": 5
 *     },
 *     {
 *         "support": "Отдел ТП",
 *         "average_grade_by_support": 5
 *     }
 * ]
 *
 *
 */
