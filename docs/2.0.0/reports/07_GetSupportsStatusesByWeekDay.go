package docs

/**
 *
 * @api {GET} /helpdesk/reports/supports_statuses 07. Время нахождения сотрудника ТП в определенном статусе в разрезе дней недели
 * @apiName GetSupportsStatusesByWeekDay
 * @apiGroup 06. Отчёты
 * @apiVersion  2.0.0
 * @apiSampleRequest off
 * @apiDescription Отчет показывает сколько времени было проведено сотрудником в том или ином статусе в разрезе дней недели за указанный промежуток времени.
 * Если в какой-то из дней недели небыло сотрудников у которых менялся статус - то в этот день список сотрудников будет пустым.
 *
 * @apiExample  Example usage:
 * http://localhost:8888/helpdesk/reports/supports_statuses?start_date=2021-05-01&end_date=2021-06-01
 *
 * @apiHeader {String} BearerToken Авторизационный токен
 *
 * @apiParam  {String} start_date 	Дата начала выборки, включительно
 * @apiParam  {String} end_date		Дата конца выборки, данные за этот день не учитываются
 *
 * @apiSuccess (Success 200) {Uint} 					week_day 							ID дня недели: 0 - понедельник, 1 - вторник, 2 - среда, 3 - четверг, 4 - пятница, 5 - суббота, 6 - воскресенье
 * @apiSuccess (Success 200) {SupportStatusesList[]} 	supports_list						Список сотрудников тех поддержки с указанием статусов в которых они были за указанный промежуток времени и времени которое в них провели
 * @apiSuccess (Success 200) {String}					supports_list.support_name			ФИО Сотрудника тех поддержки
 * @apiSuccess (Success 200) {SupportStatusDuration[]}	supports_list.statuses				Список статусов которые были у сотрудника тех поддержки за указанный промежуток времени с указанием времени проведенном в этом статусе
 * @apiSuccess (Success 200) {String}					supports_list.statuses.status_name	Название статуса
 * @apiSuccess (Success 200) {String}					supports_list.statuses.duration		Время проведенное в данном статусе в формате "00h00m00s"
 *
 *
 * @apiSuccessExample {json} Success-Response:
 *[
 *    {
 *        "week_day": 0,
 *        "supports_list": null
 *    },
 *    {
 *        "week_day": 1,
 *        "supports_list": [
 *            {
 *                "support_name": "Евгений Николаевич Табаков",
 *                "statuses": [
 *                    {
 *                        "status_name": "Принимаю запросы",
 *                        "duration": "7h22m41s"
 *                    },
 *                    {
 *                        "status_name": "Перерыв",
 *                        "duration": "2m25s"
 *                    },
 *                    {
 *                        "status_name": "Не принимаю запросы",
 *                        "duration": "5m23s"
 *                    }
 *                ]
 *            },
 *            {
 *                "support_name": "Вячеслав Викторович Тищенко",
 *                "statuses": [
 *                    {
 *                        "status_name": "Принимаю запросы",
 *                        "duration": "7h17m42s"
 *                    },
 *                    {
 *                        "status_name": "Работа в офисе",
 *                        "duration": "5m27s"
 *                    }
 *                ]
 *            }
 *        ]
 *    },
 *    {
 *        "week_day": 2,
 *        "supports_list": null
 *    },
 *    {
 *        "week_day": 3,
 *        "supports_list": null
 *    },
 *    {
 *        "week_day": 4,
 *        "supports_list": [
 *            {
 *                "support_name": "Евгений Николаевич Табаков",
 *                "statuses": [
 *                    {
 *                        "status_name": "Принимаю запросы",
 *                        "duration": "6h30m57s"
 *                    }
 *                ]
 *            },
 *            {
 *                "support_name": "Вячеслав Викторович Тищенко",
 *                "statuses": [
 *                    {
 *                        "status_name": "Принимаю запросы",
 *                        "duration": "64727h46m40s"
 *                    }
 *                ]
 *            }
 *        ]
 *    },
 *    {
 *        "week_day": 5,
 *        "supports_list": null
 *    },
 *    {
 *        "week_day": 6,
 *        "supports_list": null
 *    }
 *]
 *
 *
 */
