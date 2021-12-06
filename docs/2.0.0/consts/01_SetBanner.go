package docs

/**
 *
 * @api {POST} /helpdesk/const/banner 01. Модификация текста баннера
 * @apiName SetBanner
 * @apiGroup 07. Настройка
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam  {String} text Текст баннера
 *
 * @apiSuccess (Success 200) {String} status Статус ответа на запрос
 *
 * @apiParamExample  {json} Request-Example:
 * {
 *	 		"text":"text"
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
