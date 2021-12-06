package docs

/**
*
* @api {POST} /helpdesk/ticket/filtered_tickets_list 05. Получение отфильтрованого списка запросов
* @apiName apiName
* @apiGroup 03. Запросы в тех. поддержку
* @apiVersion  2.0.0
* @apiSampleRequest off
* @apiDescription Если у обратившегося пользователя стоит доступ `full_search` фильтр накладывается на все запросы.
* Если стоит `can_resolve` накладывается на список запросов которые требовали и требуют согласования и там где пользователь автор запроса.
* Для всех остальных накладывается на список где пользователь автор запроса.
* Поле support_id игнорируется для всех у кого не стоит доступ `full_search`. Если в фильтре ничего не указано отправится список всех запросов с учётом доступов пользователя.
*
* @apiHeader {String} BearerToken 	Авторизационный токен
*
* @apiParam  {Uint64} 	ticket_id 		ИД запроса
* @apiParam  {String} 	start_date 		Дата начала выборки
* @apiParam  {String} 	end_date 		Дата конца выборки
* @apiParam  {Uint64} 	category_id 	ИД категории
* @apiParam  {Uint64[]} section_id 		Массив ид раздела категории
* @apiParam  {String} 	text 			Текст который должен содержаться в запросе
* @apiParam  {Uint64} 	status_id 		ИД статуса запроса
* @apiParam  {Uint64[]} author_id 		Массив ид автора запроса
* @apiParam  {Uint64[]} support_id 		Массив ид сотрудника работавшего над запросом
* @apiParam  {String} 	filial 			Филиал запроса
* @apiParam  {String} 	comment 		Текст который должен содержаться в комментариях к этому запросу
*
* @apiSuccess (Success 200) {String[]}      fields                              Список полей которые должны отрисоваться на стороне фронта
* @apiSuccess (Success 200) {Ticket[]}      tickets                             Массив запросов в тех. поддержку
* @apiSuccess (Success 200) {Uint64} 	    tickets.ticket_id 					ID запроса
* @apiSuccess (Success 200) {String} 	    tickets.ticket_date 				Дата создания запроса
* @apiSuccess (Success 200) {String}        tickets.category   					Категория
* @apiSuccess (Success 200) {String}  		tickets.section            			Раздел категории
* @apiSuccess (Success 200) {String} 		tickets.ticket_text 				Текст запроса
* @apiSuccess (Success 200) {String}     	tickets.status 						Статус запроса
* @apiSuccess (Success 200) {String}     	tickets.ticket_author               Автор запроса
* @apiSuccess (Success 200) {String}     	tickets.support               		Сотрудник ТП работающий над запросом
* @apiSuccess (Success 200) {String}     	tickets.filial               		Отделение за которым закрелен автор запроса
* @apiSuccess (Success 200) {Uint}     		tickets.grade               		Оценка запроса
*
* @apiParamExample  {json} Request-Example:
* {
*     "ticket_id": 131,
*     "start_date" : "2021-06-17 10:00:00",
*     "end_date" : "2021-06-17 10:30:00",
*     "category_id" : 1,
*     "section_id" : [1, 2, 3],
*     "text" : "asdfasf",
*     "status_id" : 1,
*     "author_id" : [1, 2, 3],
*     "support_id" : [1, 2, 3],
*     "filial" : "asfasdcv"
* }
*
*
* @apiSuccessExample {json} Ответ для обычного пользователя и сотрудника бэк-офиса:
* {
*   "fields": [
*     "ticket_id",
*     "ticket_date",
*     "category",
*     "section",
*     "ticket_text",
*     "status",
*     "filial",
*     "ticket_author",
*     "grade"
*   ],
*   "tickets": [
*     {
*       "ticket_id": 141,
*       "ticket_date": "2021-06-17T14:42:30Z",
*       "significant": false,
*       "category": "Валютообмен",
*       "section": "Спецоперация",
*       "ticket_text": "asfsdfgsgvxcfasf",
*       "status": "В процессе реализации",
*       "ticket_status_id": 5,
*       "filial": "not found",
*       "ticket_author": "Вячеслав Викторович Тищенко",
*       "support": "",
*       "grade": 0
*     }
*  ]
* }
*
* @apiSuccessExample {json} Ответ для тех у кого есть разрешение `full_search`:
* {
*   "fields": [
*    "ticket_id",
*    "ticket_date",
*    "category",
*    "section",
*    "ticket_text",
*    "status",
*    "filial",
*    "ticket_author",
*    "support",
*    "grade"
*   ],
*   "tickets": [
*     {
*       "ticket_id": 140,
*       "ticket_date": "2021-06-17T13:20:57Z",
*       "significant": true,
*       "category": "1С ",
*       "section": "Изменение/удаление кассовых ордеров",
*       "ticket_text": "выпывп",
*       "status": "В процессе реализации",
*       "ticket_status_id": 5,
*       "filial": "not found",
*       "ticket_author": "Евгений Николаевич Табаков",
*       "support": "Евгений Николаевич Табаков",
*       "grade": 0
*     }
*  ]
* }
*
 */
