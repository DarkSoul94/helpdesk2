
ALTER TABLE `category_section` ADD COLUMN `service` TINYINT NOT NULL DEFAULT 0;

INSERT INTO category_section SET
category_id = (SELECT category_id FROM category WHERE category_name = 'Сервисная категория'),
category_section_name = 'Запрос по api',
significant_category_section = false,
old_category_section = false,
need_approval = false,
`service` = true;
