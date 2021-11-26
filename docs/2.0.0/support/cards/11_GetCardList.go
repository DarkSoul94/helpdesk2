package docs

/**
 *
 * @api {GET} /support/card/list 11. Получить список карточек суппорта
 * @apiName GetCardList
 * @apiGroup 05. Сотрудник ТП
 * @apiVersion  0.1.1
 * @apiSampleRequest off
 *
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {[]Card} 		cards 									Массив карточек суппорта
 * @apiSuccess (Success 200) {Uint64} 		cards.id 								ИД карточки супорта
 * @apiSuccess (Success 200) {CardUser} 	cards.support 					Объект содержащий данные суппорта
 * @apiSuccess (Success 200) {Uint64} 		cards.support.id 				ИД суппорта
 * @apiSuccess (Success 200) {String} 		cards.support.name 			ФИО суппорта
 * @apiSuccess (Success 200) {Bool} 			cards.is_senior 				Признак старшего группы саппортов
 * @apiSuccess (Success 200) {CardUser} 	cards.senior 						Объект содержащий данные о старшем суппорте в группу которого входит текущий сотрудник
 * @apiSuccess (Success 200) {Uint64} 		cards.senior.id 				ИД суппорта
 * @apiSuccess (Success 200) {String} 		cards.senior.name 			ФИО суппорта
 * @apiSuccess (Success 200) {String} 		cards.color 						Цвет отображения (в шестнадцатиричной системе)
 *
 *
 * @apiSuccessExample {json} Список карточек:
 * [
 *   {
 *     "id": 13,
 *     "support": {
 *       "id": 4,
 *       "name": "Вячеслав Викторович Тищенко"
 *     },
 *     "is_senior": true,
 *     "senior": null,
 *     "color": "0xFFFFF0"
 *   },
 *   {
 *     "id": 14,
 *     "support": {
 *       "id": 5,
 *       "name": "Евгений Николаевич Табаков"
 *     },
 *     "is_senior": false,
 *     "senior": {
 *       "id": 4,
 *       "name": "Вячеслав Викторович Тищенко"
 *     },
 *     "color": "0xFFFFF0"
 *   }
 * ]
 *
 */
