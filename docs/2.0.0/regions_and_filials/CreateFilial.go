package docs

/**
 *
 * @api {POST} /helpdesk/filial/create Добаление отделения в базу
 * @apiName CreateFilial
 * @apiGroup 07. Настройка
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam  {Uint64} 	region_id 	ID региона в базе
 * @apiParam  {String}  filial 		Название отделения
 * @apiParam  {String}  ip 			Первые 3 октета ip-адреса отделения
 *
 * @apiSuccess (Success 200) {String} status 		Статус ответа на запрос
 * @apiSuccess (Success 200) {Uint64} filial_id 	ID созданого отделения
 *
 * @apiParamExample  {json} Создание филиала:
 * {
 *	 	"region_id": 1,
 *     	"filial":"Николаевское отделение №3",
 *     	"ip":"10.54.3"
 * }
 *
 * @apiSuccessExample {json} Создание филиала:
 * {
 *     "filial_id": 3,
 *     "status": "ok"
 * }
 *
 * @apiError FilialErr_Exist Отделение с таким ip уже существует
 *
 */
