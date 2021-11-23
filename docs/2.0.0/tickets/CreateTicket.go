package docs

/**
*
* @api {post} /helpdesk/ticket/create Создание запроса в тех. поддержку
* @apiName CreateTicket
* @apiGroup 03. Запросы в тех. поддержку
* @apiVersion  2.0.0
* @apiSampleRequest off
*
*
* @apiHeader {String} BearerToken 	Авторизационный токен
*
* @apiParam {Uint64} 	        section_id 				ID раздела категории.
* @apiParam {String} 	        ticket_text 			Текст запроса.
* @apiParam {file[]}            files                   Файлы.
* @apiParam {String}	        files.file_name 		Имя файла.
* @apiParam {String}	        files.file_data 		Данные файла.
*
* @apiSuccess (Success 200) {String} status 	Статус ответа на запрос
* @apiSuccess (Success 200) {Uint64} ticket_id 	ID созданого запроса
*
*
* @apiParamExample  {json} Request-Example:
* {
*   "section_id": 1,
*   "ticket_text": "Не вышел чек",
*   "files": [
*       {
*           "file_name": "скрин1.новый_клиент.jpg",
*           "file_data": "a few byte count"
*       },
*       {
*           "file_name": "скрин2.jpg",
*           "file_data": "a few byte count"
*       }
*   ]
* }
*
*
* @apiSuccessExample {json} Success-Response:
* {
*     "status": "ok",
*     "ticket_id": 8
* }
*
* @apiError ErrInvalidID Неверный ID категории
* @apiError ErrBlankTicketText Пустой текст запроса
* @apiError ErrBlankCategoryName Пустое имя категории
* @apiError ErrBlankSectionName Пустое имя раздела категории
*
*
 */
