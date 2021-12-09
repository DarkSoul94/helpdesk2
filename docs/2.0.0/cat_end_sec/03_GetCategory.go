package cat_end_sec

/**
*
* @api {GET} /helpdesk/category/ 03. Получение списка категорий
* @apiSampleRequest off
* @apiName GetCategory
* @apiGroup 02. Категории и разделы категорий
* @apiVersion  2.0.0
* @apiDescription Получение списка категорий.
*
* @apiHeader {String} BearerToken Авторизационный токен
*
* @apiSuccess (Success 200) {Category} 	category         			Массив объектов "категория запроса"
* @apiSuccess (Success 200) {Uint64}    category.category_id        ID категории
* @apiSuccess (Success 200) {String}    category.category_name      Имя категории
* @apiSuccess (Success 200) {Bool}      category.significant        Признак высокого приоритета у категории
* @apiSuccess (Success 200) {Bool}      category.old_category       Признак того что категория устарела
* @apiSuccess (Success 200) {Uint}      category.price       		Цена мотивации за запрос данной категории
*
* @apiSuccessExample {json} Success-Response:
*[
*    {
*        "category_id": 1,
*        "category_name": "Оборудование",
*        "significant": false,
*        "old_category": false,
*		 "price": 5
*    },
*    {
*        "category_id": 2,
*        "category_name": "1C",
*        "significant": false,
*        "old_category": false,
*		 "price": 1
*    }
*]
*
 */
