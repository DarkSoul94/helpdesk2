package docs

/**
 *
 * @api {GET} /support/card Получить карточку сотрудника тех.поддержки
 * @apiName GetSupportCard
 * @apiGroup 05. Сотрудник ТП
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiExample  Example usage:
 * http://localhost:8888/helpdesk/support/card?id=13
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {Uint64} 		id 								ИД карточки супорта
 * @apiSuccess (Success 200) {CardUser} 	support 					Объект содержащий данные суппорта
 * @apiSuccess (Success 200) {Uint64} 		support.id 				ИД суппорта
 * @apiSuccess (Success 200) {String} 		support.name 			ФИО суппорта
 * @apiSuccess (Success 200) {String} 		internal_number 	Внутренний номер телефонии
 * @apiSuccess (Success 200) {String} 		mobile_number 		Мобильный номер телефона
 * @apiSuccess (Success 200) {String} 		birth_date 				Дата рождения
 * @apiSuccess (Success 200) {Bool} 			is_senior 				Признак старшего группы саппортов
 * @apiSuccess (Success 200) {CardUser} 	senior 						Объект содержащий данные о старшем суппорте в группу которого входит текущий сотрудник
 * @apiSuccess (Success 200) {Uint64} 		senior.id 				ИД суппорта
 * @apiSuccess (Success 200) {String} 		senior.name 			ФИО суппорта
 * @apiSuccess (Success 200) {Float64}		wager 						Ставка за смену
 * @apiSuccess (Success 200) {String} 		comment 					Комментарий
 * @apiSuccess (Success 200) {String} 		color 						Цвет отображения (в шестнадцатиричной системе)
 *
 *
 * @apiSuccessExample {json} Карточка старшего смены:
 * {
 *   "id": 13,
 *   "support": {
 *     "id": 4,
 *     "name": "Вячеслав Викторович Тищенко"
 *   },
 *   "internal_number": "1484",
 *   "mobile_number": "",
 *   "birth_date": "",
 *   "is_senior": true,
 *   "senior": null,
 *   "wager": 500,
 *   "comment": "test",
 *   "color": "0xFFFFF0"
 * }
 *
 * @apiSuccessExample {json} Карточка обычного суппорта:
 *{
 *  "id": 14,
 *  "support": {
 *    "id": 5,
 *    "name": "Евгений Николаевич Табаков"
 *  },
 *  "internal_number": "1487",
 *  "mobile_number": "",
 *  "birth_date": "",
 *  "is_senior": false,
 *  "senior": {
 *    "id": 4,
 *    "name": "Вячеслав Викторович Тищенко"
 *  },
 *  "wager": 500,
 *  "comment": "test",
 *  "color": "0xFFFFF0"
 *}
 *
 */
