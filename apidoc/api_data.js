define({ "api": [
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
    "title": "03. Создание раздела категории",
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
    "filename": "./docs/2.0.0/cat_end_sec/03_CreateCategorySection.go",
    "groupTitle": "02._Категории_и_разделы_категорий"
  },
  {
    "type": "GET",
    "url": "/helpdesk/section/",
    "title": "05. Получение списка разделов категорий без учета устаревших",
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
    "filename": "./docs/2.0.0/cat_end_sec/05_GetCategorySection.go",
    "groupTitle": "02._Категории_и_разделы_категорий"
  },
  {
    "type": "GET",
    "url": "/helpdesk/section/section_list",
    "title": "06. Получение всего списка разделов категорий",
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
    "filename": "./docs/2.0.0/cat_end_sec/06_GetCategorySectionList.go",
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
    "title": "04. Обновление разделов категории",
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
    "filename": "./docs/2.0.0/cat_end_sec/04_UpdateCategorySection.go",
    "groupTitle": "02._Категории_и_разделы_категорий"
  },
  {
    "type": "GET",
    "url": "/helpdesk/resolve_ticket/check_exist",
    "title": "06. Проверка есть ли запросы ожидающие согласования",
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
    "filename": "./docs/2.0.0/tickets/06_CheckNeedResolveTicketExist.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "post",
    "url": "/helpdesk/comment/create",
    "title": "12. Создание нового комментария в запросе",
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
    "filename": "./docs/2.0.0/tickets/comments/12_CreateCommentsHistory.go",
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
    "type": "get",
    "url": "/file/",
    "title": "13. Получение файла по его ид",
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
    "filename": "./docs/2.0.0/tickets/files/13_GetFile.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "GET",
    "url": "/helpdesk/resolve_ticket/resolve_tickets_list",
    "title": "07. Получение списка запросов в тех. поддержку ожидающих согласования",
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
    "filename": "./docs/2.0.0/tickets/07_GetResolveTicketsList.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "GET",
    "url": "/helpdesk/ticket/ticket",
    "title": "04. Получение запроса",
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
    "filename": "./docs/2.0.0/tickets/04_GetTicket.go",
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
    "title": "05. Получение списка запросов в тех. поддержку",
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
    "filename": "./docs/2.0.0/tickets/05_GetTicketsList.go",
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
    "title": "03. Обновление запроса в ТП",
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
    "filename": "./docs/2.0.0/tickets/03_UpdateTicket.go",
    "groupTitle": "03._Запросы_в_тех._поддержку"
  },
  {
    "type": "GET",
    "url": "/helpdesk/group/for_resolve",
    "title": "Получение списка груп пользователей с правами согласовывать запросы",
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
    "filename": "./docs/2.0.0/groups/GetGroupsListForResolve.go",
    "groupTitle": "04._Пользователи",
    "sampleRequest": [
      {
        "url": "http://localhost:8888//helpdesk/group/for_resolve"
      }
    ]
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
    "version": "0.1.1",
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
    "version": "0.1.1",
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
    "version": "0.1.1",
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
  }
] });
