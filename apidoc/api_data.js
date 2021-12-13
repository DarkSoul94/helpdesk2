define({ "api": [
  {
    "type": "POST",
    "url": "/helpdesk/auth/login",
    "title": "01. Авторизация",
    "name": "SignIn",
    "group": "01._Авторизация",
    "version": "2.0.0",
    "description": "<p>Для авторизации необходимо передать адрес доменной электронной почты, а также пароль от нее. Если такой пользователь есть на LDAP сервере - авторизация пройдет без ошибок и в качестве ответа прийдет авторизационный токен для этого пользователя, а также информация по пользователю и по доступам которые есть у группы к которой он относится.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>Доменная электронная почта</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>Пароль от доменной электронной почты</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n  \"username\": \"ivanov.i.i@limefin.com\",\n  \"password\": \"Qwerty123456\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>Авторизационный токен</p>"
          },
          {
            "group": "Success 200",
            "type": "user",
            "optional": false,
            "field": "user",
            "description": "<p>Пользователь</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "user.user_name",
            "description": "<p>Имя пользователя</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "user.email",
            "description": "<p>Электронный адрес пользователя</p>"
          },
          {
            "group": "Success 200",
            "type": "group",
            "optional": false,
            "field": "user.group",
            "description": "<p>Группа в которую пользователь входит</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "user.group.group_id",
            "description": "<p>ID группы</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "user.group.group_name",
            "description": "<p>Название группы</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "user.group.create_ticket",
            "description": "<p>Разрешение на создание запросов в тех. поддержку</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "user.group.get_all_tickets",
            "description": "<p>Разрешение на получение списка всех запросов</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "user.group.see_additional_info",
            "description": "<p>Разрешение на просмотр доп. информации в запросе</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "user.group.can_resolve_ticket",
            "description": "<p>Разрешение на согласование запросов</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "user.group.work_on_tickets",
            "description": "<p>Разрешение на работу с запросом как сотрудник тех. поддержки</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "user.group.change_settings",
            "description": "<p>Разрешение на изменение настроек системы</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"token\": \"<token>\"\n   \"user\": {\n       \"user_name\": \"Табаков Евгений Николаевич\",\n       \"email\": \"tabakov.e.n@limefin.com\",\n       \"group\": {\n           \"group_id\": 3,\n           \"group_name\": \"admin\",\n           \"create_ticket\": true,\n           \"get_all_tickets\": true,\n           \"see_additional_info\": true,\n           \"can_resolve_ticket\": true,\n           \"work_on_tickets\": true,\n           \"change_settings\": true\n       }\n   }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/auth/01_login.go",
    "groupTitle": "01._Авторизация"
  },
  {
    "type": "POST",
    "url": "/helpdesk/category/create",
    "title": "01. Создание категории",
    "name": "CreateCategory",
    "group": "02._Категории_и_разделы_категорий",
    "version": "2.0.0",
    "description": "<p>Создание категории.</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "category_name",
            "description": "<p>Имя новой категории, при создании только раздела категории, можно опустить</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "significant",
            "description": "<p>Признак высокого приоритета у категории, при создании только раздела категории, можно опустить</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "old_category",
            "description": "<p>Признак того что категория устарела, при создании только раздела категории, можно опустить</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint",
            "optional": false,
            "field": "price",
            "description": "<p>Цена мотивации за запрос данной категории, при создании только раздела категории, можно опустить</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Запрос на создание категории:",
          "content": "{\n\t\"category_name\": 1C,\n\t\"significant\": false,\n\t\"old_category\": false,\n\t\"price\": 5\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "categoryid",
            "description": "<p>ID нового раздела категории</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус выполнения запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Ответ при создании категории:",
          "content": "{\n   \"category_id\": 14,\n   \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrBlankCategoryName",
            "description": "<p>Пустое имя категории</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrSectionAlreadyExist",
            "description": "<p>Категория с таким именем уже существует</p>"
          }
        ]
      }
    },
    "filename": "./docs/2.0.0/cat_end_sec/01_CreateCategory.go",
    "groupTitle": "02._Категории_и_разделы_категорий"
  },
  {
    "type": "POST",
    "url": "/helpdesk/section/create",
    "title": "04. Создание раздела категории",
    "name": "CreateSection",
    "group": "02._Категории_и_разделы_категорий",
    "version": "2.0.0",
    "description": "<p>Создание раздела категории.</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "section_id",
            "description": "<p>ID раздела категории. Для создания раздела должен равнятся 0</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "section_name",
            "description": "<p>Имя нового раздела категории. Уникально в рамках одной категории</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "significant",
            "description": "<p>Признак высокого приоритета у раздела</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "old_category",
            "description": "<p>Признак того что раздел категории устарел</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "need_approval",
            "description": "<p>Признак того что для раздела необходимо согласование</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "template",
            "description": "<p>Шаблон заполнения запроса</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "category_id",
            "description": "<p>ID категории</p>"
          },
          {
            "group": "Parameter",
            "type": "[]Uint64",
            "optional": false,
            "field": "approval_groups",
            "description": "<p>ID групп которые будут согласовывать запросы с этим разделом категорий</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Запрос на создание только раздела:",
          "content": "{\n   \"section_id\": 0,\n   \"section_name\": \"Перемещение менеджера\",\n   \"significant\": false,\n   \"old_category\": false,\n   \"need_approval\": false,\n\t \"template\":\"template for ticket\",\n\t \"category_id\": 2,\n\t \"approval_groups\": [1, 2]\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "category_section_id",
            "description": "<p>ID нового раздела категории</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Ответ при создании раздела:",
          "content": "{\n   \"category_section_id\": 14,\n   \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrBlankSectionName",
            "description": "<p>Пустое имя раздела категории</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrBlankCategoryName",
            "description": "<p>Пустое имя категории</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrSectionAlreadyExist",
            "description": "<p>Раздел категории с таким именем уже существует</p>"
          }
        ]
      }
    },
    "filename": "./docs/2.0.0/cat_end_sec/04_CreateCategorySection.go",
    "groupTitle": "02._Категории_и_разделы_категорий"
  },
  {
    "type": "GET",
    "url": "/helpdesk/category/",
    "title": "03. Получение списка категорий",
    "name": "GetCategory",
    "group": "02._Категории_и_разделы_категорий",
    "version": "2.0.0",
    "description": "<p>Получение списка категорий.</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Category",
            "optional": false,
            "field": "category",
            "description": "<p>Массив объектов &quot;категория запроса&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "category.category_id",
            "description": "<p>ID категории</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "category.category_name",
            "description": "<p>Имя категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "category.significant",
            "description": "<p>Признак высокого приоритета у категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "category.old_category",
            "description": "<p>Признак того что категория устарела</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "category.price",
            "description": "<p>Цена мотивации за запрос данной категории</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n   {\n       \"category_id\": 1,\n       \"category_name\": \"Оборудование\",\n       \"significant\": false,\n       \"old_category\": false,\n\t\t \"price\": 5\n   },\n   {\n       \"category_id\": 2,\n       \"category_name\": \"1C\",\n       \"significant\": false,\n       \"old_category\": false,\n\t\t \"price\": 1\n   }\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/cat_end_sec/03_GetCategory.go",
    "groupTitle": "02._Категории_и_разделы_категорий"
  },
  {
    "type": "GET",
    "url": "/helpdesk/section/",
    "title": "06. Получение списка разделов категорий без учета устаревших",
    "name": "GetCategorySection",
    "group": "02._Категории_и_разделы_категорий",
    "version": "2.0.0",
    "description": "<p>Получение списка разделов категорий. Разделы категорий которые помечены устаревшими или которые входят в устаревшую категорию не отображаются.</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Section[]",
            "optional": false,
            "field": "section",
            "description": "<p>Массив объектов &quot;раздел категории&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "section.section_id",
            "description": "<p>ID раздела категории</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "section.section_name",
            "description": "<p>Имя нового раздела категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "section.significant",
            "description": "<p>Высокий приоритет выполнения у раздела категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "section.old_category",
            "description": "<p>Устаревший раздел категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "section.need_approval",
            "description": "<p>Необходимость согласования</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "section.template",
            "description": "<p>Шаблон заполнения запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "category",
            "optional": false,
            "field": "section.category",
            "description": "<p>Категория запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "section.category.category_id",
            "description": "<p>ID категории</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "section.category.category_name",
            "description": "<p>Имя категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "section.category.significant",
            "description": "<p>Высокий приоритет выполнения у категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "section.category.old_category",
            "description": "<p>Устаревшая категория</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "section.category.price",
            "description": "<p>Цена мотивации за запрос данной категории</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n   {\n       \"section_id\": 1,\n       \"section_name\": \"Проблемы с кассовым аппаратом\",\n       \"significant\": false,\n       \"old_category\": false,\n       \"need_approval\": false,\n\t \t  \"template\":\"template for ticket\",\n       \"category\": {\n           \"category_id\": 1,\n           \"category_name\": \"Оборудование\",\n           \"significant\": false,\n           \"old_category\": false,\n\t\t \t \"price\": 5\n       }\n   },\n   {\n       \"section_id\": 2,\n       \"section_name\": \"Настройка нового оборудования\",\n       \"significant\": false,\n       \"old_category\": false,\n       \"need_approval\": false,\n\t \t  \"template\":\"template for ticket\",\n       \"category\": {\n           \"category_id\": 1,\n           \"category_name\": \"Оборудование\",\n           \"significant\": false,\n           \"old_category\": false,\n\t\t \t \"price\": 5\n       }\n   }\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/cat_end_sec/06_GetCategorySection.go",
    "groupTitle": "02._Категории_и_разделы_категорий"
  },
  {
    "type": "GET",
    "url": "/helpdesk/section/section_list",
    "title": "07. Получение всего списка разделов категорий",
    "name": "GetCategorySectionList",
    "group": "02._Категории_и_разделы_категорий",
    "version": "2.0.0",
    "description": "<p>Получение всего списка разделов категорий для проведения их настройки.</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Category[]",
            "optional": false,
            "field": "category",
            "description": "<p>Массив объектов &quot;категория&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "category.category_id",
            "description": "<p>ID категории</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "category.category_name",
            "description": "<p>Имя категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "category.significant",
            "description": "<p>Высокий приоритет выполнения у категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "category.old_category",
            "description": "<p>Устаревшая категория</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "category.price",
            "description": "<p>Цена мотивации за запрос данной категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Section[]",
            "optional": false,
            "field": "category.sections",
            "description": "<p>Массив объектов &quot;раздел категории&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "category.section.section_id",
            "description": "<p>ID раздела категории</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "category.section.section_name",
            "description": "<p>Имя нового раздела категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "category.section.significant",
            "description": "<p>Высокий приоритет выполнения у раздела категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "category.section.old_category",
            "description": "<p>Устаревший раздел категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "category.section.need_approval",
            "description": "<p>Необходимость согласования</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "category.section.category_id",
            "description": "<p>ID категории в которую входит раздел</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "category.section.template",
            "description": "<p>Шаблон заполнения запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n   {\n   \"category_id\": 4,\n   \"category_name\": \"Устаревшая\",\n   \"significant\": false,\n   \"old_category\": true,\n   \"price\": 0,\n   \"sections\": [\n     {\n       \"section_id\": 6,\n       \"section_name\": \"Уборка комнаты с игрушками\",\n       \"significant\": false,\n       \"old_category\": false,\n       \"need_approval\": false,\n       \"category_id\": 4,\n\t \t \"template\":\"template for ticket\",\n     }\n   ]\n },\n {\n   \"category_id\": 5,\n   \"category_name\": \"Валютообмен\",\n   \"significant\": false,\n   \"old_category\": false,\n   \"price\": 0,\n   \"sections\": [\n     {\n       \"section_id\": 7,\n       \"section_name\": \"Спецоперация\",\n       \"significant\": false,\n       \"old_category\": false,\n       \"need_approval\": false,\n       \"category_id\": 5,\n\t \t \"template\":\"template for ticket\",\n     }\n   ]\n }\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/cat_end_sec/07_GetCategorySectionList.go",
    "groupTitle": "02._Категории_и_разделы_категорий"
  },
  {
    "type": "POST",
    "url": "/helpdesk/category/update",
    "title": "02. Обновление категории",
    "name": "UpdateCategory",
    "group": "02._Категории_и_разделы_категорий",
    "version": "2.0.0",
    "description": "<p>Обновление категории.</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "category_id",
            "description": "<p>ID категории.</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "category_name",
            "description": "<p>Имя категории.</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "significant",
            "description": "<p>Признак высокого приоритета у категории.</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "old_category",
            "description": "<p>Признак того что категория устарела.</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint",
            "optional": false,
            "field": "price",
            "description": "<p>Цена мотивации за запрос данной категории.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t\"category_id\": 2,\n\t\"category_name\": \"1C\",\n\t\"significant\": false,\n\t\"old_category\": true,\n\t\"price\": 5\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус выполнения запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n   \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrBlankCategoryName",
            "description": "<p>Пустое имя категории</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrCategoryDoesNotExist",
            "description": "<p>Указанная категория не существует</p>"
          }
        ]
      }
    },
    "filename": "./docs/2.0.0/cat_end_sec/02_UpdateCategory.go",
    "groupTitle": "02._Категории_и_разделы_категорий"
  },
  {
    "type": "post",
    "url": "/helpdesk/section/update",
    "title": "05. Обновление разделов категории",
    "name": "UpdateCategorySection",
    "group": "02._Категории_и_разделы_категорий",
    "version": "2.0.0",
    "description": "<p>Обновление разделов категории.</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "section_id",
            "description": "<p>ID раздела категории</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "section_name",
            "description": "<p>Имя нового раздела категории</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "significant",
            "description": "<p>Признак высокого приоритета у раздела</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "old_category",
            "description": "<p>Признак того что раздел категории устарел</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "need_approval",
            "description": "<p>Признак того что для раздела необходимо согласование</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "template",
            "description": "<p>Шаблон заполнения запроса</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "category_id",
            "description": "<p>ID категории к которой принадлежит раздел</p>"
          },
          {
            "group": "Parameter",
            "type": "[]Uint64",
            "optional": false,
            "field": "approval_groups",
            "description": "<p>ID групп которые будут согласовывать запросы с этим разделом категорий</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n   \"section_id\": 1,\n   \"section_name\": \"Изменение по действующему займу\",\n   \"significant\": false,\n   \"old_category\": false,\n   \"need_approval\": false,\n\t \"template\":\"template for ticket\",\n   \"category_id\": 2,\n\t \"approval_groups\": [1, 2]\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус выполнения запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n   \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrBlankCategoryName",
            "description": "<p>Пустое имя категории</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrBlankSectionName",
            "description": "<p>Пустое имя раздела категории</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrSectionDoesNotExist",
            "description": "<p>Указанный раздел категории не существует</p>"
          }
        ]
      }
    },
    "filename": "./docs/2.0.0/cat_end_sec/05_UpdateCategorySection.go",
    "groupTitle": "02._Категории_и_разделы_категорий"
  },
  {
    "type": "POST",
    "url": "/helpdesk/service/auto_create",
    "title": "03. Автоматическое создание запроса",
    "name": "AutoCreateTicket",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "description": "<p>Автоматическое создание запроса из данных присланных по API сторонними сервисами.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "text",
            "description": "<p>Текст запроса</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "user_email",
            "description": "<p>Доменная почта пользователя</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "user_ip",
            "description": "<p>IP-адресс компьютера с которого отправлялся запрос</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "priority",
            "description": "<p>Признак являеться ли запрос приоритетным</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n   \"text\":\"api ticket\",\n   \"user_email\" : \"tishchenko.v.v@limefin.com\",\n   \"user_ip\":\"10.54.1.101\",\n\t  \"priority\" : true\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>статус выполнения запроса</p>"
          },
          {
            "group": "200",
            "type": "Int",
            "optional": false,
            "field": "ticket_id",
            "description": "<p>номер созданного запроса(0 - если была ошибка)</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Запрос успешно создан:",
          "content": "{\n  \"status\": \"ok\",\n  \"ticket_id\": 49\n}",
          "type": "json"
        },
        {
          "title": "Запрос без email:",
          "content": "{\n\t \"error\": \"Email is blank\",\n\t \"status\": \"error\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/tickets/03_AutoCreateTicket.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "GET",
    "url": "/helpdesk/resolve_ticket/check_exist",
    "title": "12. Проверка есть ли запросы ожидающие согласования",
    "name": "CheckNeedResolveTicketExist",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "exist",
            "description": "<p>Есть ли запросы в базе, если есть то <code>true</code></p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\t\t\"exist\": true,\n    \t\"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/tickets/12_CheckNeedResolveTicketExist.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "post",
    "url": "/helpdesk/comment/create",
    "title": "16. Создание нового комментария в запросе",
    "name": "CreateCommentsHistory",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_id",
            "description": "<p>ID запроса к которому относится комментарий</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "comment_text",
            "description": "<p>Текст комментария</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Запрос на создание комментария:",
          "content": "{\n  \"ticket_id\": 2,\n  \"comment_text\": \"комментарий к запросу\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "comment_id",
            "description": "<p>ID созданного комментария</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"comment_id\": 5,\n    \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrBlankComment",
            "description": "<p>Поле <code>comment_text</code> пустое</p>"
          }
        ]
      }
    },
    "filename": "./docs/2.0.0/tickets/comments/16_CreateCommentsHistory.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "post",
    "url": "/helpdesk/ticket/create",
    "title": "01. Создание запроса в тех. поддержку",
    "name": "CreateTicket",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "section_id",
            "description": "<p>ID раздела категории.</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ticket_text",
            "description": "<p>Текст запроса.</p>"
          },
          {
            "group": "Parameter",
            "type": "file[]",
            "optional": false,
            "field": "files",
            "description": "<p>Файлы.</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "files.file_name",
            "description": "<p>Имя файла.</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "files.file_data",
            "description": "<p>Данные файла.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n  \"section_id\": 1,\n  \"ticket_text\": \"Не вышел чек\",\n  \"files\": [\n      {\n          \"file_name\": \"скрин1.новый_клиент.jpg\",\n          \"file_data\": \"a few byte count\"\n      },\n      {\n          \"file_name\": \"скрин2.jpg\",\n          \"file_data\": \"a few byte count\"\n      }\n  ]\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_id",
            "description": "<p>ID созданого запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"status\": \"ok\",\n    \"ticket_id\": 8\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrInvalidID",
            "description": "<p>Неверный ID категории</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrBlankTicketText",
            "description": "<p>Пустой текст запроса</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrBlankCategoryName",
            "description": "<p>Пустое имя категории</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrBlankSectionName",
            "description": "<p>Пустое имя раздела категории</p>"
          }
        ]
      }
    },
    "filename": "./docs/2.0.0/tickets/01_CreateTicket.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "POST",
    "url": "/helpdesk/ticket/generate_tickets",
    "title": "02. Создание большого количества запросов",
    "name": "CreateTicket",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "text",
            "description": "<p>Текст запроса</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "section_id",
            "description": "<p>ИД раздела категории</p>"
          },
          {
            "group": "Parameter",
            "type": "User[]",
            "optional": false,
            "field": "users",
            "description": "<p>Массив обектов &quot;пользователь&quot;</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "users.user_id",
            "description": "<p>ИД пользователя которому создать запрос</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "users.count",
            "description": "<p>Количество запросов которые необходимо создать</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n  \"text\":\"тест много запросов3\",\n  \"section_id\":1,\n  \"users\":[\n      \t{\n          \"user_id\": 6,\n          \"count\": 1\n      \t}\n      ]\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Запросы успешно созданы:",
          "content": "{\n    \"status\": \"ok\"\n}",
          "type": "json"
        },
        {
          "title": "Пустой текст запроса:",
          "content": "{\n \"error\": \"Ticket text is blank\",\n \"status\": \"error\"\n}",
          "type": "json"
        },
        {
          "title": "ИД не существующей категории:",
          "content": "{\n \"error\": \"Such category section doesn't exist\",\n \"status\": \"error\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/tickets/02_GenerateTickets.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "POST",
    "url": "/helpdesk/ticket/ticket_grade",
    "title": "15. Оценка запроса в тех.поддержку",
    "name": "CreateTicketGrade",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_id",
            "description": "<p>ID оцениваемого запроса</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_grade",
            "description": "<p>Оценка</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n  \"ticket_id\":5,\n  \"ticket_grade\":5\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket_grade_id",
            "description": "<p>ID созданного объекта</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n   \"status\": \"ok\",\n   \"ticket_grade_id\": \"1\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/tickets/15_TicketGrade.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "get",
    "url": "/file/",
    "title": "17. Получение файла по его ид",
    "name": "GetFile",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "description": "<p>В запросе указывается ид файла который необходимо получить, в ответ возвращается файл в виде json объекта или ошибка.</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/file/?file_id=23",
        "type": "json"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "file_id",
            "description": "<p>ID получаемого файла</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "file_id",
            "description": "<p>ID файла</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "file_name",
            "description": "<p>Имя файла</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_id",
            "description": "<p>ID запроса к которому относится файл</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "file_data",
            "description": "<p>Данные файлы</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "file_date",
            "description": "<p>Дата добавления файла</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"file_id\": 48,\n  \"file_name\": \"скрин1.новый_клиент.jpg\",\n  \"file_data\": \"a few byte count\",\n  \"file_date\": \"2021-04-12T13:33:20Z\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/tickets/files/17_GetFile.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "GET",
    "url": "/helpdesk/resolve_ticket/resolve_tickets_list",
    "title": "13. Получение списка запросов в тех. поддержку ожидающих согласования",
    "name": "GetResolveTicketsList",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "count",
            "description": "<p>Количество запросов которые сервер должен вернуть</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "offset",
            "description": "<p>Сдвиг по списку запросов (общее колчество запросов которые клиент уже получил)</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/resolve_ticket/resolve_tickets_list?count=5&offset=0",
        "type": "json"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String[]",
            "optional": false,
            "field": "fields",
            "description": "<p>Список полей которые должны отрисоваться на стороне фронта</p>"
          },
          {
            "group": "Success 200",
            "type": "Ticket[]",
            "optional": false,
            "field": "tickets",
            "description": "<p>Массив запросов в тех. поддержку</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "tickets.ticket_id",
            "description": "<p>ID запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.ticket_date",
            "description": "<p>Дата создания запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.category",
            "description": "<p>Категория</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.section",
            "description": "<p>Раздел категории</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.ticket_text",
            "description": "<p>Текст запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.status",
            "description": "<p>Статус запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.ticket_author",
            "description": "<p>Автор запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.filial",
            "description": "<p>Отделение за которым закрелен автор запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"fields\": [\n        \t\"ticket_id\",\n        \t\"ticket_date\",\n        \t\"category\",\n        \t\"section\",\n        \t\"ticket_text\",\n        \t\"status\",\n\t\t\t\"filial\",\n        \t\"ticket_author\",\n    ],\n    \"tickets\": [\n       {\n           \"ticket_id\": 51,\n           \"ticket_date\": \"2021-05-26T11:40:41+03:00\",\n           \"category\": \"Оборудование\",\n           \"section\": \"Настройка интернета\",\n           \"ticket_text\": \"adasdasdads\",\n           \"status\": \"В ожидании согласования\",\n           \"filial\": \"not found\",\n           \"ticket_author\": \"Артем Владимирович Шелкопляс\"\n       },\n       {\n           \"ticket_id\": 49,\n           \"ticket_date\": \"2021-05-26T10:33:09+03:00\",\n           \"category\": \"1С\",\n           \"section\": \"Удаление кассовых\",\n           \"ticket_text\": \"gdfgdfbdfbdfb\",\n           \"status\": \"В ожидании согласования\",\n           \"filial\": \"not found\",\n           \"ticket_author\": \"Вячеслав Викторович Тищенко\"\n       },\n    ]\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/tickets/13_GetResolveTicketsList.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "GET",
    "url": "/helpdesk/ticket/ticket",
    "title": "06. Получение запроса",
    "name": "GetTicket",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "description": "<p>Метод который возвращает запрос по указаному ID с прикрепленными к нему коментариями и файлами. Если у пользователся нет доступа <code>see_additional_info</code> - у него не будут отображатся поля: <code>ticket_author</code>, <code>support</code>, <code>resolved_user</code>, <code>service_comment</code>.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "ticket_id",
            "description": "<p>Номер запрашиваемого запроса</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/ticket/ticket?ticket_id=1",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_id",
            "description": "<p>ID запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket_date",
            "description": "<p>Дата создания запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "CategorySection",
            "optional": false,
            "field": "category_section",
            "description": "<p>Раздел категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "category_section.section_id",
            "description": "<p>ID раздела категории</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "category_section.section_name",
            "description": "<p>Имя раздела категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "category_section.significant",
            "description": "<p>Приоритет выполнения</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "category_section.old_category",
            "description": "<p>Утаревший раздел категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "category_section.need_approval",
            "description": "<p>Необходимость согласования</p>"
          },
          {
            "group": "Success 200",
            "type": "Category",
            "optional": false,
            "field": "category_section.category",
            "description": "<p>Категория</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "category_section.category.category_id",
            "description": "<p>ID категории</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "category_section.category.category_name",
            "description": "<p>Имя категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "category_section.category.significant",
            "description": "<p>Приоритет выполнения</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "category_section.category.old_category",
            "description": "<p>Утаревшая категория</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "category_section.category.price",
            "description": "<p>Количество мотивации за запрос данной категории</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket_text",
            "description": "<p>Текст запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "TicketStatus",
            "optional": false,
            "field": "ticket_status",
            "description": "<p>Статус запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_status.ticket_status_id",
            "description": "<p>ID статуса запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket_status.ticket_status_name",
            "description": "<p>Имя статуса запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "filial",
            "description": "<p>Отделение за которым закрелен автор запроса. Если в базе филиал не найден вернет <code>&quot;not_found&quot;</code>. Если у пользователя нет прав на просмотр информации вернет пустое поле</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ip",
            "description": "<p>IP адрес автора запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "User",
            "optional": false,
            "field": "ticket_author",
            "description": "<p>Автор запроса. Если у пользователя нет прав на просмотр информации вернет <code>null</code></p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_author.user_id",
            "description": "<p>ID автора запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket_author.user_name",
            "description": "<p>Имя автора запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket_author.email",
            "description": "<p>Почта автора запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_author.group_id",
            "description": "<p>ID группы к которой принадлежит автор запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket_author.department",
            "description": "<p>Подразделение к которому относится автор запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "User",
            "optional": false,
            "field": "support",
            "description": "<p>Сотрудник тех. поддержки. Если у пользователя нет прав на просмотр информации вернет <code>null</code></p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "support.user_id",
            "description": "<p>ID сотрудника тех.поддержки</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "support.user_name",
            "description": "<p>Имя сотрудника тех.поддержки</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "support.email",
            "description": "<p>Почта сотрудника тех.поддержки</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "support.group_id",
            "description": "<p>ID группы к которой принадлежит пользователь</p>"
          },
          {
            "group": "Success 200",
            "type": "User",
            "optional": false,
            "field": "resolved_user",
            "description": "<p>Пользователь согласовавший запрос. Если у пользователя нет прав на просмотр информации вернет <code>null</code></p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "resolved_user.user_id",
            "description": "<p>ID пользователя согласовавшего запрос</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "resolved_user.user_name",
            "description": "<p>Имя пользователя согласовавшего запрос</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "resolved_user.email",
            "description": "<p>Почта пользователя согласовавшего запрос</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "resolved_user.group_id",
            "description": "<p>ID группы к которой принадлежит пользователь</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "service_comment",
            "description": "<p>Сервисный комментарий для сотрудников тех. поддержки</p>"
          },
          {
            "group": "Success 200",
            "type": "comments[]",
            "optional": false,
            "field": "comments",
            "description": "<p>Комментарии</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "comments.comment_id",
            "description": "<p>ID комментария</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "comments.comment_date",
            "description": "<p>Дата добавления комментария</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "comment_author",
            "description": "<p>Автор комментария</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "comments.comment_text",
            "description": "<p>Текст комментария</p>"
          },
          {
            "group": "Success 200",
            "type": "Files[]",
            "optional": false,
            "field": "files",
            "description": "<p>Файлы</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "files.file_id",
            "description": "<p>ID файла</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "files.file_name",
            "description": "<p>Имя файла</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "files.file_date",
            "description": "<p>Дата добавления файла</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Вид запроса для админа и сотрудника ТП:",
          "content": "{\n \"ticket_id\": 2,\n \"ticket_date\": \"2021-05-18T16:49:30+03:00\",\n \"category_section\": {\n   \"section_id\": 2,\n   \"section_name\": \"Удаление кассовых\",\n   \"significant\": false,\n   \"old_category\": false,\n   \"need_approval\": true,\n   \"category\": {\n     \"category_id\": 2,\n     \"category_name\": \"1С\"\n   }\n },\n \"ticket_text\": \"delete\",\n \"ticket_status\": {\n   \"ticket_status_id\": 9,\n   \"ticket_status_name\": \"Выполнен\"\n },\n \"filial\": \"not found\",\n  \"ip\": \"10.54.86.26\",\n \"ticket_author\": {\n   \"user_id\": 5,\n   \"user_name\": \"Владислав Сергеевич Маспанов\",\n   \"email\": \"maspanov.v.s@limefin.com\",\n   \"group_id\": 3,\n   \"department\": \"Техническая поддержка\"\n },\n \"support\": {\n   \"user_id\": 6,\n   \"user_name\": \"Вячеслав Викторович Тищенко\",\n   \"email\": \"tishchenko.v.v@limefin.com\",\n   \"group_id\": 2\n },\n \"resolved_user\": {\n   \"user_id\": 6,\n   \"user_name\": \"Вячеслав Викторович Тищенко\",\n   \"email\": \"tishchenko.v.v@limefin.com\",\n   \"group_id\": 2\n },\n \"service_comment\": \"\",\n \"comments\": [],\n \"files\": []\n}\n*",
          "type": "json"
        },
        {
          "title": "Вид запроса для остальных пользователей:",
          "content": "{\n\"ticket_id\": 2,\n\"ticket_date\": \"2021-05-18T16:49:30+03:00\",\n\"category_section\": {\n\t \"section_id\": 2,\n   \"section_name\": \"Удаление кассовых\",\n   \"significant\": false,\n   \"old_category\": false,\n   \"need_approval\": true,\n   \"category\": {\n     \"category_id\": 2,\n     \"category_name\": \"1С\"\n   }\n },\n \"ticket_text\": \"delete\",\n \"ticket_status\": {\n   \"ticket_status_id\": 9,\n   \"ticket_status_name\": \"Выполнен\"\n },\n \"filial\": \" \",\n \"ip\": \" \",\n \"ticket_author\": null,\n \"support\": null,\n \"resolved_user\": null,\n \"service_comment\": \"\",\n \"comments\": [],\n \"files\": []\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/tickets/06_GetTicket.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "get",
    "url": "/helpdesk/ticket_status/",
    "title": "08. Получение списка статусов запроса в ТП для выбора внутри запроса",
    "name": "GetTicketStatus",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "description": "<p>Возвращает массив объектов</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "TicketStatus[]",
            "optional": false,
            "field": "ticket_status",
            "description": "<p>Массив объеквтов &quot;статус запроса&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_status.ticket_status_id",
            "description": "<p>ID статуса запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket_status.ticket_status_name",
            "description": "<p>Имя статуса запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n  {\n      \"ticket_status_id\": 1,\n      \"ticket_status_name\": \"Новый\"\n  },\n  {\n      \"ticket_status_id\": 2,\n      \"ticket_status_name\": \"В ожидании\"\n  }\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/tickets/08_GetTicketStatusListAll.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "get",
    "url": "/helpdesk/ticket_status/history",
    "title": "10. Получение истории изменения статусов запроса",
    "name": "GetTicketStatusHistory",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_id",
            "description": "<p>ID запроса по которому нужно посмотреть историю статусов</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "StatusHistory[]",
            "optional": false,
            "field": "status_history",
            "description": "<p>Массив объектов &quot;история изменения статуса&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status_history.curr_status_time",
            "description": "<p>Время присвоения текущего статуса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status_history.curr_status",
            "description": "<p>Название текущего статуса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status_history.changed_user",
            "description": "<p>Пользователь сменивший статус</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "status_history.difference",
            "description": "<p>Время нахождения в статусе</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n  {\n    \"curr_status_time\": \"2021-11-24T12:43:53Z\",\n    \"curr_status\": \"В ожидании\",\n    \"changed_user\": \"Вячеслав Викторович Тищенко\",\n    \"difference\": 1784\n  }\n]",
          "type": "json"
        }
      ]
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/ticket_status/history?ticket_id=2",
        "type": "json"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrStatusHistoryNotExist",
            "description": "<p>В базе нет записей истории изменения статусов по данному запросу</p>"
          }
        ]
      }
    },
    "filename": "./docs/2.0.0/tickets/10_GetTicketStatusHistory.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "get",
    "url": "/helpdesk/ticket_status/list",
    "title": "09. Получение всего списка статусов у запроса",
    "name": "GetTicketStatusList",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "description": "<p>Возвращает массив объектов</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "TicketStatus[]",
            "optional": false,
            "field": "ticket_status",
            "description": "<p>Массив объеквтов &quot;статус запроса&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_status.ticket_status_id",
            "description": "<p>ID статуса запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket_status.ticket_status_name",
            "description": "<p>Имя статуса запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n  {\n      \"ticket_status_id\": 1,\n      \"ticket_status_name\": \"Новый\"\n  },\n  {\n      \"ticket_status_id\": 2,\n      \"ticket_status_name\": \"В ожидании\"\n  }\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/tickets/09_GetTicketStatusList.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "GET",
    "url": "/helpdesk/ticket/tickets_list",
    "title": "04. Получение списка запросов в тех. поддержку",
    "name": "GetTicketsList",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "description": "<p>В зависимости от прав закрепленных за группой пользователей будет отличатся результирующий список запросов. В случае если у пользователя есть разрешение <code>get_all_tickets</code> - получит весь список запросов, при <code>can_resolve_ticket</code> - получит написанные собой запросы, а также запросы на согласование, при <code>work_on_tickets</code> - получит распределенные на него запросы. В случае если нет ни одного из этих разрешений пользователь получит список только написанных им запросов. Поле <code>support</code> отображается только для пользователя с правами <code>admin</code>.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "count",
            "description": "<p>Количество запросов которые сервер должен вернуть</p>"
          },
          {
            "group": "Parameter",
            "type": "Int",
            "optional": false,
            "field": "offset",
            "description": "<p>Сдвиг по списку запросов (общее колчество запросов которые клиент уже получил)</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/ticket/tickets_list?count=5&offset=0",
        "type": "json"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String[]",
            "optional": false,
            "field": "fields",
            "description": "<p>Список полей которые должны отрисоваться на стороне фронта</p>"
          },
          {
            "group": "Success 200",
            "type": "Ticket[]",
            "optional": false,
            "field": "tickets",
            "description": "<p>Массив запросов в тех. поддержку</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "tickets.ticket_id",
            "description": "<p>ID запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.ticket_date",
            "description": "<p>Дата создания запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "tickets.significant",
            "description": "<p>Важность запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.category",
            "description": "<p>Категория</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.section",
            "description": "<p>Раздел категории</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.ticket_text",
            "description": "<p>Текст запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.status",
            "description": "<p>Статус запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.ticket_author",
            "description": "<p>Автор запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.support",
            "description": "<p>Сотрудник ТП работающий над запросом</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.filial",
            "description": "<p>Отделение за которым закрелен автор запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "tickets.grade",
            "description": "<p>Оценка запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "tickets.sort_priority",
            "description": "<p>Приоритет отображения запроса (1 - наивысший приоритет). В рамках одного приоритета запросы сортируются по своим ИД по убыванию.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Ответ при запросе списка админом:",
          "content": "{\n    \"fields\": [\n\t\t\t\"ticket_id\",\n\t\t\t\"ticket_date\",\n\t\t\t\"category\",\n\t\t\t\"section\",\n\t\t\t\"ticket_text\",\n\t\t\t\"status\",\n\t\t\t\"filial\",\n\t\t\t\"ticket_author\",\n\t\t\t\"support\",\n\t\t\t\"grade\"\n    ],\n    \"tickets\": [\n\t\t{\n\t\t\t\"ticket_id\": 68,\n\t\t\t\"ticket_date\": \"2021-07-28T06:55:17Z\",\n\t\t\t\"significant\": false,\n\t\t\t\"category\": \"Оборудование\",\n\t\t\t\"section\": \"Проблема с принтером/печатью\",\n\t\t\t\"ticket_text\": \"afasdasdqds\",\n\t\t\t\"status\": \"Отклонен\",\n\t\t\t\"ticket_status_id\": 8,\n\t\t\t\"filial\": \"not found\",\n\t\t\t\"ticket_author\": \"Евгений Николаевич Табаков\",\n\t\t\t\"support\": \"Вячеслав Викторович Тищенко\",\n\t\t\t\"grade\": 0,\n\t\t\t\"sort_priority\": 1\n\t\t}\n    ]\n}",
          "type": "json"
        },
        {
          "title": "Ответ при запросе списка остальными пользователями:",
          "content": "{\n    \"fields\": [\n\t\t\t\"ticket_id\",\n\t\t\t\"ticket_date\",\n\t\t\t\"category\",\n\t\t\t\"section\",\n\t\t\t\"ticket_text\",\n\t\t\t\"status\",\n\t\t\t\"filial\",\n\t\t\t\"ticket_author\",\n\t\t\t\"grade\"\n    ],\n    \"tickets\": [\n        {\n\t\t\t\"ticket_id\": 68,\n\t\t\t\"ticket_date\": \"2021-07-28T06:55:17Z\",\n\t\t\t\"significant\": false,\n\t\t\t\"category\": \"Оборудование\",\n\t\t\t\"section\": \"Проблема с принтером/печатью\",\n\t\t\t\"ticket_text\": \"afasdasdqds\",\n\t\t\t\"status\": \"Отклонен\",\n\t\t\t\"ticket_status_id\": 8,\n\t\t\t\"filial\": \"not found\",\n\t\t\t\"ticket_author\": \"Евгений Николаевич Табаков\",\n\t\t\t\"support\": \"\",\n\t\t\t\"grade\": 0,\n\t\t\t\"sort_priority\": 4\n        }\n    ]\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/tickets/04_GetTicketsList.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "POST",
    "url": "/helpdesk/resolve_ticket/resolve",
    "title": "14. Согласование запроса",
    "name": "ResolveTicket",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_id",
            "description": "<p>ID запроса который согласовывается.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n  \"ticket_id\":5\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrTicketDoesNotExist",
            "description": "<p>Запроса с таким ID не существует</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrDoesNotNeedApproval",
            "description": "<p>Запрос не нуждается в согласовании</p>"
          }
        ]
      }
    },
    "filename": "./docs/2.0.0/tickets/14_ResolveTicket.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "POST",
    "url": "/helpdesk/ticket/steal",
    "title": "11. Взять чужой запрос себе в работу.",
    "name": "StealTicket",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_id",
            "description": "<p>ID запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n  \"ticket_id\":5\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"status\": \"ok\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "{\n     \"status\": \"error\",\n\t\t\"error\":\"Ticket is complete\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrTicketIsComplete",
            "description": "<p>Ticket is complete</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrTicketDoesNotExist",
            "description": "<p>Ticket with this id not exist</p>"
          }
        ]
      }
    },
    "filename": "./docs/2.0.0/tickets/11_StealTicket.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "post",
    "url": "/helpdesk/ticket/update",
    "title": "07. Обновление запроса в ТП",
    "name": "UpdateTicket",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "description": "<p>Обновление запроса в тех.поддержку.</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_id",
            "description": "<p>ID запроса</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "section_id",
            "description": "<p>ID раздела категории. Передается только при изменении раздела категории в запросе</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_status_id",
            "description": "<p>ID статуса запроса. Передается только при изменении статуса запроса</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "support_id",
            "description": "<p>ID сотрудника ТП, передается только при изменении сотрудника ТП в запросе</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "service_comment",
            "description": "<p>Сервисный комментарий для сотрудников ТП, передается только при добавлении/изменении сервисного комментария</p>"
          },
          {
            "group": "Parameter",
            "type": "files[]",
            "optional": false,
            "field": "files",
            "description": "<p>Файлы, передается только при добавлении/изменении файлов</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "files.file_name",
            "description": "<p>Имя файла</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "files.file_data",
            "description": "<p>Данные файла</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"ticket_id\": 2,\n    \"section_id\": 1,\n    \"ticket_status_id\": 9,\n    \"service_comment\": \"test\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrInvalidID",
            "description": "<p>Неверный ID категории</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrBlankTicketText",
            "description": "<p>Пустой текст запроса</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrBlankCategoryName",
            "description": "<p>Пустое имя категории</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrBlankSectionName",
            "description": "<p>Пустое имя раздела категории</p>"
          }
        ]
      }
    },
    "filename": "./docs/2.0.0/tickets/07_UpdateTicket.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "POST",
    "url": "/helpdesk/ticket/filtered_tickets_list",
    "title": "05. Получение отфильтрованого списка запросов",
    "name": "apiName",
    "group": "03._Запросы_в_тех._поддержку",
    "version": "2.0.0",
    "description": "<p>Если у обратившегося пользователя стоит доступ <code>full_search</code> фильтр накладывается на все запросы. Если стоит <code>can_resolve</code> накладывается на список запросов которые требовали и требуют согласования и там где пользователь автор запроса. Для всех остальных накладывается на список где пользователь автор запроса. Поле support_id игнорируется для всех у кого не стоит доступ <code>full_search</code>. Если в фильтре ничего не указано отправится список всех запросов с учётом доступов пользователя.</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "ticket_id",
            "description": "<p>ИД запроса</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "start_date",
            "description": "<p>Дата начала выборки</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "end_date",
            "description": "<p>Дата конца выборки</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "category_id",
            "description": "<p>ИД категории</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64[]",
            "optional": false,
            "field": "section_id",
            "description": "<p>Массив ид раздела категории</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "text",
            "description": "<p>Текст который должен содержаться в запросе</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "status_id",
            "description": "<p>ИД статуса запроса</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64[]",
            "optional": false,
            "field": "author_id",
            "description": "<p>Массив ид автора запроса</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64[]",
            "optional": false,
            "field": "support_id",
            "description": "<p>Массив ид сотрудника работавшего над запросом</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "filial",
            "description": "<p>Филиал запроса</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "comment",
            "description": "<p>Текст который должен содержаться в комментариях к этому запросу</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"ticket_id\": 131,\n    \"start_date\" : \"2021-06-17 10:00:00\",\n    \"end_date\" : \"2021-06-17 10:30:00\",\n    \"category_id\" : 1,\n    \"section_id\" : [1, 2, 3],\n    \"text\" : \"asdfasf\",\n    \"status_id\" : 1,\n    \"author_id\" : [1, 2, 3],\n    \"support_id\" : [1, 2, 3],\n    \"filial\" : \"asfasdcv\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String[]",
            "optional": false,
            "field": "fields",
            "description": "<p>Список полей которые должны отрисоваться на стороне фронта</p>"
          },
          {
            "group": "Success 200",
            "type": "Ticket[]",
            "optional": false,
            "field": "tickets",
            "description": "<p>Массив запросов в тех. поддержку</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "tickets.ticket_id",
            "description": "<p>ID запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.ticket_date",
            "description": "<p>Дата создания запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.category",
            "description": "<p>Категория</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.section",
            "description": "<p>Раздел категории</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.ticket_text",
            "description": "<p>Текст запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.status",
            "description": "<p>Статус запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.ticket_author",
            "description": "<p>Автор запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.support",
            "description": "<p>Сотрудник ТП работающий над запросом</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tickets.filial",
            "description": "<p>Отделение за которым закрелен автор запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "tickets.grade",
            "description": "<p>Оценка запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Ответ для обычного пользователя и сотрудника бэк-офиса:",
          "content": "{\n  \"fields\": [\n    \"ticket_id\",\n    \"ticket_date\",\n    \"category\",\n    \"section\",\n    \"ticket_text\",\n    \"status\",\n    \"filial\",\n    \"ticket_author\",\n    \"grade\"\n  ],\n  \"tickets\": [\n    {\n      \"ticket_id\": 141,\n      \"ticket_date\": \"2021-06-17T14:42:30Z\",\n      \"significant\": false,\n      \"category\": \"Валютообмен\",\n      \"section\": \"Спецоперация\",\n      \"ticket_text\": \"asfsdfgsgvxcfasf\",\n      \"status\": \"В процессе реализации\",\n      \"ticket_status_id\": 5,\n      \"filial\": \"not found\",\n      \"ticket_author\": \"Вячеслав Викторович Тищенко\",\n      \"support\": \"\",\n      \"grade\": 0\n    }\n ]\n}",
          "type": "json"
        },
        {
          "title": "Ответ для тех у кого есть разрешение `full_search`:",
          "content": "{\n  \"fields\": [\n   \"ticket_id\",\n   \"ticket_date\",\n   \"category\",\n   \"section\",\n   \"ticket_text\",\n   \"status\",\n   \"filial\",\n   \"ticket_author\",\n   \"support\",\n   \"grade\"\n  ],\n  \"tickets\": [\n    {\n      \"ticket_id\": 140,\n      \"ticket_date\": \"2021-06-17T13:20:57Z\",\n      \"significant\": true,\n      \"category\": \"1С \",\n      \"section\": \"Изменение/удаление кассовых ордеров\",\n      \"ticket_text\": \"выпывп\",\n      \"status\": \"В процессе реализации\",\n      \"ticket_status_id\": 5,\n      \"filial\": \"not found\",\n      \"ticket_author\": \"Евгений Николаевич Табаков\",\n      \"support\": \"Евгений Николаевич Табаков\",\n      \"grade\": 0\n    }\n ]\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/tickets/05_GetFilteredTicketList.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "POST",
    "url": "/helpdesk/group/create",
    "title": "03. Создание группы прав пользователей",
    "name": "CreateGroup",
    "group": "04._Пользователи",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "group_name",
            "description": "<p>Имя группы</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "create_ticket",
            "description": "<p>Разрешение принимать запросы</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "get_all_tickets",
            "description": "<p>Разрешение получать список запросов</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "see_additional_info",
            "description": "<p>Разрешение смотреть расширенную информацию по запросу</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "can_resolve_ticket",
            "description": "<p>Разрешение согласовывать запросы</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "work_on_tickets",
            "description": "<p>Разрешение работать над запросом</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "change_settings",
            "description": "<p>Разрешение изменять данные в базе</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n  \"group_name\": \"support\",\n  \"create_ticket\": false,\n  \"get_all_tickets\": false,\n  \"see_additional_info\": false,\n  \"can_resolve_ticket\": false,\n  \"work_on_tickets\": false,\n  \"change_settings\": false\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "group_id",
            "description": "<p>ID созданной группы</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"group_id\": 3,\n    \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrGroupAlreadyExist",
            "description": "<p>Такая группа уже существует</p>"
          }
        ]
      }
    },
    "filename": "./docs/2.0.0/users_groups/03_CreateGroup.go",
    "groupTitle": "04._Пользователи"
  },
  {
    "type": "GET",
    "url": "/helpdesk/group/",
    "title": "04. Получение списка групп прав",
    "name": "GetAllGroup",
    "group": "04._Пользователи",
    "version": "2.0.0",
    "description": "<p>Возвращает массив объектов</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Group[]",
            "optional": false,
            "field": "group",
            "description": "<p>Массив объектов &quot;группа пользователя&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "group.group_id",
            "description": "<p>ID группы</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "group.group_name",
            "description": "<p>Имя группы</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "group.create_ticket",
            "description": "<p>Разрешение принимать запросы</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "group.get_all_tickets",
            "description": "<p>Разрешение получать список запросов</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "group.see_additional_info",
            "description": "<p>Разрешение смотреть расширенную информацию по запросу</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "group.can_resolve_ticket",
            "description": "<p>Разрешение согласовывать запросы</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "group.work_on_tickets",
            "description": "<p>Разрешение работать над запросом</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "group.change_settings",
            "description": "<p>Разрешение изменять данные в базе</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n  {\n      \"group_id\": 1,\n      \"group_name\": \"regular_user\",\n      \"create_ticket\": true,\n      \"get_all_tickets\": false,\n      \"see_additional_info\": false,\n      \"can_resolve_ticket\": false,\n      \"work_on_tickets\": false,\n      \"change_settings\": false\n  },\n  {\n      \"group_id\": 2,\n      \"group_name\": \"support\",\n      \"create_ticket\": false,\n      \"get_all_tickets\": false,\n      \"see_additional_info\": false,\n      \"can_resolve_ticket\": false,\n      \"work_on_tickets\": false,\n      \"change_settings\": false\n  }\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/users_groups/04_GetAllGroup.go",
    "groupTitle": "04._Пользователи"
  },
  {
    "type": "GET",
    "url": "/helpdesk/user/",
    "title": "01. Получение списка пользователей",
    "name": "GetAllUsers",
    "group": "04._Пользователи",
    "version": "2.0.0",
    "description": "<p>Возвращает массив объектов</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "User[]",
            "optional": false,
            "field": "user",
            "description": "<p>Массив объектов пользователь</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "user.user_id",
            "description": "<p>ID пользователя</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "user.user_name",
            "description": "<p>Имя пользователя</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "user.email",
            "description": "<p>Электронная почта пользователя</p>"
          },
          {
            "group": "Success 200",
            "type": "Group",
            "optional": false,
            "field": "user.group",
            "description": "<p>Группа в которой состоит пользователь</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "user.group.group_id",
            "description": "<p>ID группы</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "user.group.group_name",
            "description": "<p>Имя группы</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n  {\n      \"user_id\": 1,\n      \"user_name\": \"Евгений Николаевич Табаков\",\n      \"email\": \"tabakov.e.n@limefin.com\",\n      \"group\": {\n          \"group_id\": 2,\n          \"group_name\": \"admin\"\n      }\n  },\n  {\n      \"user_id\": 2,\n      \"user_name\": \"Вячеслав Викторович Тищенко\",\n      \"email\": \"tishchenko.v.v@limefin.com\",\n      \"group\": {\n          \"group_id\": 2,\n          \"group_name\": \"admin\"\n      }\n  }\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/users_groups/01_GetAllUsers.go",
    "groupTitle": "04._Пользователи"
  },
  {
    "type": "GET",
    "url": "/helpdesk/user/departments_list",
    "title": "07. Получение списка отделов сотрудников",
    "name": "GetDepartmentsList",
    "group": "04._Пользователи",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "Departmen[]",
            "optional": false,
            "field": "departments",
            "description": "<p>Список отделов</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"departments\": [\n    \"Техническая поддержка\",\n    \"Разработчики\"\n  ],\n  \"status\": \"ok\"\n}",
          "type": "type"
        }
      ]
    },
    "filename": "./docs/2.0.0/users_groups/07_GetDepartmentsList.go",
    "groupTitle": "04._Пользователи"
  },
  {
    "type": "GET",
    "url": "/helpdesk/group/for_resolve",
    "title": "06. Получение списка груп пользователей с правами согласовывать запросы",
    "name": "GetGroupsListForResolve",
    "group": "04._Пользователи",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Groups",
            "optional": false,
            "field": "groups",
            "description": "<p>Массив объектов групп</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "groups.group_id",
            "description": "<p>ID группы</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "groups.group_name",
            "description": "<p>Название группы</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"groups\": [\n    {\n      \"group_id\": 4,\n      \"group_name\": \"Сотрудник бэк-офиса\"\n    }\n  ],\n  \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/users_groups/06_GetGroupsListForResolve.go",
    "groupTitle": "04._Пользователи",
    "sampleRequest": [
      {
        "url": "http://localhost:8888//helpdesk/group/for_resolve"
      }
    ]
  },
  {
    "type": "POST",
    "url": "/helpdesk/group/update",
    "title": "05. Обновление данных группы прав",
    "name": "UpdateGroup",
    "group": "04._Пользователи",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "group_id",
            "description": "<p>ID группы</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "group_name",
            "description": "<p>Имя группы</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "create_ticket",
            "description": "<p>Разрешение принимать запросы</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "get_all_tickets",
            "description": "<p>Разрешение получать список запросов</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "see_additional_info",
            "description": "<p>Разрешение смотреть расширенную информацию по запросу</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "can_resolve_ticket",
            "description": "<p>Разрешение согласовывать запросы</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "work_on_tickets",
            "description": "<p>Разрешение работать над запросом</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "change_settings",
            "description": "<p>Разрешение изменять данные в базе</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n  \"group_id\": 2,\n  \"group_name\": \"support\",\n  \"create_ticket\": false,\n  \"get_all_tickets\": false,\n  \"see_additional_info\": false,\n  \"can_resolve_ticket\": false,\n  \"work_on_tickets\": false,\n  \"change_settings\": false\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "GroupErr_NotExist",
            "description": "<p>Такой группы не существует</p>"
          }
        ]
      }
    },
    "filename": "./docs/2.0.0/users_groups/05_UpdateGroup.go",
    "groupTitle": "04._Пользователи"
  },
  {
    "type": "POST",
    "url": "/helpdesk/user/update",
    "title": "02. Обновление данных пользователя",
    "name": "UpdateUser",
    "group": "04._Пользователи",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "user_id",
            "description": "<p>ID пользователя1</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "group_id",
            "description": "<p>ID группы</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n  \"user_id\": 2,\n  \"group_id\": 1\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \t\"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/users_groups/02_UpdateUser.go",
    "groupTitle": "04._Пользователи"
  },
  {
    "type": "POST",
    "url": "/helpdesk/support/change_status",
    "title": "04. Сменить статус сотрудника ТП",
    "name": "ChangeSupportStatus",
    "group": "05._Сотрудник_ТП",
    "version": "2.0.0",
    "description": "<p>Смена рабочего статуса сотрудника ТП: &quot;принимаю запросы&quot;, &quot;не принимаю запросы&quot; и т.д. Используется как для смены статуса самому себе, так и для смены статуса другому сотруднику. Сменить статус другому сотруднику может только админ.</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "support_id",
            "description": "<p>ID сотрудника ТП. Если ID сотрудника не передается, то ID берется из авторизационного токена</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "support_status_id",
            "description": "<p>ID нового статуса сотрудника ТП</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Смена статуса другому суппорту:",
          "content": "{\n\t\t\"support_id\": 4,\n\t\t\"support_status_id\": 1\n}",
          "type": "json"
        },
        {
          "title": "Смена статуса себе:",
          "content": "{\n        \"support_status_id\": 1\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/support/supports/04_ChangeSupportStatus.go",
    "groupTitle": "05._Сотрудник_ТП"
  },
  {
    "type": "POST",
    "url": "/helpdesk/support/close_shift",
    "title": "09. Закрытие смены сотруднику ТП",
    "name": "CloseShift",
    "group": "05._Сотрудник_ТП",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "support_id",
            "description": "<p>ID сотрудника которому нужно закрыть смену</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"support_id\" : 6\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/support/supports/09_ShiftClose.go",
    "groupTitle": "05._Сотрудник_ТП"
  },
  {
    "type": "POST",
    "url": "/helpdesk/support/create_lateness",
    "title": "08. Отправка причины опоздания",
    "name": "CreateLateness",
    "group": "05._Сотрудник_ТП",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "support_id",
            "description": "<p>ID сотрудника которому нужно открыть смену</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "cause",
            "description": "<p>Причина опоздания</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"support_id\" : 4,\n    \"cause\": \"test\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/support/supports/08_CreateLateness.go",
    "groupTitle": "05._Сотрудник_ТП"
  },
  {
    "type": "GET",
    "url": "/helpdesk/support/active_support_list",
    "title": "06. Получение списка активных сотрудников ТП",
    "name": "GetActiveSupportList",
    "group": "05._Сотрудник_ТП",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "ActiveSupport[]",
            "optional": false,
            "field": "active_support",
            "description": "<p>Массив объектов &quot;Активный сотрудник ТП&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "active_support.user_id",
            "description": "<p>ID сотрудника</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "active_support.user_name",
            "description": "<p>ФИО сотрудника</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n    {\n        \"user_id\": 5,\n        \"user_name\": \"Евгений Николаевич Табаков\"\n    },\n    {\n        \"user_id\": 4,\n        \"user_name\": \"Вячеслав Викторович Тищенко\"\n    }\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/support/supports/06_GetActiveSupportList.go",
    "groupTitle": "05._Сотрудник_ТП"
  },
  {
    "type": "GET",
    "url": "/support/card/list",
    "title": "11. Получить список карточек суппорта",
    "name": "GetCardList",
    "group": "05._Сотрудник_ТП",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "[]Card",
            "optional": false,
            "field": "cards",
            "description": "<p>Массив карточек суппорта</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "cards.id",
            "description": "<p>ИД карточки супорта</p>"
          },
          {
            "group": "Success 200",
            "type": "CardUser",
            "optional": false,
            "field": "cards.support",
            "description": "<p>Объект содержащий данные суппорта</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "cards.support.id",
            "description": "<p>ИД суппорта</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "cards.support.name",
            "description": "<p>ФИО суппорта</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "cards.is_senior",
            "description": "<p>Признак старшего группы саппортов</p>"
          },
          {
            "group": "Success 200",
            "type": "CardUser",
            "optional": false,
            "field": "cards.senior",
            "description": "<p>Объект содержащий данные о старшем суппорте в группу которого входит текущий сотрудник</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "cards.senior.id",
            "description": "<p>ИД суппорта</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "cards.senior.name",
            "description": "<p>ФИО суппорта</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "cards.color",
            "description": "<p>Цвет отображения (в шестнадцатиричной системе)</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Список карточек:",
          "content": "[\n  {\n    \"id\": 13,\n    \"support\": {\n      \"id\": 4,\n      \"name\": \"Вячеслав Викторович Тищенко\"\n    },\n    \"is_senior\": true,\n    \"senior\": null,\n    \"color\": \"0xFFFFF0\"\n  },\n  {\n    \"id\": 14,\n    \"support\": {\n      \"id\": 5,\n      \"name\": \"Евгений Николаевич Табаков\"\n    },\n    \"is_senior\": false,\n    \"senior\": {\n      \"id\": 4,\n      \"name\": \"Вячеслав Викторович Тищенко\"\n    },\n    \"color\": \"0xFFFFF0\"\n  }\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/support/cards/11_GetCardList.go",
    "groupTitle": "05._Сотрудник_ТП"
  },
  {
    "type": "GET",
    "url": "/support/card/seniors",
    "title": "13. Получить список старших суппортов",
    "name": "GetSeniorSupportsList",
    "group": "05._Сотрудник_ТП",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "[]SeniorSupport",
            "optional": false,
            "field": "seniors",
            "description": "<p>Массив старших суппортов</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "seniors.id",
            "description": "<p>ИД старшего суппорта</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "seniors.name",
            "description": "<p>ФИО суппорта</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Список старших суппортов:",
          "content": "[\n  {\n    \"id\": 4,\n    \"name\": \"Вячеслав Викторович Тищенко\"\n  }\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/support/cards/13_GetSeniorList.go",
    "groupTitle": "05._Сотрудник_ТП"
  },
  {
    "type": "GET",
    "url": "/support/status_list",
    "title": "02. Получение списка возможных статусов для работников ТП",
    "name": "GetStatusesForSupport",
    "group": "05._Сотрудник_ТП",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "SupportStatus[]",
            "optional": false,
            "field": "support_status",
            "description": "<p>Массив объектов &quot;статус сотрудника ТП&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "support_status.support_status_id",
            "description": "<p>ID статуса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "support_status.support_status_name",
            "description": "<p>Описание статуса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n  {\n      \"support_status_id\": 1,\n      \"support_status_name\": \"Принимаю запросы\"\n  },\n  {\n      \"support_status_id\": 2,\n      \"support_status_name\": \"Перерыв\"\n  },\n  {\n      \"support_status_id\": 3,\n      \"support_status_name\": \"Работа в офисе\"\n  },\n  {\n      \"support_status_id\": 4,\n      \"support_status_name\": \"Не принимаю запросы\"\n  }\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/support/supports/02_GetStatusesForSupport.go",
    "groupTitle": "05._Сотрудник_ТП"
  },
  {
    "type": "GET",
    "url": "/support/card",
    "title": "12. Получить карточку сотрудника тех.поддержки",
    "name": "GetSupportCard",
    "group": "05._Сотрудник_ТП",
    "version": "2.0.0",
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/support/card?id=13",
        "type": "json"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "id",
            "description": "<p>ИД карточки супорта</p>"
          },
          {
            "group": "Success 200",
            "type": "CardUser",
            "optional": false,
            "field": "support",
            "description": "<p>Объект содержащий данные суппорта</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "support.id",
            "description": "<p>ИД суппорта</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "support.name",
            "description": "<p>ФИО суппорта</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "internal_number",
            "description": "<p>Внутренний номер телефонии</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "mobile_number",
            "description": "<p>Мобильный номер телефона</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "birth_date",
            "description": "<p>Дата рождения</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "is_senior",
            "description": "<p>Признак старшего группы саппортов</p>"
          },
          {
            "group": "Success 200",
            "type": "CardUser",
            "optional": false,
            "field": "senior",
            "description": "<p>Объект содержащий данные о старшем суппорте в группу которого входит текущий сотрудник</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "senior.id",
            "description": "<p>ИД суппорта</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "senior.name",
            "description": "<p>ФИО суппорта</p>"
          },
          {
            "group": "Success 200",
            "type": "Float64",
            "optional": false,
            "field": "wager",
            "description": "<p>Ставка за смену</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "comment",
            "description": "<p>Комментарий</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "color",
            "description": "<p>Цвет отображения (в шестнадцатиричной системе)</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Карточка старшего смены:",
          "content": "{\n  \"id\": 13,\n  \"support\": {\n    \"id\": 4,\n    \"name\": \"Вячеслав Викторович Тищенко\"\n  },\n  \"internal_number\": \"1484\",\n  \"mobile_number\": \"\",\n  \"birth_date\": \"\",\n  \"is_senior\": true,\n  \"senior\": null,\n  \"wager\": 500,\n  \"comment\": \"test\",\n  \"color\": \"0xFFFFF0\"\n}",
          "type": "json"
        },
        {
          "title": "Карточка обычного суппорта:",
          "content": "{\n \"id\": 14,\n \"support\": {\n   \"id\": 5,\n   \"name\": \"Евгений Николаевич Табаков\"\n },\n \"internal_number\": \"1487\",\n \"mobile_number\": \"\",\n \"birth_date\": \"\",\n \"is_senior\": false,\n \"senior\": {\n   \"id\": 4,\n   \"name\": \"Вячеслав Викторович Тищенко\"\n },\n \"wager\": 500,\n \"comment\": \"test\",\n \"color\": \"0xFFFFF0\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/support/cards/12_GetCard.go",
    "groupTitle": "05._Сотрудник_ТП"
  },
  {
    "type": "GET",
    "url": "/support/get_current_statuses",
    "title": "10. Получение списка сотрудников ТП с их текущим рабочим статусом",
    "name": "GetSupportCurrentStatuses",
    "group": "05._Сотрудник_ТП",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус выполнения запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "support_current_status[]",
            "optional": false,
            "field": "support_current_status",
            "description": "<p>Массив статистика по запросам и текущих статусов сотрудников ТП работающих сегодня</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "support_current_status.support_id",
            "description": "<p>ID сотрудника ТП</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "support_current_status.support",
            "description": "<p>Имя сотрудника ТП</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "support_current_status.status",
            "description": "<p>Текущий статус сотрудника ТП</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "support_current_status.shift_status",
            "description": "<p>Текущий статус смены сотрудника ТП</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "support_current_status.in_work",
            "description": "<p>Количество запросов в работе у указанного сотрудника</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "support_current_status.postproned",
            "description": "<p>Количество отложеных запросов у указанного сотрудника</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "support_current_status.complete",
            "description": "<p>Количество выполненных за сегодня запросов у указанного сотрудника</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "support_current_status.revision",
            "description": "<p>Количество запросов за сегодня отправленных на доработку у указанного сотрудника</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "support_current_status.priority",
            "description": "<p>Признак приоритета по распределению запросов на саппорта</p>"
          },
          {
            "group": "Success 200",
            "type": "total",
            "optional": false,
            "field": "total",
            "description": "<p>Суммарная статистика по кол-ву запросов в работе, отложеных, выполненых по сотрудникам за сегодня</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "total.total_in_work",
            "description": "<p>Общее кол-во запросов в работе</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "total.total_postproned",
            "description": "<p>Общее кол-во отложенных запросов</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "total.total_complete",
            "description": "<p>Общее кол-во выполненных за сегодня запросов</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "total.total_revision",
            "description": "<p>Общее кол-во запросов отправленных на доработку за сегодня</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "wait_ticket_count",
            "description": "<p>Кол-во запросов в очереди на распределение</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"ok\",\n  \"support_current_status\": [\n    {\n      \"support_id\": 5,\n      \"support\": \"Артем Владимирович Шелкопляс\",\n      \"status\": \"Принимаю запросы\",\n      \"shift_status\": true,\n      \"in_work\": 1,\n      \"postproned\": 0,\n      \"complete\": 0\n    },\n    {\n      \"support_id\": 6,\n      \"support\": \"Вячеслав Викторович Тищенко\",\n      \"status\": \"Не принимаю запросы\",\n      \"shift_status\": false,\n      \"in_work\": 0,\n      \"postproned\": 0,\n      \"complete\": 0\n    }\n  ],\n  \"total\": {\n    \"total_in_work\": 1,\n    \"total_postproned\": 0,\n    \"total_complete\": 0\n  },\n  \"wait_ticket_count\": 0\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/support/supports/10_GetSupportCurrentStatuses.go",
    "groupTitle": "05._Сотрудник_ТП"
  },
  {
    "type": "GET",
    "url": "/helpdesk/support/support_list",
    "title": "01. Получение списка всех сотрудников ТП",
    "name": "GetSupportList",
    "group": "05._Сотрудник_ТП",
    "version": "2.0.0",
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/support/support_list",
        "type": "json"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Support[]",
            "optional": false,
            "field": "support",
            "description": "<p>Массив объектов &quot;Cотрудник ТП&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "support.user_id",
            "description": "<p>ID сотрудника</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "support.user_name",
            "description": "<p>ФИО сотрудника</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n  {\n    \"user_id\": 4,\n    \"user_name\": \"Евгений Николаевич Табаков\"\n  },\n  {\n    \"user_id\": 5,\n    \"user_name\": \"Артем Владимирович Шелкопляс\"\n  },\n  {\n    \"user_id\": 6,\n    \"user_name\": \"Вячеслав Викторович Тищенко\"\n  },\n  {\n    \"user_id\": 7,\n    \"user_name\": \"Александр Игоревич Кудряшов\"\n  }\n]",
          "type": "type"
        }
      ]
    },
    "filename": "./docs/2.0.0/support/supports/01_GetSupportList.go",
    "groupTitle": "05._Сотрудник_ТП"
  },
  {
    "type": "GET",
    "url": "/helpdesk/support/shift_status",
    "title": "05. Получение текущего статуса смены",
    "name": "GetSupportShiftStatus",
    "group": "05._Сотрудник_ТП",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "shift_status",
            "description": "<p>Статус смены. True - открыта</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Смена закрыта:",
          "content": "{\n  \"shift_status\": false,\n  \"status\": \"ok\"\n}",
          "type": "json"
        },
        {
          "title": "Смена открыта:",
          "content": "{\n  \"shift_status\": true,\n  \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/support/supports/05_GetShiftStatus.go",
    "groupTitle": "05._Сотрудник_ТП"
  },
  {
    "type": "GET",
    "url": "/helpdesk/support/get_support_status",
    "title": "03. Получение текущего статуса сотрудника ТП",
    "name": "GetSupportStatus",
    "group": "05._Сотрудник_ТП",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "support_status_id",
            "description": "<p>ID статуса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "support_status_name",
            "description": "<p>Описание статуса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"support_status_id\": 4,\n  \"support_status_name\": \"Не принимаю запросы\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/support/supports/03_GetSupportStatus.go",
    "groupTitle": "05._Сотрудник_ТП"
  },
  {
    "type": "POST",
    "url": "/helpdesk/support/open_shift",
    "title": "07. Открытие смены сотруднику ТП",
    "name": "OpenShift",
    "group": "05._Сотрудник_ТП",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "support_id",
            "description": "<p>ID сотрудника которому нужно открыть смену</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"support_id\" : 6\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/support/supports/07_ShiftOpen.go",
    "groupTitle": "05._Сотрудник_ТП"
  },
  {
    "type": "POST",
    "url": "/support/card/update",
    "title": "14. Обновить карточку суппорта",
    "name": "UpdateSupportCard",
    "group": "05._Сотрудник_ТП",
    "version": "2.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "id",
            "description": "<p>ИД карточки</p>"
          },
          {
            "group": "Parameter",
            "type": "Stirng",
            "optional": false,
            "field": "internal_number",
            "description": "<p>Внутренний номер телефонии</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "mobile_number",
            "description": "<p>Мобильный номер (в одном из форматов: 0671112233, +380671112233, 380671112233, 067 111 22 33, 067-111-22-33)</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "birth_date",
            "description": "<p>Дата рождения (в формате ДД.ММ.ГГГГ)</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "is_senior",
            "description": "<p>Признак старшего суппорта</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "senior_id",
            "description": "<p>ИД старшего суппорта (если есть)</p>"
          },
          {
            "group": "Parameter",
            "type": "Float64",
            "optional": false,
            "field": "wager",
            "description": "<p>Ставка за смену</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "comment",
            "description": "<p>Комментарий</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "color",
            "description": "<p>Цвет отображения (если назначен старший цвет берется из карточки старшего)</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n  \"id\": 14,\n  \"internal_number\": \"1487\",\n  \"mobile_number\": \"\",\n  \"birth_date\": \"\",\n  \"is_senior\": false,\n  \"senior_id\": 4,\n  \"wager\": 500,\n  \"comment\": \"test\",\n  \"color\": \"0xFFFFF\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус выполнения запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/support/cards/14_UpdateCard.go",
    "groupTitle": "05._Сотрудник_ТП"
  },
  {
    "type": "GET",
    "url": "/helpdesk/reports/average_grades",
    "title": "03. Получение списка средних оценок за запросы",
    "name": "GetAverageGrades",
    "group": "06._Отчёты",
    "version": "2.0.0",
    "description": "<p>Получение списка усредненных оценок запросов по каждому сотруднику тех. поддержки. Последняя запись - усредненная оценка по отделу.</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "start_date",
            "description": "<p>Дата начала выборки, включительно</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "end_date",
            "description": "<p>Дата конца выборки, данные за этот день не учитываются</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/reports/average_grades?start_date=2021-05-01&end_date=2021-06-01",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "AverageGrade[]",
            "optional": false,
            "field": "average_grade",
            "description": "<p>Массив объектов <code>средняя оценка за запросы по сотруднику</code></p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "average_grade.support",
            "description": "<p>Сотрудник тех. поддержки</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "average_grade.average_grade_by_support",
            "description": "<p>Средняя оценка запросов</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n    {\n        \"support\": \"Вячеслав Викторович Тищенко\",\n        \"average_grade_by_support\": 5\n    },\n    {\n        \"support\": \"Отдел ТП\",\n        \"average_grade_by_support\": 5\n    }\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/reports/03_GetAverageGrades.go",
    "groupTitle": "06._Отчёты"
  },
  {
    "type": "GET",
    "url": "/helpdesk/reports/motivation",
    "title": "01. Отображние мотивации сотрудников ТП",
    "name": "GetMotivation",
    "group": "06._Отчёты",
    "version": "2.0.0",
    "description": "<p>Получение мотивации сотрудников тп в виде таблицы, где данные рассортированы по сотрудникам тп и категориям запросов. Каждый объект отображает название категории, количество выполненых запросов, и мотивацию за их выполнение. Последний обект &quot;Total&quot; отображает сумарное количество выполненых сотрудником запросов и его полную мотивацию.</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "start_date",
            "description": "<p>Дата начала выборки, включительно</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "end_date",
            "description": "<p>Дата конца выборки, данные за этот день не учитываются</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/reports/motivation?start_date=2021-03-01&end_date=2021-05-01",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "[]Period",
            "optional": false,
            "field": "period",
            "description": "<p>Результат за указанный период разделенный по месяцам</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Motivation",
            "optional": false,
            "field": "period.motivation",
            "description": "<p>Массив мотиваций по сотрудникам тех. поддержки</p>"
          },
          {
            "group": "Success 200",
            "type": "Support",
            "optional": false,
            "field": "period.motivation.support",
            "description": "<p>Объект с информацией по сотруднику тех. поддержки</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "period.motivation.support.id",
            "description": "<p>ID сотрудника тех. поддержки</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "period.motivation.support.name",
            "description": "<p>Имя сотрудника тех. поддержки</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "period.motivation.support.color",
            "description": "<p>Цвет для отображения сотрудника тех. поддержки в графике</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Categories",
            "optional": false,
            "field": "period.motivation.categories",
            "description": "<p>Массив объектов с информацией по количеству запросов выполненных сотрудником в разрезе категорий запросов</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "period.motivation.categories.id",
            "description": "<p>ID категории запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "period.motivation.categories.name",
            "description": "<p>Название категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "period.motivation.categories.tickets_count",
            "description": "<p>Количество запросов по категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "period.motivation.total_tickets_count",
            "description": "<p>Общее количество запросов по категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Float64",
            "optional": false,
            "field": "period.motivation.total_motivation",
            "description": "<p>Общая мотивация по сотруднику тех. поддержки</p>"
          },
          {
            "group": "Success 200",
            "type": "Float64",
            "optional": false,
            "field": "period.motivation.total_by_shifts",
            "description": "<p>Общая сумма оплаты по открытым сменам сотрудника за указанный период.</p>"
          },
          {
            "group": "Success 200",
            "type": "Float64",
            "optional": false,
            "field": "period.motivation.total_payment",
            "description": "<p>Общая сумма оплаты с учетом мотивации и оплаты смен.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"2021-10-01 ~ 2021-10-30\": [\n    {\n      \"support\": {\n        \"id\": 4,\n        \"name\": \"Вячеслав Викторович Тищенко\",\n        \"color\": \"\"\n      },\n      \"categories\": [\n        {\n          \"id\": 1,\n          \"name\": \"Сервисная категория\",\n          \"tickets_count\": 3\n        }\n      ],\n      \"total_tickets_count\": 0,\n      \"total_motivation\": 3,\n      \"total_by_shifts\": 1500,\n      \"total_payment\": 1503\n    },\n    {\n      \"support\": {\n        \"id\": 5,\n        \"name\": \"Артем Владимирович Шелкопляс\",\n        \"color\": \"0xFFFFFF\"\n      },\n      \"categories\": [],\n      \"total_tickets_count\": 0,\n      \"total_motivation\": 0,\n      \"total_by_shifts\": 0,\n      \"total_payment\": 0\n    },\n    {\n      \"support\": {\n        \"id\": 6,\n        \"name\": \"Евгений Николаевич Табаков\",\n        \"color\": \"\"\n      },\n      \"categories\": [],\n      \"total_tickets_count\": 0,\n      \"total_motivation\": 0,\n      \"total_by_shifts\": 0,\n      \"total_payment\": 0\n    },\n    {\n      \"support\": {\n        \"id\": 0,\n        \"name\": \"Итого\",\n        \"color\": \"\"\n      },\n      \"categories\": [\n        {\n          \"id\": 1,\n          \"name\": \"Сервисная категория\",\n          \"tickets_count\": 3\n        }\n      ],\n      \"total_tickets_count\": 0,\n      \"total_motivation\": 3,\n      \"total_by_shifts\": 1500,\n      \"total_payment\": 1503\n    }\n  ]\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/reports/01_GetMotivation.go",
    "groupTitle": "06._Отчёты"
  },
  {
    "type": "POST",
    "url": "/helpdesk/reports/tickets_grades",
    "title": "04. Получение списка оценок запросов",
    "name": "GetReturnedTickets",
    "group": "06._Отчёты",
    "version": "2.0.0",
    "description": "<p>Принимает в себя список ид пользователей и/или список отделов пользователей по которым нужно получить список оцененных запросов и их оценки</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "start_date",
            "description": "<p>Дата начала периода выборки</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "end_date",
            "description": "<p>Дата конца периода выборки</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64[]",
            "optional": false,
            "field": "users_id",
            "description": "<p>Массив ид пользователей по которым нужно сделать выборку</p>"
          },
          {
            "group": "Parameter",
            "type": "String[]",
            "optional": false,
            "field": "departments",
            "description": "<p>Массив отделов по которым нужно сделать выборку</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"start_date\" : \"2021-06-01\",\n    \"end_date\" : \"2021-06-20\",\n    \"users_id\":\n            [\n                4,\n                5,\n                6\n            ],\n\t\t\"departments\": []\n}",
          "type": "type"
        }
      ]
    },
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "DepartmentsTicketsGrades[]",
            "optional": false,
            "field": "departments_tickets_grades",
            "description": "<p>Массив обектов &quot;оценки запросов по пользователям по отделу&quot;</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "departments_tickets_grades.department",
            "description": "<p>Название отдела</p>"
          },
          {
            "group": "200",
            "type": "UsersTicketsGrades[]",
            "optional": false,
            "field": "departments_tickets_grades.users_grades",
            "description": "<p>Массив обектов &quot;оценки запросов по пользователю&quot;</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "departments_tickets_grades.users_grades.user_name",
            "description": "<p>ФИО пользователя</p>"
          },
          {
            "group": "200",
            "type": "TicketsGrades[]",
            "optional": false,
            "field": "departments_tickets_grades.users_grades.tickets_grades",
            "description": "<p>Массив обектов &quot;оценка запроса&quot;</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "departments_tickets_grades.users_grades.tickets_grades.ticket_id",
            "description": "<p>ИД запроса</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "departments_tickets_grades.users_grades.tickets_grades.ticket_grade",
            "description": "<p>Оценка запроса</p>"
          },
          {
            "group": "200",
            "type": "Float64",
            "optional": false,
            "field": "departments_tickets_grades.users_grades.average_user_grade",
            "description": "<p>Средняя оценка по пользователю</p>"
          },
          {
            "group": "200",
            "type": "Float64",
            "optional": false,
            "field": "departments_tickets_grades.avarege_department_grade",
            "description": "<p>Средняя оценка по отделу</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n  {\n    \"department\": \"Техническая поддержка\",\n    \"users_grades\": [\n      {\n        \"user_name\": \"Вячеслав Викторович Тищенко\",\n        \"tickets_grades\": [\n          {\n            \"ticket_id\": 1,\n            \"ticket_grade\": 5\n          },\n          {\n            \"ticket_id\": 3,\n            \"ticket_grade\": 5\n          },\n          {\n            \"ticket_id\": 4,\n            \"ticket_grade\": 2\n          },\n          {\n            \"ticket_id\": 31,\n            \"ticket_grade\": 3\n          }\n        ],\n        \"average_user_grade\": 3.75\n      },\n      {\n        \"user_name\": \"Евгений Николаевич Табаков\",\n        \"tickets_grades\": [\n          {\n            \"ticket_id\": 19,\n            \"ticket_grade\": 4\n          },\n          {\n            \"ticket_id\": 22,\n            \"ticket_grade\": 5\n          }\n        ],\n        \"average_user_grade\": 4.5\n      }\n    ],\n    \"avarege_department_grade\": 4.13\n  },\n  {\n    \"department\": \"Разработчики\",\n    \"users_grades\": [\n      {\n        \"user_name\": \"Артем Владимирович Шелкопляс\",\n        \"tickets_grades\": [\n          {\n            \"ticket_id\": 2,\n            \"ticket_grade\": 3\n          },\n          {\n            \"ticket_id\": 5,\n            \"ticket_grade\": 5\n          },\n          {\n            \"ticket_id\": 15,\n            \"ticket_grade\": 5\n          }\n        ],\n        \"average_user_grade\": 4.33\n      }\n    ],\n    \"avarege_department_grade\": 4.33\n  }\n]",
          "type": "type"
        }
      ]
    },
    "filename": "./docs/2.0.0/reports/04_GetTicketsGrades.go",
    "groupTitle": "06._Отчёты"
  },
  {
    "type": "GET",
    "url": "/helpdesk/reports/returned_tickets",
    "title": "05. Список запросов возвращенных обратно в работу после статус выполнено или отклонено",
    "name": "GetReturnedTickets",
    "group": "06._Отчёты",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "start_date",
            "description": "<p>Дата начала выборки, включительно</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "end_date",
            "description": "<p>Дата конца выборки, данные за этот день не учитываются</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/reports/returned_tickets?start_date=2021-05-01&end_date=2021-05-07",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Ticket[]",
            "optional": false,
            "field": "ticket",
            "description": "<p>Массив объектов &quot;запрос&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "ticket.ticket_id",
            "description": "<p>Номер запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket.ticket_date",
            "description": "<p>Дата создания запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket.category",
            "description": "<p>Категория запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket.section",
            "description": "<p>Раздел категории запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket.ticket_text",
            "description": "<p>Текст запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket.status",
            "description": "<p>Текущий статус запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket.author",
            "description": "<p>ФИО автора запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket.support",
            "description": "<p>ФИО сотрудника тех.поддержки</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "ticket.ticket_grade",
            "description": "<p>Оценка запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n  {\n    \"ticket_id\": 23,\n    \"ticket_date\": \"2021-06-11 15:27:37\",\n    \"category\": \"Оборудование\",\n    \"section\": \"Проблема с принтером/печатью\",\n    \"ticket_text\": \"Не печатает принтер\",\n    \"status\": \"В работе\",\n    \"author\": \"Евгений Николаевич Табаков\",\n    \"support\": \"Артем Владимирович Шелкопляс\",\n    \"ticket_grade\": 0\n  },\n  {\n    \"ticket_id\": 19,\n    \"ticket_date\": \"2021-06-09 16:39:54\",\n    \"category\": \"1С\",\n    \"section\": \"Изменение/удаление кассовых ордеров\",\n    \"ticket_text\": \"удалить кассовый ордер\",\n    \"status\": \"Выполнен\",\n    \"author\": \"Евгений Николаевич Табаков\",\n    \"support\": \"Артем Владимирович Шелкопляс\",\n    \"ticket_grade\": 4\n  }\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/reports/05_GetReturnedTickets.go",
    "groupTitle": "06._Отчёты"
  },
  {
    "type": "GET",
    "url": "/helpdesk/reports/supports_statuses_history",
    "title": "09. История статусов суппортов за определенную дату",
    "name": "GetSupportStatusHistory",
    "group": "06._Отчёты",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "date",
            "description": "<p>Дата за которую необходимо получить статусы</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/reports/supports_statuses_history?date=2021-07-20",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "SupportStatus[]",
            "optional": false,
            "field": "support_status",
            "description": "<p>Массив объектов &quot;статус суппорта&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "support_status.support",
            "description": "<p>Имя суппорта</p>"
          },
          {
            "group": "Success 200",
            "type": "Statuses[]",
            "optional": false,
            "field": "support_status.statuses",
            "description": "<p>Массив объектов &quot;статусы&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "support_status.statuses.time",
            "description": "<p>Время выбора статуса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "support_status.statuses.name",
            "description": "<p>Имя выбранного статуса</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "support_status.statuses.difference",
            "description": "<p>Длительность нахождения в статусе в секундах</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n {\n   \"support\": \"Артем Владимирович Шелкопляс\",\n   \"statuses\": [\n     {\n       \"time\": \"14:27:44\",\n       \"name\": \"Принимаю запросы\",\n       \"difference\": 146\n     },\n     {\n       \"time\": \"14:30:10\",\n       \"name\": \"Работа в офисе\",\n       \"difference\": 101\n     },\n     {\n      \"time\": \"14:31:51\",\n      \"name\": \"Принимаю запросы\",\n      \"difference\": 0\n    }\n   ]\n },\n {\n   \"support\": \"Александр Игоревич Кудряшов\",\n   \"statuses\": [\n     {\n       \"time\": \"14:28:51\",\n       \"name\": \"Принимаю запросы\",\n       \"difference\": 89\n     },\n     {\n       \"time\": \"14:30:20\",\n       \"name\": \"Работа в офисе\",\n       \"difference\": 15\n     },\n     {\n      \"time\": \"14:30:35\",\n      \"name\": \"Принимаю запросы\",\n      \"difference\": 0\n    }\n   ]\n },\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/reports/09_GetSupportStatusHistory.go",
    "groupTitle": "06._Отчёты"
  },
  {
    "type": "GET",
    "url": "/helpdesk/reports/supports_shifts",
    "title": "08. Время открытия и закрытия смен супортов за период",
    "name": "GetSupportsShiftsOpeningTime",
    "group": "06._Отчёты",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "start_date",
            "description": "<p>Дата начала выборки</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "end_date",
            "description": "<p>Дата конца выборки</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/reports/supports_shifts?start_date=2021-08-04&end_date=2021-11-10",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "[]Period",
            "optional": false,
            "field": "period",
            "description": "<p>Результат за указанный период разделенный по месяцам</p>"
          },
          {
            "group": "200",
            "type": "[]SupportShifts",
            "optional": false,
            "field": "period.support_shifts",
            "description": "<p>Массив объектов &quot;смены суппорта&quot;</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "period.support_shifts.support",
            "description": "<p>ФИО суппорта</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "period.support_shifts.with_out_grace_time",
            "description": "<p>Время опоздания свыше льготного периода</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "period.support_shifts.shifts_count",
            "description": "<p>Количество смен</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "period.support_shifts.total_minutes_count",
            "description": "<p>Общее время опоздания без учета льготного периода</p>"
          },
          {
            "group": "200",
            "type": "[]Shift",
            "optional": false,
            "field": "period.support_shifts.shifts",
            "description": "<p>Массив объектов &quot;смена&quot;</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "period.support_shifts.shifts.opening_date",
            "description": "<p>Время открытия смены</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "period.support_shifts.shifts.closing_date",
            "description": "<p>Время закрытия смены</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "period.support_shifts.shifts.count_of_minutes_late",
            "description": "<p>Опоздание за эту смены в минутах</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"2021-08-04 ~ 2021-08-31\": [],\n  \"2021-09-01 ~ 2021-09-30\": [],\n  \"2021-10-01 ~ 2021-10-31\": [],\n  \"2021-11-01 ~ 2021-11-10\": [\n    {\n      \"support\": \"Артем Владимирович Шелкопляс\",\n      \"with_out_grace_time\": \"6h28m0s\",\n      \"shifts_count\": 1,\n      \"total_minutes_count\": \"6h48m0s\",\n      \"shifts\": [\n        {\n          \"opening_date\": \"2021-11-03 14:48:00\",\n          \"closing_date\": \" \",\n          \"count_of_minutes_late\": 408\n        }\n      ]\n    },\n    {\n      \"support\": \"Вячеслав Викторович Тищенко\",\n      \"with_out_grace_time\": \"0s\",\n      \"shifts_count\": 1,\n      \"total_minutes_count\": \"0s\",\n      \"shifts\": [\n        {\n          \"opening_date\": \"2021-11-04 14:50:38\",\n          \"closing_date\": \"2021-11-09 09:51:23\",\n          \"count_of_minutes_late\": 0\n        }\n      ]\n    }\n  ]\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/reports/08_GetSupportsShifts.go",
    "groupTitle": "06._Отчёты"
  },
  {
    "type": "GET",
    "url": "/helpdesk/reports/supports_statuses",
    "title": "07. Время нахождения сотрудника ТП в определенном статусе в разрезе дней недели",
    "name": "GetSupportsStatusesByWeekDay",
    "group": "06._Отчёты",
    "version": "2.0.0",
    "description": "<p>Отчет показывает сколько времени было проведено сотрудником в том или ином статусе в разрезе дней недели за указанный промежуток времени. Если в какой-то из дней недели небыло сотрудников у которых менялся статус - то в этот день список сотрудников будет пустым.</p>",
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/reports/supports_statuses?start_date=2021-05-01&end_date=2021-06-01",
        "type": "json"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "start_date",
            "description": "<p>Дата начала выборки, включительно</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "end_date",
            "description": "<p>Дата конца выборки, данные за этот день не учитываются</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "week_day",
            "description": "<p>ID дня недели: 0 - понедельник, 1 - вторник, 2 - среда, 3 - четверг, 4 - пятница, 5 - суббота, 6 - воскресенье</p>"
          },
          {
            "group": "Success 200",
            "type": "SupportStatusesList[]",
            "optional": false,
            "field": "supports_list",
            "description": "<p>Список сотрудников тех поддержки с указанием статусов в которых они были за указанный промежуток времени и времени которое в них провели</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "supports_list.support_name",
            "description": "<p>ФИО Сотрудника тех поддержки</p>"
          },
          {
            "group": "Success 200",
            "type": "SupportStatusDuration[]",
            "optional": false,
            "field": "supports_list.statuses",
            "description": "<p>Список статусов которые были у сотрудника тех поддержки за указанный промежуток времени с указанием времени проведенном в этом статусе</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "supports_list.statuses.status_name",
            "description": "<p>Название статуса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "supports_list.statuses.duration",
            "description": "<p>Время проведенное в данном статусе в формате &quot;00h00m00s&quot;</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n   {\n       \"week_day\": 0,\n       \"supports_list\": null\n   },\n   {\n       \"week_day\": 1,\n       \"supports_list\": [\n           {\n               \"support_name\": \"Евгений Николаевич Табаков\",\n               \"statuses\": [\n                   {\n                       \"status_name\": \"Принимаю запросы\",\n                       \"duration\": \"7h22m41s\"\n                   },\n                   {\n                       \"status_name\": \"Перерыв\",\n                       \"duration\": \"2m25s\"\n                   },\n                   {\n                       \"status_name\": \"Не принимаю запросы\",\n                       \"duration\": \"5m23s\"\n                   }\n               ]\n           },\n           {\n               \"support_name\": \"Вячеслав Викторович Тищенко\",\n               \"statuses\": [\n                   {\n                       \"status_name\": \"Принимаю запросы\",\n                       \"duration\": \"7h17m42s\"\n                   },\n                   {\n                       \"status_name\": \"Работа в офисе\",\n                       \"duration\": \"5m27s\"\n                   }\n               ]\n           }\n       ]\n   },\n   {\n       \"week_day\": 2,\n       \"supports_list\": null\n   },\n   {\n       \"week_day\": 3,\n       \"supports_list\": null\n   },\n   {\n       \"week_day\": 4,\n       \"supports_list\": [\n           {\n               \"support_name\": \"Евгений Николаевич Табаков\",\n               \"statuses\": [\n                   {\n                       \"status_name\": \"Принимаю запросы\",\n                       \"duration\": \"6h30m57s\"\n                   }\n               ]\n           },\n           {\n               \"support_name\": \"Вячеслав Викторович Тищенко\",\n               \"statuses\": [\n                   {\n                       \"status_name\": \"Принимаю запросы\",\n                       \"duration\": \"64727h46m40s\"\n                   }\n               ]\n           }\n       ]\n   },\n   {\n       \"week_day\": 5,\n       \"supports_list\": null\n   },\n   {\n       \"week_day\": 6,\n       \"supports_list\": null\n   }\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/reports/07_GetSupportsStatusesByWeekDay.go",
    "groupTitle": "06._Отчёты"
  },
  {
    "type": "GET",
    "url": "/helpdesk/reports/tickets_status_difference",
    "title": "02. Время нахождения запроса в разных статусах",
    "name": "GetTicketStatusDifference",
    "group": "06._Отчёты",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "start_date",
            "description": "<p>Дата начала выборки, включительно</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "end_date",
            "description": "<p>Дата конца выборки, данные за этот день не учитываются</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/reports/tickets_status_difference?start_date=2021-05-01&end_date=2021-06-01",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "TicketStatuDifferenceTime[]",
            "optional": false,
            "field": "ticket_status_difference_time",
            "description": "<p>Массив обектов &quot;время нахождения запроса в разных статусах&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "ticket_status_difference_time.ticket_id",
            "description": "<p>Ид запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket_status_difference_time.support_name",
            "description": "<p>Имя сотрудника тех.-поддержки</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket_status_difference_time.section",
            "description": "<p>Имя раздела категории</p>"
          },
          {
            "group": "Success 200",
            "type": "StatusDifference[]",
            "optional": false,
            "field": "ticket_status_difference_time.status_difference",
            "description": "<p>Масиив обектов &quot;время нахождения в статусе&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket_status_difference_time.status_difference.status",
            "description": "<p>Имя статуса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "ticket_status_difference_time.status_difference.diff_time",
            "description": "<p>Время нахождения в этом статусе</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n    {\n        \"ticket_id\": 3,\n        \"support_name\": \"Вячеслав Викторович Тищенко\",\n        \"section\": \"Прочее\",\n        \"status_difference\": [\n            {\n                \"status\": \"В ожидании\",\n                \"diff_time\": \"12m35s\"\n            },\n            {\n                \"status\": \"Отправлен на доработку\",\n                \"diff_time\": \"1m26s\"\n            }\n        ]\n    },\n    {\n        \"ticket_id\": 4,\n        \"support_name\": \"Вячеслав Викторович Тищенко\",\n        \"section\": \"Прочее\",\n        \"status_difference\": [\n            {\n                \"status\": \"В ожидании\",\n                \"diff_time\": \"5s\"\n            },\n            {\n                \"status\": \"В работе\",\n                \"diff_time\": \"2m42s\"\n            },\n            {\n                \"status\": \"В процессе реализации\",\n                \"diff_time\": \"17s\"\n            },\n            {\n                \"status\": \"Отправлен на доработку\",\n                \"diff_time\": \"18s\"\n            },\n            {\n                \"status\": \"Отложен\",\n                \"diff_time\": \"30s\"\n            }\n        ]\n    }\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/reports/02_GetTicketStatusDifference.go",
    "groupTitle": "06._Отчёты"
  },
  {
    "type": "GET",
    "url": "/helpdesk/reports/tickets_count",
    "title": "06. Количество поступивших запросов в разрезе часов и дней",
    "name": "GetTicketsCountByDaysHours",
    "group": "06._Отчёты",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "start_date",
            "description": "<p>Дата начала выборки, включительно</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "end_date",
            "description": "<p>Дата конца выборки, данные за этот день не учитываются</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/reports/tickets_count?start_date=2021-04-01&end_date=2021-05-01",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "CountByDay[]",
            "optional": false,
            "field": "count_by_day",
            "description": "<p>Массив обектов &quot;Количество запросов за часы дня&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "count_by_day.day",
            "description": "<p>Дата</p>"
          },
          {
            "group": "Success 200",
            "type": "CountByHour[]",
            "optional": false,
            "field": "count_by_day.count_by_hour",
            "description": "<p>Массив обектов &quot;Количество запросов за час&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "count_by_day.count_by_hour.hour",
            "description": "<p>Временной диапазон</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint",
            "optional": false,
            "field": "count_by_day.count_by_hour.count",
            "description": "<p>Количество запросов</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n  {\n    \"date\": \"2021-06-01\",\n    \"count_by_hour\": [\n      {\n        \"hour\": \"19:00:00 - 19:59:59\",\n        \"count\": 2\n      },\n      {\n        \"hour\": \"20:00:00 - 20:59:59\",\n        \"count\": 6\n      }\n    ]\n  },\n  {\n    \"date\": \"2021-06-02\",\n    \"count_by_hour\": [\n      {\n        \"hour\": \"14:00:00 - 14:59:59\",\n        \"count\": 1\n      },\n      {\n        \"hour\": \"19:00:00 - 19:59:59\",\n        \"count\": 1\n      }\n    ]\n  },\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/reports/06_GetTicketsCountByDaysHours.go",
    "groupTitle": "06._Отчёты"
  },
  {
    "type": "POST",
    "url": "/helpdesk/filial/create",
    "title": "04. Добаление отделения в базу",
    "name": "CreateFilial",
    "group": "07._Настройка",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "region_id",
            "description": "<p>ID региона в базе</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "filial",
            "description": "<p>Название отделения</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ip",
            "description": "<p>Первые 3 октета ip-адреса отделения</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Создание филиала:",
          "content": "{\n\t \t\"region_id\": 1,\n    \t\"filial\":\"Николаевское отделение №3\",\n    \t\"ip\":\"10.54.3\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "filial_id",
            "description": "<p>ID созданого отделения</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Создание филиала:",
          "content": "{\n    \"filial_id\": 3,\n    \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "FilialErr_Exist",
            "description": "<p>Отделение с таким ip уже существует</p>"
          }
        ]
      }
    },
    "filename": "./docs/2.0.0/regions_and_filials/04_CreateFilial.go",
    "groupTitle": "07._Настройка"
  },
  {
    "type": "POST",
    "url": "/helpdesk/region/create",
    "title": "01. Создание региона",
    "name": "CreateRegion",
    "group": "07._Настройка",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "region",
            "description": "<p>Название региона</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n        \"region\":\"Николаевская область\"\n}",
          "type": "type"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "region_id",
            "description": "<p>ID созданого отделения</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"status\" : \"ok\",\n    \"region_id\" : 1,\n}",
          "type": "type"
        }
      ]
    },
    "filename": "./docs/2.0.0/regions_and_filials/01_CreateRegion.go",
    "groupTitle": "07._Настройка"
  },
  {
    "type": "DELETE",
    "url": "/helpdesk/filial/",
    "title": "06. Удаление филиала из базы",
    "name": "DeleteFilial",
    "group": "07._Настройка",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/filial/?filial_id=23",
        "type": "json"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "filial_id",
            "description": "<p>ID отделения в базе</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ErrFilialDoesNotExist",
            "description": "<p>Такой филиал отсутствует в базе</p>"
          }
        ]
      }
    },
    "filename": "./docs/2.0.0/regions_and_filials/06_DeleteFilial.go",
    "groupTitle": "07._Настройка"
  },
  {
    "type": "DELETE",
    "url": "/helpdesk/region/",
    "title": "03. Удаление региона из базы",
    "name": "DeleteRegion",
    "group": "07._Настройка",
    "version": "2.0.0",
    "description": "<p>При удалении региона удаляются также и филиалы которые в него входят</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/region/?region_id=23",
        "type": "json"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "region_id",
            "description": "<p>ID отделения в базе</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/regions_and_filials/03_DeleteRegion.go",
    "groupTitle": "07._Настройка"
  },
  {
    "type": "GET",
    "url": "/helpdesk/const/banner",
    "title": "02. Получение текущего текста баннера",
    "name": "GetBanner",
    "group": "07._Настройка",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "text",
            "description": "<p>Текст баннера</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\t \t\t\"text\":\"text\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/consts/02_GetBanner.go",
    "groupTitle": "07._Настройка"
  },
  {
    "type": "GET",
    "url": "/helpdesk/filial/filial_list",
    "title": "07. Получение списка отделений из базы",
    "name": "GetFilialList",
    "group": "07._Настройка",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "Region[]",
            "optional": false,
            "field": "region",
            "description": "<p>Массив объектов &quot;регион&quot;</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "region.region_id",
            "description": "<p>ID региона в базе</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "region.region",
            "description": "<p>Имя региона</p>"
          },
          {
            "group": "200",
            "type": "Filial[]",
            "optional": false,
            "field": "region.filials",
            "description": "<p>Массив отделений относящихся к региону</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "region.filials.filial_id",
            "description": "<p>ID отделения в базе</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "region.filials.filial",
            "description": "<p>Название отделения</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "region.filials.ip",
            "description": "<p>Первые 3 октета ip-адреса отделения</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[\n {\n   \"region_id\": 1,\n   \"region\": \"Николаевская область\",\n   \"filials\": [\n     {\n       \"filial_id\": 4,\n       \"filial\": \"Николаевское отделение №5\",\n       \"ip\": \"10.54.6\"\n     },\n     {\n       \"filial_id\": 3,\n       \"filial\": \"Николаевское отделение №2\",\n       \"ip\": \"10.54.2\"\n     },\n     {\n       \"filial_id\": 1,\n       \"filial\": \"Николаевское отделение №1\",\n       \"ip\": \"10.54.1\"\n     }\n   ]\n },\n {\n   \"region_id\": 2,\n   \"region\": \"Киевская область\",\n   \"filials\": [\n     {\n       \"filial_id\": 2,\n       \"filial\": \"Киевское отделение №1\",\n       \"ip\": \"10.1.1\"\n     }\n   ]\n },\n {\n   \"region_id\": 3,\n   \"region\": \"Одесская область\"\n }\n]",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/regions_and_filials/07_GetFilialList.go",
    "groupTitle": "07._Настройка"
  },
  {
    "type": "POST",
    "url": "/helpdesk/const/banner",
    "title": "01. Модификация текста баннера",
    "name": "SetBanner",
    "group": "07._Настройка",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "text",
            "description": "<p>Текст баннера</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t \t\t\"text\":\"text\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/consts/01_SetBanner.go",
    "groupTitle": "07._Настройка"
  },
  {
    "type": "POST",
    "url": "/helpdesk/filial/update",
    "title": "05. Обновление отделения в базе",
    "name": "UpdateFilial",
    "group": "07._Настройка",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "filial_id",
            "description": "<p>ID отделения в базе</p>"
          },
          {
            "group": "Parameter",
            "type": "Region",
            "optional": false,
            "field": "region",
            "description": "<p>Регион в котором находится отделение</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "region.region_id",
            "description": "<p>ID региона в котором находится отделение</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "filial",
            "description": "<p>Название отделения</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ip",
            "description": "<p>Первые 3 октета ip-адреса отделения</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t   \"filial_id\":1,\n    \"region\":{\n\t\t\"region_id\": 1\n},\n    \"filial\":\"Николаевское отделение №1\",\n    \"ip\":\"10.54.5\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/regions_and_filials/05_UpdateFilial.go",
    "groupTitle": "07._Настройка"
  },
  {
    "type": "POST",
    "url": "/helpdesk/region/update",
    "title": "02. Обновление региона в базе",
    "name": "UpdateRegion",
    "group": "07._Настройка",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "region_id",
            "description": "<p>ID региона в базе</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "region",
            "description": "<p>Название региона</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t\t\"region_id\": 2,\n\t\t\"region\": \"Киевская область\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/regions_and_filials/02_UpdateRegion.go",
    "groupTitle": "07._Настройка"
  },
  {
    "type": "GET",
    "url": "/helpdesk/table/lateness_conf",
    "title": "07. Получить настройки опозданий",
    "name": "GetLatenessConf",
    "group": "8._Табель",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "LatenessConf",
            "optional": false,
            "field": "lateness_conf",
            "description": "<p>Объект &quot;настройки опозданий&quot;</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "lateness_conf.late_penalty",
            "description": "<p>Штраф за минуту опоздания</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "lateness_conf.grace_time",
            "description": "<p>Количество льготных минут</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Cтатус выполнения</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"lateness_conf\": {\n    \"late_penalty\": 5,\n    \"grace_time\": 20\n  },\n  \"status\": \"ok\"\n}",
          "type": "type"
        }
      ]
    },
    "filename": "./docs/2.0.0/scheduler/07_GetLatenessConfig.go",
    "groupTitle": "8._Табель"
  },
  {
    "type": "GET",
    "url": "/helpdesk/table/offices_list",
    "title": "01. Получить список офисов",
    "name": "GetOfficesList",
    "group": "8._Табель",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "[]Office",
            "optional": false,
            "field": "actual",
            "description": "<p>Список актуальных офисов</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "actual.id",
            "description": "<p>ИД смены</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "actual.name",
            "description": "<p>Название смены</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "actual.color",
            "description": "<p>Цвет отображения в графике</p>"
          },
          {
            "group": "200",
            "type": "Bool",
            "optional": false,
            "field": "actual.deleted",
            "description": "<p>Признак удалена смена или нет.</p>"
          },
          {
            "group": "200",
            "type": "[]Office",
            "optional": false,
            "field": "deleted",
            "description": "<p>Список актуальных офисов</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "deleted.id",
            "description": "<p>ИД смены</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "deleted.name",
            "description": "<p>Название смены</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "deleted.color",
            "description": "<p>Цвет отображения в графике</p>"
          },
          {
            "group": "200",
            "type": "Bool",
            "optional": false,
            "field": "deleted.deleted",
            "description": "<p>Признак удалена смена или нет.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"actual\": [\n    {\n      \"id\": 9,\n      \"name\": \"вторая\",\n      \"color\": \"#BEFF2E\",\n      \"deleted\": false\n    },\n    {\n      \"id\": 10,\n      \"name\": \"артилерийская\",\n      \"color\": \"#1FC91EF2\",\n      \"deleted\": false\n    }\n  ],\n  \"deleted\": [\n    {\n      \"id\": 8,\n      \"name\": \"первая\",\n      \"color\": \"#487C7CFF\",\n      \"deleted\": true\n    }\n  ],\n  \"status\": \"ok\"\n}",
          "type": "Json"
        }
      ]
    },
    "filename": "./docs/2.0.0/scheduler/01_GetOfficesList.go",
    "groupTitle": "8._Табель"
  },
  {
    "type": "GET",
    "url": "/helpdesk/table/schedule",
    "title": "03. Получение графика смен",
    "name": "GetShiftsSchedule",
    "group": "8._Табель",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8888/helpdesk/table/schedule?date=2021-10-01",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "[]Office",
            "optional": false,
            "field": "legend",
            "description": "<p>Список оффисов для легенды графика (список офисов который включает в себя актуальные офисы, а также те удаленные которые уже были отмечены в графике за указанный период)</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "legend.id",
            "description": "<p>ИД смены</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "legend.name",
            "description": "<p>Название смены</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "legend.color",
            "description": "<p>Цвет для отображения на графике</p>"
          },
          {
            "group": "200",
            "type": "Bool",
            "optional": false,
            "field": "legend.deleted",
            "description": "<p>Признак являеться ли смена удаленной</p>"
          },
          {
            "group": "200",
            "type": "[]ShiftsScheduleCell",
            "optional": false,
            "field": "shifts_schedule",
            "description": "<p>Массив объектов &quot;ячейка графика смен&quot;</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "shifts_schedule.id",
            "description": "<p>ИД ячейки</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "shifts_schedule.support_id",
            "description": "<p>ИД сотрудника тп</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "shifts_schedule.shift_id",
            "description": "<p>ИД смены</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "shifts_schedule.date",
            "description": "<p>Дата смены</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "shifts_schedule.start_time",
            "description": "<p>Время начала смены</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "shifts_schedule.end_time",
            "description": "<p>Время конца смены</p>"
          },
          {
            "group": "200",
            "type": "Bool",
            "optional": false,
            "field": "shifts_schedule.sick_leave",
            "description": "<p>Признак больничного</p>"
          },
          {
            "group": "200",
            "type": "Bool",
            "optional": false,
            "field": "shifts_schedule.vacation",
            "description": "<p>Признак отпуска</p>"
          },
          {
            "group": "200",
            "type": "Bool",
            "optional": false,
            "field": "shifts_schedule.late",
            "description": "<p>Признак было ли опоздание в эту смену</p>"
          },
          {
            "group": "200",
            "type": "[]Support",
            "optional": false,
            "field": "regular_supports",
            "description": "<p>Список обычных суппортов</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "regular_supports.id",
            "description": "<p>ID суппорта</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "regular_supports.name",
            "description": "<p>Имя суппорта</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "regular_supports.color",
            "description": "<p>Цвет отображения</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "regular_supports.senior_id",
            "description": "<p>ID старшего суппорта (если не назначен прийдет <code>0</code>)</p>"
          },
          {
            "group": "200",
            "type": "[]Support",
            "optional": false,
            "field": "senior_supports",
            "description": "<p>Список старших суппортов</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "senior_supports.id",
            "description": "<p>ID суппорта</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "senior_supports.name",
            "description": "<p>Имя суппорта</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "senior_supports.color",
            "description": "<p>Цвет отображения</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "senior_supports.senior_id",
            "description": "<p>ID старшего суппорта (если не назначен прийдет <code>0</code>)</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"legend\": [\n    {\n      \"id\": 9,\n      \"name\": \"вторая\",\n      \"color\": \"#BEFF2E\",\n      \"deleted\": false\n    },\n    {\n      \"id\": 10,\n      \"name\": \"артилерийская\",\n      \"color\": \"#1FC91EF2\",\n      \"deleted\": false\n    },\n    {\n      \"id\": 11,\n      \"name\": \"Новая\",\n      \"color\": \"#D92B48\",\n      \"deleted\": false\n    },\n    {\n      \"id\": 12,\n      \"name\": \"*новая\",\n      \"color\": \"#677FD9FF\",\n      \"deleted\": false\n    }\n  ],\n  \"shifts_schedule\": [\n    {\n      \"id\": 62,\n      \"support_id\": 4,\n      \"office_id\": 9,\n      \"start_time\": \"08:00\",\n      \"end_time\": \"20:00\",\n      \"date\": \"2021-12-08\",\n      \"vacation\": false,\n      \"sick_leave\": true,\n      \"late\": false\n    },\n    {\n      \"id\": 63,\n      \"support_id\": 4,\n      \"office_id\": 9,\n      \"start_time\": \"08:00\",\n      \"end_time\": \"20:00\",\n      \"date\": \"2021-12-07\",\n      \"vacation\": false,\n      \"sick_leave\": false,\n      \"late\": false\n    }\n  ],\n  \"status\": \"ok\",\n\t \"regular_supports\": [\n\t   {\n\t     \"id\": 4,\n\t     \"name\": \"Вячеслав Викторович Тищенко\",\n\t     \"color\": \"0xFFFFFF\",\n\t     \"senior_id\": 5\n\t   },\n\t   {\n\t     \"id\": 6,\n\t     \"name\": \"Евгений Николаевич Табаков\",\n\t     \"color\": \"0xFFFFFF\",\n\t     \"senior_id\": 5\n\t   },\n\t   {\n\t     \"id\": 7,\n\t     \"name\": \"Владислав Сергеевич Маспанов\",\n\t     \"color\": \"0xFFFFFF\",\n\t     \"senior_id\": 0\n\t   }\n\t ],\n  \"senior_supports\": [\n   {\n     \"id\": 5,\n     \"name\": \"Артем Владимирович Шелкопляс\",\n     \"color\": \"0xFFFFFF\",\n     \"senior_id\": 0\n   }\n ],\n}",
          "type": "Json"
        }
      ]
    },
    "filename": "./docs/2.0.0/scheduler/03_GetShiftsSchedule.go",
    "groupTitle": "8._Табель"
  },
  {
    "type": "GET",
    "url": "/helpdesk/table/lateness",
    "title": "05. Получить список опозданий сотрудников ТП за месяц",
    "name": "GetSupportLateness",
    "group": "8._Табель",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "[]Decisions",
            "optional": false,
            "field": "desicions",
            "description": "<p>Список возможных решений по опозданию</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "desicions.id",
            "description": "<p>ID решения</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "desicions.text",
            "description": "<p>Текст решения</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Lateness",
            "optional": false,
            "field": "lateness",
            "description": "<p>Список опозданий сотрудников ТП</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "lateness.id",
            "description": "<p>ID записи об опоздании</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "lateness.date",
            "description": "<p>Дата и время создания записи об опоздании</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "lateness.name",
            "description": "<p>Имя сотрудника ТП</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "lateness.cause",
            "description": "<p>Причина опоздания ТП</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "lateness.decision_id",
            "description": "<p>ID решения (если <code>0</code> - решение по опозданию отсутствует)</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "lateness.difference",
            "description": "<p>Кол-во минут на которые опоздал сотрудник</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус выполнения запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"desicions\": [\n    {\n      \"id\": 1,\n      \"text\": \"Помилован\"\n    },\n    {\n      \"id\": 2,\n      \"text\": \"Казнен\"\n    }\n  ],\n  \"lateness\": [\n    {\n      \"id\": 1,\n      \"date\": \"2021-10-05 11:08:12\",\n      \"name\": \"Вячеслав Викторович Тищенко\"\n      \"cause\": \"test\",\n      \"decision_id\": 1,\n      \"difference\": 3\n    },\n    {\n      \"id\": 2,\n      \"date\": \"2021-10-02 11:08:12\",\n      \"name\": \"Вячеслав Викторович Тищенко\"\n      \"cause\": \"test2\",\n      \"decision_id\": 0,\n      \"difference\": 10\n    }\n  ],\n  \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/scheduler/05_GetSupportLateness.go",
    "groupTitle": "8._Табель"
  },
  {
    "type": "POST",
    "url": "/helpdesk/table/lateness_conf",
    "title": "08. Изменить настройки опозданий",
    "name": "SetLatenessConf",
    "group": "8._Табель",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "late_penalty",
            "description": "<p>Штраф за минуту опоздания</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "grace_time",
            "description": "<p>Количество льготных минут</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"late_penalty\":5,\n    \"grace_time\":20\n}",
          "type": "Json"
        }
      ]
    },
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Cтатус выполнения</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"status\": \"ok\"\n}",
          "type": "Json"
        }
      ]
    },
    "filename": "./docs/2.0.0/scheduler/08_SetLatenessConfig.go",
    "groupTitle": "8._Табель"
  },
  {
    "type": "POST",
    "url": "/helpdesk/table/lateness/update",
    "title": "06. Установить решение по опозданию",
    "name": "UpdateLateness",
    "group": "8._Табель",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "lateness_id",
            "description": "<p>ID записи об опоздании</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "decision_id",
            "description": "<p>ID решения по опозданию</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"lateness_id\": 1,\n    \"decision_id\": 1\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус выполнения запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/2.0.0/scheduler/06_UpdateLateness.go",
    "groupTitle": "8._Табель"
  },
  {
    "type": "POST",
    "url": "/helpdesk/table/offices_list",
    "title": "02. Обновление списка оффисов",
    "name": "UpdateOfficesList",
    "group": "8._Табель",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "[]Office",
            "optional": false,
            "field": "office",
            "description": "<p>Массив объектов &quot;Офис&quot;</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "office.id",
            "description": "<p>Ид смены, для новой смены 0 или без этого поля</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "office.name",
            "description": "<p>Название смены</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "office.start_time",
            "description": "<p>Время начала работы</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "office.end_time",
            "description": "<p>Время конца работы</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "office.color",
            "description": "<p>Цвет которым офис будет отображаться на графике</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "office.deleted",
            "description": "<p>Признак удален офис или нет.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "[\n  {\n    \"id\": 1,\n    \"name\": \"первая\",\n    \"start_time\": \"8:00\",\n    \"end_time\": \"20:00\",\n    \"color\": \"0xffffff\",\n    \"deleted\":false\n  },\n  {\n    \"id\":2,\n    \"name\": \"вторая\",\n    \"start_time\": \"9:00\",\n    \"end_time\": \"21:00\",\n    \"color\": \"0x3b6a32\",\n    \"deleted\":true\n  },\n    {\n    \"name\": \"артилерийская\",\n    \"start_time\": \"7:30\",\n    \"end_time\": \"19:30\",\n    \"color\": \"0x5b90f3\",\n    \"deleted\":false\n  }\n]",
          "type": "JSON"
        }
      ]
    },
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"status\": \"ok\"\n}",
          "type": "JSON"
        }
      ]
    },
    "filename": "./docs/2.0.0/scheduler/02_UpdateOfficesList.go",
    "groupTitle": "8._Табель"
  },
  {
    "type": "POST",
    "url": "/helpdesk/table/update_schedule",
    "title": "04. Обновить график смен",
    "name": "UpdateShiftsSchedule",
    "group": "8._Табель",
    "version": "2.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "[]ShiftsScheduleCell",
            "optional": false,
            "field": "shiftsSchedule",
            "description": "<p>Массив объектов &quot;ячейка графика смен&quot;</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "shiftsSchedule.id",
            "description": "<p>ИД ячейки. При добалении новой ячейки ид = 0</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "shiftsSchedule.support_id",
            "description": "<p>ИД сотрудника ТП</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "shiftsSchedule.office_id",
            "description": "<p>ИД смены</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "shiftsSchedule.date",
            "description": "<p>Дата в графике</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "shiftsSchedule.start_time",
            "description": "<p>Время начала смены</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "shiftsSchedule.end_time",
            "description": "<p>Время конца смены</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "shiftsSchedule.vacation",
            "description": "<p>Признак отпуска</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "shiftsSchedule.sick_leave",
            "description": "<p>Признак больничного</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Обновить смену в графике:",
          "content": "[\n  {\n    \"id\": 8,\n    \"support_id\": 4,\n    \"office_id\": 10,\n    \"date\": \"2021-10-15\"\n    \"start_time\": \"08:00\",\n    \"end_time\": \"20:00\",\n\t\t \"vacation\": false,\n    \"sick_leave\": false\n  }\n]",
          "type": "Json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"status\": \"ok\"\n}",
          "type": "Json"
        }
      ]
    },
    "filename": "./docs/2.0.0/scheduler/04_UpdateShiftsSchedule.go",
    "groupTitle": "8._Табель"
  }
] });
