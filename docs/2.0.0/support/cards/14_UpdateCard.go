package docs

/**
 *
 * @api {POST} /support/card/update 14. Обновить карточку суппорта
 * @apiName UpdateSupportCard
 * @apiGroup 05. Сотрудник ТП
 * @apiVersion  0.1.1
 * @apiSampleRequest off
 *
 * @apiParam  {Uint64} 	id 								ИД карточки
 * @apiParam  {Stirng} 	internal_number 	Внутренний номер телефонии
 * @apiParam  {String} 	mobile_number 		Мобильный номер (в одном из форматов: 0671112233, +380671112233, 380671112233, 067 111 22 33, 067-111-22-33)
 * @apiParam  {String} 	birth_date 				Дата рождения (в формате ДД.ММ.ГГГГ)
 * @apiParam  {Bool} 		is_senior 				Признак старшего суппорта
 * @apiParam  {Uint64} 	senior_id 				ИД старшего суппорта (если есть)
 * @apiParam  {Float64} wager							Ставка за смену
 * @apiParam  {String} 	comment						Комментарий
 * @apiParam  {String} 	color							Цвет отображения (если назначен старший цвет берется из карточки старшего)
 *
 * @apiSuccess (Success 200) {String} status Статус выполнения запроса
 *
 * @apiParamExample  {json} Request-Example:
 * {
 *   "id": 14,
 *   "internal_number": "1487",
 *   "mobile_number": "",
 *   "birth_date": "",
 *   "is_senior": false,
 *   "senior_id": 4,
 *   "wager": 500,
 *   "comment": "test",
 *   "color": "0xFFFFF"
 * }
 *
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "status": "ok"
 * }
 *
 *
 */
