
ALTER TABLE `category` ADD COLUMN `service` TINYINT NOT NULL DEFAULT 0;

INSERT INTO category SET
category_name = "Сервисная категория",
significant_category = false,
old_category = false,
price = 1.00,
`service` = true;
