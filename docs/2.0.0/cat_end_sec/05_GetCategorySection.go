package cat_end_sec

/**
 *
 * @api {GET} /helpdesk/section/ 05. Получение списка разделов категорий без учета устаревших
 * @apiSampleRequest off
 * @apiName GetCategorySection
 * @apiGroup 02. Категории и разделы категорий
 * @apiVersion  2.0.0
 * @apiDescription Получение списка разделов категорий. Разделы категорий которые помечены устаревшими или которые входят в устаревшую категорию не отображаются.
 *
 * @apiHeader {String} BearerToken Авторизационный токен
 *
 * @apiSuccess (Success 200) {Section[]}		section                  			Массив объектов "раздел категории"
 * @apiSuccess (Success 200) {Uint64}        	section.section_id                  ID раздела категории
 * @apiSuccess (Success 200) {String}        	section.section_name                Имя нового раздела категории
 * @apiSuccess (Success 200) {Bool}          	section.significant                 Высокий приоритет выполнения у раздела категории
 * @apiSuccess (Success 200) {Bool}          	section.old_category                Устаревший раздел категории
 * @apiSuccess (Success 200) {Bool}          	section.need_approval               Необходимость согласования
 * @apiSuccess (Success 200) {String}        	section.template               		Шаблон заполнения запроса
 * @apiSuccess (Success 200) {category}      	section.category                    Категория запроса
 * @apiSuccess (Success 200) {Uint64}        	section.category.category_id        ID категории
 * @apiSuccess (Success 200) {String}        	section.category.category_name      Имя категории
 * @apiSuccess (Success 200) {Bool}          	section.category.significant        Высокий приоритет выполнения у категории
 * @apiSuccess (Success 200) {Bool}          	section.category.old_category       Устаревшая категория
 * @apiSuccess (Success 200) {Uint}      		section.category.price       		Цена мотивации за запрос данной категории
 *
 * @apiSuccessExample {json} Success-Response:
 *[
 *    {
 *        "section_id": 1,
 *        "section_name": "Проблемы с кассовым аппаратом",
 *        "significant": false,
 *        "old_category": false,
 *        "need_approval": false,
 *	 	  "template":"template for ticket",
 *        "category": {
 *            "category_id": 1,
 *            "category_name": "Оборудование",
 *            "significant": false,
 *            "old_category": false,
 *		 	 "price": 5
 *        }
 *    },
 *    {
 *        "section_id": 2,
 *        "section_name": "Настройка нового оборудования",
 *        "significant": false,
 *        "old_category": false,
 *        "need_approval": false,
 *	 	  "template":"template for ticket",
 *        "category": {
 *            "category_id": 1,
 *            "category_name": "Оборудование",
 *            "significant": false,
 *            "old_category": false,
 *		 	 "price": 5
 *        }
 *    }
 *]
 *
 */
