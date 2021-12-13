package docs

/**
 *
 * @api {GET} /helpdesk/reports/supports_statuses_history 09. История статусов суппортов за определенную дату
 * @apiName GetSupportStatusHistory
 * @apiGroup 06. Отчёты
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken Авторизационный токен
 *
 * @apiParam  {String} date 	Дата за которую необходимо получить статусы
 *
 * @apiExample  Example usage:
 * http://localhost:8888/helpdesk/reports/supports_statuses_history?date=2021-07-20
 *
 * @apiSuccess (Success 200) {SupportStatus[]} 	support_status 						Массив объектов "статус суппорта"
 * @apiSuccess (Success 200) {String} 			support_status.support 				Имя суппорта
 * @apiSuccess (Success 200) {Statuses[]} 		support_status.statuses 			Массив объектов "статусы"
 * @apiSuccess (Success 200) {String} 			support_status.statuses.time		Время выбора статуса
 * @apiSuccess (Success 200) {String} 			support_status.statuses.name		Имя выбранного статуса
 * @apiSuccess (Success 200) {Uint64} 			support_status.statuses.difference 	Длительность нахождения в статусе в секундах

 *
 *
 * @apiSuccessExample {json} Success-Response:
 * [
 *  {
 *    "support": "Артем Владимирович Шелкопляс",
 *    "statuses": [
 *      {
 *        "time": "14:27:44",
 *        "name": "Принимаю запросы",
 *        "difference": 146
 *      },
 *      {
 *        "time": "14:30:10",
 *        "name": "Работа в офисе",
 *        "difference": 101
 *      },
 *      {
 *       "time": "14:31:51",
 *       "name": "Принимаю запросы",
 *       "difference": 0
 *     }
 *    ]
 *  },
 *  {
 *    "support": "Александр Игоревич Кудряшов",
 *    "statuses": [
 *      {
 *        "time": "14:28:51",
 *        "name": "Принимаю запросы",
 *        "difference": 89
 *      },
 *      {
 *        "time": "14:30:20",
 *        "name": "Работа в офисе",
 *        "difference": 15
 *      },
 *      {
 *       "time": "14:30:35",
 *       "name": "Принимаю запросы",
 *       "difference": 0
 *     }
 *    ]
 *  },
 * ]
 */
