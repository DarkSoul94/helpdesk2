package cat_end_sec

/**
*
* @api {POST} /helpdesk/section/create Создание раздела категории
* @apiName CreateSection
* @apiGroup 02. Категории и разделы категорий
* @apiVersion  2.0.0
* @apiDescription Создание раздела категории.
* @apiSampleRequest off
*
* @apiHeader {String} BearerToken Авторизационный токен
*
* @apiParam {Uint64} 	section_id 		ID раздела категории. Для создания раздела должен равнятся 0
* @apiParam {String} 	section_name 	Имя нового раздела категории. Уникально в рамках одной категории
* @apiParam {Bool} 		significant 	Признак высокого приоритета у раздела
* @apiParam {Bool} 		old_category 	Признак того что раздел категории устарел
* @apiParam {Bool} 		need_approval 	Признак того что для раздела необходимо согласование
* @apiParam {String} 	template 		Шаблон заполнения запроса
* @apiParam {Uint64} 	category_id 	ID категории
*
*
* @apiParamExample {json} Запрос на создание только раздела:
* {
*    "section_id": 0,
*    "section_name": "Перемещение менеджера",
*    "significant": false,
*    "old_category": false,
*    "need_approval": false,
*	 "template":"template for ticket",
*	 "category_id": 2,
* }
*
* @apiSuccess (Success 200) {Uint64} 	category_section_id 	ID нового раздела категории
* @apiSuccess (Success 200) {String}	status					Статус запроса
*
* @apiSuccessExample {json} Ответ при создании раздела:
* {
*    "category_section_id": 14,
*    "status": "ok"
* }
*
*
* @apiError ErrBlankSectionName Пустое имя раздела категории
* @apiError ErrBlankCategoryName Пустое имя категории
* @apiError ErrSectionAlreadyExist Раздел категории с таким именем уже существует
*
 */
