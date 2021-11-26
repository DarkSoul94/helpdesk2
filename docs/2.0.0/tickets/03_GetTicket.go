package docs

/**
 *
 * @api {GET} /helpdesk/ticket/ticket 03. Получение запроса
 * @apiName GetTicket
 * @apiGroup 03. Запросы в тех. поддержку
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 * @apiDescription Метод который возвращает запрос по указаному ID с прикрепленными к нему коментариями и файлами.
 * Если у пользователся нет доступа `see_additional_info` - у него не будут отображатся поля: `ticket_author`, `support`, `resolved_user`, `service_comment`.
 *
 * @apiParam {Int}	ticket_id	Номер запрашиваемого запроса
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiExample  Example usage:
 * http://localhost:8888/helpdesk/ticket/ticket?ticket_id=1
 *
 * @apiSuccess (Success 200) {Uint64} 	        	ticket_id 									ID запроса
 * @apiSuccess (Success 200) {String} 	        	ticket_date 								Дата создания запроса
 * @apiSuccess (Success 200) {CategorySection}  	category_section 							Раздел категории
 * @apiSuccess (Success 200) {Uint64} 	        	category_section.section_id 				ID раздела категории
 * @apiSuccess (Success 200) {String} 	        	category_section.section_name 				Имя раздела категории
 * @apiSuccess (Success 200) {Bool} 				category_section.significant 				Приоритет выполнения
 * @apiSuccess (Success 200) {Bool} 				category_section.old_category 				Утаревший раздел категории
 * @apiSuccess (Success 200) {Bool} 				category_section.need_approval 				Необходимость согласования
 * @apiSuccess (Success 200) {Category}          	category_section.category                 	Категория
 * @apiSuccess (Success 200) {Uint64} 	        	category_section.category.category_id 		ID категории
 * @apiSuccess (Success 200) {String} 	        	category_section.category.category_name 	Имя категории
 * @apiSuccess (Success 200) {Bool} 				category_section.category.significant 		Приоритет выполнения
 * @apiSuccess (Success 200) {Bool} 				category_section.category.old_category 		Утаревшая категория
 * @apiSuccess (Success 200) {Bool} 				category_section.category.price 			Количество мотивации за запрос данной категории
 * @apiSuccess (Success 200) {String} 	        	ticket_text 								Текст запроса
 * @apiSuccess (Success 200) {TicketStatus}     	ticket_status 								Статус запроса
 * @apiSuccess (Success 200) {Uint64} 	        	ticket_status.ticket_status_id 		        ID статуса запроса
 * @apiSuccess (Success 200) {String} 	        	ticket_status.ticket_status_name 	        Имя статуса запроса
 * @apiSuccess (Success 200) {String} 				filial 										Отделение за которым закрелен автор запроса. Если в базе филиал не найден вернет <code>"not_found"</code>. Если у пользователя нет прав на просмотр информации вернет пустое поле
 * @apiSuccess (Success 200) {String} 				ip 											IP адрес автора запроса
 * @apiSuccess (Success 200) {User} 				ticket_author 								Автор запроса. Если у пользователя нет прав на просмотр информации вернет <code>null</code>
 * @apiSuccess (Success 200) {Uint64} 	        	ticket_author.user_id 						ID автора запроса
 * @apiSuccess (Success 200) {String} 	        	ticket_author.user_name 					Имя автора запроса
 * @apiSuccess (Success 200) {String} 	        	ticket_author.email 						Почта автора запроса
 * @apiSuccess (Success 200) {Uint64} 	        	ticket_author.group_id 						ID группы к которой принадлежит автор запроса
 * @apiSuccess (Success 200) {String} 	        	ticket_author.department 					Подразделение к которому относится автор запроса
 * @apiSuccess (Success 200) {User} 				support 									Сотрудник тех. поддержки. Если у пользователя нет прав на просмотр информации вернет <code>null</code>
 * @apiSuccess (Success 200) {Uint64} 	        	support.user_id 							ID сотрудника тех.поддержки
 * @apiSuccess (Success 200) {String} 				support.user_name 							Имя сотрудника тех.поддержки
 * @apiSuccess (Success 200) {String} 	        	support.email 								Почта сотрудника тех.поддержки
 * @apiSuccess (Success 200) {Uint64} 	        	support.group_id 							ID группы к которой принадлежит пользователь
 * @apiSuccess (Success 200) {User} 				resolved_user 								Пользователь согласовавший запрос. Если у пользователя нет прав на просмотр информации вернет <code>null</code>
 * @apiSuccess (Success 200) {Uint64} 				resolved_user.user_id 						ID пользователя согласовавшего запрос
 * @apiSuccess (Success 200) {String} 				resolved_user.user_name 					Имя пользователя согласовавшего запрос
 * @apiSuccess (Success 200) {String} 	        	resolved_user.email 						Почта пользователя согласовавшего запрос
 * @apiSuccess (Success 200) {Uint64} 	        	resolved_user.group_id 						ID группы к которой принадлежит пользователь
 * @apiSuccess (Success 200) {String} 	        	service_comment 							Сервисный комментарий для сотрудников тех. поддержки
 * @apiSuccess (Success 200) {comments[]}        	comments 									Комментарии
 * @apiSuccess (Success 200) {Uint64} 				comments.comment_id 						ID комментария
 * @apiSuccess (Success 200) {String} 				comments.comment_date 						Дата добавления комментария
 * @apiSuccess (Success 200) {String} 				comment_author 								Автор комментария
 * @apiSuccess (Success 200) {String} 				comments.comment_text 						Текст комментария
 * @apiSuccess (Success 200) {Files[]}           	files 										Файлы
 * @apiSuccess (Success 200) {Uint64} 				files.file_id 								ID файла
 * @apiSuccess (Success 200) {String} 				files.file_name 							Имя файла
 * @apiSuccess (Success 200) {String} 				files.file_date 							Дата добавления файла
 *
 *
 * @apiSuccessExample {json} Вид запроса для админа и сотрудника ТП:
 * {
 *  "ticket_id": 2,
 *  "ticket_date": "2021-05-18T16:49:30+03:00",
 *  "category_section": {
 *    "section_id": 2,
 *    "section_name": "Удаление кассовых",
 *    "significant": false,
 *    "old_category": false,
 *    "need_approval": true,
 *    "category": {
 *      "category_id": 2,
 *      "category_name": "1С"
 *    }
 *  },
 *  "ticket_text": "delete",
 *  "ticket_status": {
 *    "ticket_status_id": 9,
 *    "ticket_status_name": "Выполнен"
 *  },
 *  "filial": "not found",
 *   "ip": "10.54.86.26",
 *  "ticket_author": {
 *    "user_id": 5,
 *    "user_name": "Владислав Сергеевич Маспанов",
 *    "email": "maspanov.v.s@limefin.com",
 *    "group_id": 3,
 *    "department": "Техническая поддержка"
 *  },
 *  "support": {
 *    "user_id": 6,
 *    "user_name": "Вячеслав Викторович Тищенко",
 *    "email": "tishchenko.v.v@limefin.com",
 *    "group_id": 2
 *  },
 *  "resolved_user": {
 *    "user_id": 6,
 *    "user_name": "Вячеслав Викторович Тищенко",
 *    "email": "tishchenko.v.v@limefin.com",
 *    "group_id": 2
 *  },
 *  "service_comment": "",
 *  "comments": [],
 *  "files": []
 *}
 * *
 * * @apiSuccessExample {json} Вид запроса для остальных пользователей:
 * {
 * "ticket_id": 2,
 * "ticket_date": "2021-05-18T16:49:30+03:00",
 * "category_section": {
 *	 "section_id": 2,
 *    "section_name": "Удаление кассовых",
 *    "significant": false,
 *    "old_category": false,
 *    "need_approval": true,
 *    "category": {
 *      "category_id": 2,
 *      "category_name": "1С"
 *    }
 *  },
 *  "ticket_text": "delete",
 *  "ticket_status": {
 *    "ticket_status_id": 9,
 *    "ticket_status_name": "Выполнен"
 *  },
 *  "filial": " ",
 *  "ip": " ",
 *  "ticket_author": null,
 *  "support": null,
 *  "resolved_user": null,
 *  "service_comment": "",
 *  "comments": [],
 *  "files": []
 *}
 *
 */
