package docs

/**
 *
 * @api {GET} /helpdesk/resolve_ticket/check_exist Проверка есть ли запросы ожидающие согласования
 * @apiName CheckNeedResolveTicketExist
 * @apiGroup 03. Запросы в тех. поддержку
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {Bool} 	exist 	Есть ли запросы в базе, если есть то `true`
 * @apiSuccess (Success 200) {String} 	status 	Статус ответа на запрос
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 * 		"exist": true,
 *     	"status": "ok"
 * }
 *
 *
 */
