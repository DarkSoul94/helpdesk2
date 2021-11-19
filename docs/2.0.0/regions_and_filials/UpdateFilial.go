package docs

/**
 *
 * @api {POST} /helpdesk/filial/update Обновление отделения в базе
 * @apiName UpdateFilial
 * @apiGroup 07. Настройка
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam  {Uint64} filial_id 			ID отделения в базе
 * @apiParam  {Region} region 				Регион в котором находится отделение
 * @apiParam  {Uint64} region.region_id		ID региона в котором находится отделение
 * @apiParam  {String} filial 				Название отделения
 * @apiParam  {String} ip 					Первые 3 октета ip-адреса отделения
 *
 * @apiSuccess (Success 200) {String} status 		Статус ответа на запрос
 *
 * @apiParamExample  {json} Request-Example:
 * {
 * 	   "filial_id":1,
 *     "region":{
 *		"region_id": 1
 * },
 *     "filial":"Николаевское отделение №1",
 *     "ip":"10.54.5"
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
