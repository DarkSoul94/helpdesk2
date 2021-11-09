INSERT INTO category_section SET
category_id = (SELECT category_id FROM category WHERE category_name = 'Сервисная категория'),
category_section_name = 'Запрос по api приоритетный',
significant_category_section = true,
old_category_section = false,
need_approval = false,
`service` = true;
