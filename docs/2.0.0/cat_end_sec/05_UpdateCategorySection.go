package cat_end_sec

/**
*
* @api {post} /helpdesk/section/update 05. Обновление разделов категории
* @apiSampleRequest off
* @apiName UpdateCategorySection
* @apiGroup 02. Категории и разделы категорий
* @apiVersion  2.0.0
* @apiDescription Обновление разделов категории.
*
* @apiHeader {String} BearerToken Авторизационный токен
*
* @apiParam {Uint64}  	section_id 			ID раздела категории
* @apiParam {String}  	section_name 		Имя нового раздела категории
* @apiParam {Bool}    	significant 		Признак высокого приоритета у раздела
* @apiParam {Bool}    	old_category 		Признак того что раздел категории устарел
* @apiParam {Bool}    	need_approval 		Признак того что для раздела необходимо согласование
* @apiParam {String}  	template 			Шаблон заполнения запроса
* @apiParam {Uint64}  	category_id 		ID категории к которой принадлежит раздел
* @apiParam {[]Uint64}	approval_groups 	ID групп которые будут согласовывать запросы с этим разделом категорий

*
* @apiParamExample {json} Request-Example:
*{
*    "section_id": 1,
*    "section_name": "Изменение по действующему займу",
*    "significant": false,
*    "old_category": false,
*    "need_approval": false,
*	 "template":"template for ticket",
*    "category_id": 2,
*	 "approval_groups": [1, 2]
*}
*
* @apiSuccess (Success 200) {String} status Статус выполнения запроса
* @apiSuccessExample {json} Success-Response:
*{
*    "status": "ok"
*}
*
* @apiError ErrBlankCategoryName Пустое имя категории
* @apiError ErrBlankSectionName Пустое имя раздела категории
* @apiError ErrSectionDoesNotExist Указанный раздел категории не существует
*
 */
