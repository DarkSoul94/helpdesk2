package cat_end_sec

/**
 *
 * @api {GET} /helpdesk/section/section_list 06. Получение всего списка разделов категорий
 * @apiSampleRequest off
 * @apiName GetCategorySectionList
 * @apiGroup 02. Категории и разделы категорий
 * @apiVersion  2.0.0
 * @apiDescription Получение всего списка разделов категорий для проведения их настройки.
 *
 * @apiHeader {String} BearerToken Авторизационный токен
 *
 * @apiSuccess (Success 200) {Category[]}	category                  			Массив объектов "категория"
 * @apiSuccess (Success 200) {Uint64} 		category.category_id				ID категории
 * @apiSuccess (Success 200) {String}		category.category_name 				Имя категории
 * @apiSuccess (Success 200) {Bool}  		category.significant 				Высокий приоритет выполнения у категории
 * @apiSuccess (Success 200) {Bool}  		category.old_category 				Устаревшая категория
 * @apiSuccess (Success 200) {Uint}      	category.price 						Цена мотивации за запрос данной категории
 * @apiSuccess (Success 200) {Section[]} 	category.sections                 	Массив объектов "раздел категории"
 * @apiSuccess (Success 200) {Uint64} 		category.section.section_id         ID раздела категории
 * @apiSuccess (Success 200) {String} 		category.section.section_name       Имя нового раздела категории
 * @apiSuccess (Success 200) {Bool}   		category.section.significant        Высокий приоритет выполнения у раздела категории
 * @apiSuccess (Success 200) {Bool}   		category.section.old_category       Устаревший раздел категории
 * @apiSuccess (Success 200) {Bool}   		category.section.need_approval      Необходимость согласования
 * @apiSuccess (Success 200) {Uint64}      	category.section.category_id        ID категории в которую входит раздел
 * @apiSuccess (Success 200) {String} 		category.section.template 			Шаблон заполнения запроса
 *
 * @apiSuccessExample {json} Success-Response:
 *[
 *    {
 *    "category_id": 4,
 *    "category_name": "Устаревшая",
 *    "significant": false,
 *    "old_category": true,
 *    "price": 0,
 *    "sections": [
 *      {
 *        "section_id": 6,
 *        "section_name": "Уборка комнаты с игрушками",
 *        "significant": false,
 *        "old_category": false,
 *        "need_approval": false,
 *        "category_id": 4,
 *	 	 "template":"template for ticket",
 *      }
 *    ]
 *  },
 *  {
 *    "category_id": 5,
 *    "category_name": "Валютообмен",
 *    "significant": false,
 *    "old_category": false,
 *    "price": 0,
 *    "sections": [
 *      {
 *        "section_id": 7,
 *        "section_name": "Спецоперация",
 *        "significant": false,
 *        "old_category": false,
 *        "need_approval": false,
 *        "category_id": 5,
 *	 	 "template":"template for ticket",
 *      }
 *    ]
 *  }
 *]
 *
 */
