package cat_end_sec

/**
*
* @api {POST} /helpdesk/category/update 02. Обновление категории
* @apiSampleRequest off
* @apiName UpdateCategory
* @apiGroup 02. Категории и разделы категорий
* @apiVersion  2.0.0
* @apiDescription Обновление категории.
*
* @apiHeader {String} BearerToken Авторизационный токен
*
* @apiParam {Uint64}    category_id         ID категории.
* @apiParam {String}    category_name       Имя категории.
* @apiParam {Bool}      significant         Признак высокого приоритета у категории.
* @apiParam {Bool}      old_category        Признак того что категория устарела.
* @apiParam {Uint}      price       		Цена мотивации за запрос данной категории.
*
* @apiParamExample {json} Request-Example:
*{
*	"category_id": 2,
*	"category_name": "1C",
*	"significant": false,
*	"old_category": true,
*	"price": 5
*}
*
* @apiSuccess (Success 200) {String} status Статус выполнения запроса
*
* @apiSuccessExample {json} Success-Response:
*{
*    "status": "ok"
*}
*
* @apiError ErrBlankCategoryName Пустое имя категории
* @apiError ErrCategoryDoesNotExist Указанная категория не существует
*
 */
