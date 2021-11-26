package docs

/**
 *
 * @api {DELETE} /helpdesk/region/ 03. Удаление региона из базы
 * @apiName DeleteRegion
 * @apiGroup 07. Настройка
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 * @apiDescription При удалении региона удаляются также и филиалы которые в него входят
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiExample  Example usage:
 * http://localhost:8888/helpdesk/region/?region_id=23
 *
 * @apiParam  {Uint64} region_id 	ID отделения в базе
 *
 * @apiSuccess (Success 200) {String} status 	Статус ответа на запрос
 *
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *     "status": "ok"
 * }
 *
 */
