package docs

/**
 *
 * @api {get} /file/ 17. Получение файла по его ид
 * @apiName GetFile
 * @apiGroup 03. Запросы в тех. поддержку
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 * @apiDescription В запросе указывается ид файла который необходимо получить,
 * в ответ возвращается файл в виде json объекта или ошибка.
 *
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiExample  Example usage:
 * http://localhost:8888/helpdesk/file/?file_id=23
 *
 * @apiParam  {Uint64} file_id ID получаемого файла
 *
 * @apiSuccess (Success 200) {Uint64}	file_id		ID файла
 * @apiSuccess (Success 200) {String}	file_name 	Имя файла
 * @apiSuccess (Success 200) {Uint64}	ticket_id 	ID запроса к которому относится файл
 * @apiSuccess (Success 200) {String}	file_data 	Данные файлы
 * @apiSuccess (Success 200) {String}	file_date 	Дата добавления файла
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "file_id": 48,
 *   "file_name": "скрин1.новый_клиент.jpg",
 *   "file_data": "a few byte count",
 *   "file_date": "2021-04-12T13:33:20Z"
 * }
 *
 *
 */
