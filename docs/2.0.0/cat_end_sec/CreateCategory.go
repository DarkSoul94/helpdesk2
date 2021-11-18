package cat_end_sec

/**
*
* @api {POST} /helpdesk/category/create Создание категории
* @apiName CreateCategory
* @apiGroup 02. Категории и разделы категорий
* @apiVersion  2.0.0
* @apiDescription Создание категории.
* @apiSampleRequest off
*
* @apiHeader {String} BearerToken Авторизационный токен
*
* @apiParam {String}        category_name      	Имя новой категории, при создании только раздела категории, можно опустить
* @apiParam {Bool}          significant        	Признак высокого приоритета у категории, при создании только раздела категории, можно опустить
* @apiParam {Bool}          old_category       	Признак того что категория устарела, при создании только раздела категории, можно опустить
* @apiParam {Uint}          price       		Цена мотивации за запрос данной категории, при создании только раздела категории, можно опустить
*
*
* @apiParamExample {json} Запрос на создание категории:
* {
*	"category_name": 1C,
*	"significant": false,
*	"old_category": false,
*	"price": 5
* }
*
* @apiSuccess (Success 200) {Uint64} 	categoryid	ID нового раздела категории
* @apiSuccess (Success 200) {String}	status		Статус выполнения запроса
*
* @apiSuccessExample {json} Ответ при создании категории:
* {
*    "category_id": 14,
*    "status": "ok"
* }
*
*
* @apiError ErrBlankCategoryName Пустое имя категории
* @apiError ErrSectionAlreadyExist Категория с таким именем уже существует
*
 */
