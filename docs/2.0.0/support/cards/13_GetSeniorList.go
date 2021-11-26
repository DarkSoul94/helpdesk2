package docs

/**
 *
 * @api {GET} /support/card/seniors 13. Получить список старших суппортов
 * @apiName GetSeniorSupportsList
 * @apiGroup 05. Сотрудник ТП
 * @apiVersion  0.1.1
 * @apiSampleRequest off
 *
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {[]SeniorSupport} 		seniors 				Массив старших суппортов
 * @apiSuccess (Success 200) {Uint64} 						seniors.id 			ИД старшего суппорта
 * @apiSuccess (Success 200) {String} 						seniors.name 		ФИО суппорта
 *
 *
 * @apiSuccessExample {json} Список старших суппортов:
 * [
 *   {
 *     "id": 4,
 *     "name": "Вячеслав Викторович Тищенко"
 *   }
 * ]
 *
 */
