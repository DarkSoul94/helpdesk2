package docs

/**
*
* @api {GET} /helpdesk/reports/motivation 01. Отображние мотивации сотрудников ТП
* @apiName GetMotivation
* @apiGroup 06. Отчёты
* @apiVersion  2.0.0
* @apiSampleRequest off
* @apiDescription Получение мотивации сотрудников тп в виде таблицы, где данные рассортированы по сотрудникам тп и категориям запросов.
* Каждый объект отображает название категории, количество выполненых запросов, и мотивацию за их выполнение. Последний обект "Total" отображает сумарное количество выполненых сотрудником запросов и его полную мотивацию.
*
* @apiHeader {String} BearerToken Авторизационный токен
*
* @apiParam  {String} start_date 	Дата начала выборки, включительно
* @apiParam  {String} end_date		Дата конца выборки, данные за этот день не учитываются
*
* @apiExample  Example usage:
* http://localhost:8888/helpdesk/reports/motivation?start_date=2021-03-01&end_date=2021-05-01
*
* @apiSuccess (Success 200) {[]Period} 			period 											Результат за указанный период разделенный по месяцам
* @apiSuccess (Success 200) {[]Motivation} 		period.motivation								Массив мотиваций по сотрудникам тех. поддержки
* @apiSuccess (Success 200) {Support} 			period.motivation.support						Объект с информацией по сотруднику тех. поддержки
* @apiSuccess (Success 200) {Uint64} 			period.motivation.support.id					ID сотрудника тех. поддержки
* @apiSuccess (Success 200) {String} 			period.motivation.support.name					Имя сотрудника тех. поддержки
* @apiSuccess (Success 200) {String} 			period.motivation.support.color					Цвет для отображения сотрудника тех. поддержки в графике
* @apiSuccess (Success 200) {[]Categories} 		period.motivation.categories					Массив объектов с информацией по количеству запросов выполненных сотрудником в разрезе категорий запросов
* @apiSuccess (Success 200) {Uint64} 			period.motivation.categories.id					ID категории запроса
* @apiSuccess (Success 200) {String} 			period.motivation.categories.name				Название категории
* @apiSuccess (Success 200) {Uint64} 			period.motivation.categories.tickets_count		Количество запросов по категории
* @apiSuccess (Success 200) {Uint64} 			period.motivation.total_tickets_count			Общее количество запросов по категории
* @apiSuccess (Success 200) {Float64} 			period.motivation.total_motivation				Общая мотивация по сотруднику тех. поддержки
* @apiSuccess (Success 200) {Float64} 			period.motivation.total_by_shifts				Общая сумма оплаты по открытым сменам сотрудника за указанный период.
* @apiSuccess (Success 200) {Float64} 			period.motivation.total_payment					Общая сумма оплаты с учетом мотивации и оплаты смен.
*
* @apiSuccessExample {json} Success-Response:
* {
*   "2021-10-01 ~ 2021-10-30": [
*     {
*       "support": {
*         "id": 4,
*         "name": "Вячеслав Викторович Тищенко",
*         "color": ""
*       },
*       "categories": [
*         {
*           "id": 1,
*           "name": "Сервисная категория",
*           "tickets_count": 3
*         }
*       ],
*       "total_tickets_count": 0,
*       "total_motivation": 3,
*       "total_by_shifts": 1500,
*       "total_payment": 1503
*     },
*     {
*       "support": {
*         "id": 5,
*         "name": "Артем Владимирович Шелкопляс",
*         "color": "0xFFFFFF"
*       },
*       "categories": [],
*       "total_tickets_count": 0,
*       "total_motivation": 0,
*       "total_by_shifts": 0,
*       "total_payment": 0
*     },
*     {
*       "support": {
*         "id": 6,
*         "name": "Евгений Николаевич Табаков",
*         "color": ""
*       },
*       "categories": [],
*       "total_tickets_count": 0,
*       "total_motivation": 0,
*       "total_by_shifts": 0,
*       "total_payment": 0
*     },
*     {
*       "support": {
*         "id": 0,
*         "name": "Итого",
*         "color": ""
*       },
*       "categories": [
*         {
*           "id": 1,
*           "name": "Сервисная категория",
*           "tickets_count": 3
*         }
*       ],
*       "total_tickets_count": 0,
*       "total_motivation": 3,
*       "total_by_shifts": 1500,
*       "total_payment": 1503
*     }
*   ]
* }
 *
*/
