package auth

/**
*
* @api {POST} /helpdesk/auth/login 01. Авторизация
* @apiSampleRequest off
* @apiName SignIn
* @apiGroup 01. Авторизация
* @apiVersion  2.0.0
* @apiDescription Для авторизации необходимо передать адрес доменной электронной почты, а также пароль от нее.
* Если такой пользователь есть на LDAP сервере - авторизация пройдет без ошибок и в качестве ответа прийдет авторизационный токен для этого пользователя,
* а также информация по пользователю и по доступам которые есть у группы к которой он относится.
*
* @apiParam  {String} username Доменная электронная почта
* @apiParam  {String} password Пароль от доменной электронной почты
*
* @apiSuccess (Success 200) {String} 	token 								Авторизационный токен
* @apiSuccess (Success 200) {user} 		user 								Пользователь
* @apiSuccess (Success 200)	{String}	user.user_name						Имя пользователя
* @apiSuccess (Success 200) {String}	user.email							Электронный адрес пользователя
* @apiSuccess (Success 200) {group} 	user.group							Группа в которую пользователь входит
* @apiSuccess (Success 200) {Uint64}	user.group.group_id					ID группы
* @apiSuccess (Success 200)	{String}	user.group.group_name				Название группы
* @apiSuccess (Success 200)	{Bool}		user.group.create_ticket			Разрешение на создание запросов в тех. поддержку
* @apiSuccess (Success 200)	{Bool}		user.group.get_all_tickets			Разрешение на получение списка всех запросов
* @apiSuccess (Success 200) {Bool}		user.group.see_additional_info 		Разрешение на просмотр доп. информации в запросе
* @apiSuccess (Success 200) {Bool}		user.group.can_resolve_ticket 		Разрешение на согласование запросов
* @apiSuccess (Success 200) {Bool}		user.group.work_on_tickets 			Разрешение на работу с запросом как сотрудник тех. поддержки
* @apiSuccess (Success 200) {Bool}		user.group.change_settings 			Разрешение на изменение настроек системы
*
* @apiParamExample  {json} Request-Example:
*{
*   "username": "ivanov.i.i@limefin.com",
*   "password": "Qwerty123456"
*}
*
*
* @apiSuccessExample {json} Success-Response:
*{
*   "token": "<token>"
*    "user": {
*        "user_name": "Табаков Евгений Николаевич",
*        "email": "tabakov.e.n@limefin.com",
*        "group": {
*            "group_id": 3,
*            "group_name": "admin",
*            "create_ticket": true,
*            "get_all_tickets": true,
*            "see_additional_info": true,
*            "can_resolve_ticket": true,
*            "work_on_tickets": true,
*            "change_settings": true
*        }
*    }
*}
*
*
 */
