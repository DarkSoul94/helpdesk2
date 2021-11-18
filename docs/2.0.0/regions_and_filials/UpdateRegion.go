package docs

/**
 *
 * @api {POST} /helpdesk/region/update Обновление региона в базе
 * @apiName UpdateRegion
 * @apiGroup 07. Настройка
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam  {Uint64} region_id 			ID региона в базе
 * @apiParam  {String} region 				Название региона
 *
 * @apiSuccess (Success 200) {String} status 		Статус ответа на запрос
 *
 * @apiParamExample  {json} Request-Example:
 * {
 * 		"region_id": 2,
 * 		"region": "Киевская область"
 * }
 *
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *     "status": "ok"
 * }
 *
 *
 */
