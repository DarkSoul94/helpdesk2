package docs

/**
 *
 * @api {DELETE} /helpdesk/filial/ 06. Удаление филиала из базы
 * @apiName DeleteFilial
 * @apiGroup 07. Настройка
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiExample  Example usage:
 * http://localhost:8888/helpdesk/filial/?filial_id=23
 *
 * @apiParam  {Uint64} filial_id 	ID отделения в базе
 *
 * @apiSuccess (Success 200) {String} status 	Статус ответа на запрос
 *
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *     "status": "ok"
 * }
 *
 * @apiError ErrFilialDoesNotExist Такой филиал отсутствует в базе
 *
 */
