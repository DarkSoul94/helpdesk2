package docs

/**
*
* @api {post} /helpdesk/ticket/update Обновление запроса в ТП
* @apiName UpdateTicket
* @apiGroup 03. Запросы в тех. поддержку
* @apiVersion  2.0.0
* @apiSampleRequest off
* @apiDescription Обновление запроса в тех.поддержку.
*
* @apiHeader {String} BearerToken 	Авторизационный токен
*
* @apiParam {Uint64} 	        ticket_id 						ID запроса
* @apiParam {Uint64} 	        section_id 						ID раздела категории. Передается только при изменении раздела категории в запросе
* @apiParam {Uint64} 	        ticket_status_id 		        ID статуса запроса. Передается только при изменении статуса запроса
* @apiParam {Uint64} 	        support_id 			        	ID сотрудника ТП, передается только при изменении сотрудника ТП в запросе
* @apiParam {String} 	        service_comment 				Сервисный комментарий для сотрудников ТП, передается только при добавлении/изменении сервисного комментария
* @apiParam {files[]}           files                           Файлы, передается только при добавлении/изменении файлов
* @apiParam {String}	        files.file_name 				Имя файла
* @apiParam {String}	        files.file_data 				Данные файла
*
* @apiSuccess (Success 200) {String}    status  Статус ответа на запрос
*
* @apiParamExample  {json} Request-Example:
* {
*     "ticket_id": 2,
*     "section_id": 1,
*     "ticket_status_id": 9,
*     "service_comment": "test"
* }
*
*
* @apiSuccessExample {json} Success-Response:
* {
*     "status": "ok"
* }
*
* @apiError ErrInvalidID Неверный ID категории
* @apiError ErrBlankTicketText Пустой текст запроса
* @apiError ErrBlankCategoryName Пустое имя категории
* @apiError ErrBlankSectionName Пустое имя раздела категории
*
 */
