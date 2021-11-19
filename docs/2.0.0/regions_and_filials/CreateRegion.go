package docs

/**
 *
 * @api {POST} /helpdesk/region/create Создание региона
 * @apiName CreateRegion
 * @apiGroup 07. Настройка
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam  {String} region Название региона
 *
 * @apiSuccess (Success 200) {String} status 		Статус ответа на запрос
 * @apiSuccess (Success 200) {Uint64} region_id 	ID созданого отделения
 *
 * @apiParamExample  {type} Request-Example:
 * {
 *         "region":"Николаевская область"
 * }
 *
 *
 * @apiSuccessExample {type} Success-Response:
 * {
 *     "status" : "ok",
 *     "region_id" : 1,
 * }
 *
 *
 */
