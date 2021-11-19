package docs

/**
 *
 * @api {GET} /helpdesk/filial/filial_list Получение списка отделений из базы
 * @apiName GetFilialList
 * @apiGroup 07. Настройка
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (200) {Region[]} 	region 						Массив объектов "регион"
 * @apiSuccess (200) {Uint64} 		region.region_id 			ID региона в базе
 * @apiSuccess (200) {String} 		region.region 				Имя региона
 * @apiSuccess (200) {Filial[]} 	region.filials 				Массив отделений относящихся к региону
 * @apiSuccess (200) {Uint64} 		region.filials.filial_id 	ID отделения в базе
 * @apiSuccess (200) {String} 		region.filials.filial 		Название отделения
 * @apiSuccess (200) {String} 		region.filials.ip 			Первые 3 октета ip-адреса отделения
 *
 * @apiSuccessExample {json} Success-Response:
 * [
 *  {
 *    "region_id": 1,
 *    "region": "Николаевская область",
 *    "filials": [
 *      {
 *        "filial_id": 4,
 *        "filial": "Николаевское отделение №5",
 *        "ip": "10.54.6"
 *      },
 *      {
 *        "filial_id": 3,
 *        "filial": "Николаевское отделение №2",
 *        "ip": "10.54.2"
 *      },
 *      {
 *        "filial_id": 1,
 *        "filial": "Николаевское отделение №1",
 *        "ip": "10.54.1"
 *      }
 *    ]
 *  },
 *  {
 *    "region_id": 2,
 *    "region": "Киевская область",
 *    "filials": [
 *      {
 *        "filial_id": 2,
 *        "filial": "Киевское отделение №1",
 *        "ip": "10.1.1"
 *      }
 *    ]
 *  },
 *  {
 *    "region_id": 3,
 *    "region": "Одесская область"
 *  }
 * ]
 *
 *
 */
