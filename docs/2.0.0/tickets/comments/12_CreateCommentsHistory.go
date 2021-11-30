package docs

/**
*
* @api {post} /helpdesk/comment/create 12. Создание нового комментария в запросе
* @apiName CreateCommentsHistory
* @apiGroup 03. Запросы в тех. поддержку
* @apiVersion  2.0.0
* @apiSampleRequest off
*
*
* @apiHeader {String} BearerToken 	Авторизационный токен
*
* @apiParam  {Uint64} ticket_id 	ID запроса к которому относится комментарий
* @apiParam  {String} comment_text 	Текст комментария
*
* @apiSuccess (Success 200) {String} status 		Статус ответа на запрос
* @apiSuccess (Success 200) {Uint64} comment_id		ID созданного комментария
*
* @apiParamExample  {json} Запрос на создание комментария:
* {
*   "ticket_id": 2,
*   "comment_text": "комментарий к запросу"
* }
*
*
* @apiSuccessExample {json} Success-Response:
* {
*     "comment_id": 5,
*     "status": "ok"
* }
*
* @apiError ErrBlankComment Поле <code>comment_text</code> пустое
*
 */
