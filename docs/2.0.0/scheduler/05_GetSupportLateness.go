package scheduler

/**
 *
 * @api {GET} /helpdesk/table/lateness 05. Получить список опозданий сотрудников ТП за месяц
 * @apiName GetSupportLateness
 * @apiGroup 8. Табель
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {[]Decisions} 	desicions 						Список возможных решений по опозданию
 * @apiSuccess (Success 200) {Uint64} 			desicions.id 					ID решения
 * @apiSuccess (Success 200) {String} 			desicions.text	 			Текст решения
 * @apiSuccess (Success 200) {[]Lateness} 	lateness	 						Список опозданий сотрудников ТП
 * @apiSuccess (Success 200) {Uint64} 			lateness.id	 					ID записи об опоздании
 * @apiSuccess (Success 200) {String} 			lateness.date	 				Дата и время создания записи об опоздании
 * @apiSuccess (Success 200) {String} 			lateness.name					Имя сотрудника ТП
 * @apiSuccess (Success 200) {String} 			lateness.cause				Причина опоздания ТП
 * @apiSuccess (Success 200) {Uint64} 			lateness.decision_id	ID решения (если `0` - решение по опозданию отсутствует)
 * @apiSuccess (Success 200) {Uint64} 			lateness.difference		Кол-во минут на которые опоздал сотрудник
 * @apiSuccess (Success 200) {String} 			status								Статус выполнения запроса
 *
 * @apiSuccessExample  {json} Success-Response:
 * {
 *   "desicions": [
 *     {
 *       "id": 1,
 *       "text": "Помилован"
 *     },
 *     {
 *       "id": 2,
 *       "text": "Казнен"
 *     }
 *   ],
 *   "lateness": [
 *     {
 *       "id": 1,
 *       "date": "2021-10-05 11:08:12",
 *       "name": "Вячеслав Викторович Тищенко"
 *       "cause": "test",
 *       "decision_id": 1,
 *       "difference": 3
 *     },
 *     {
 *       "id": 2,
 *       "date": "2021-10-02 11:08:12",
 *       "name": "Вячеслав Викторович Тищенко"
 *       "cause": "test2",
 *       "decision_id": 0,
 *       "difference": 10
 *     }
 *   ],
 *   "status": "ok"
 * }
 *
 *
 */
