
CREATE TABLE IF NOT EXISTS `category` (
  `category_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `category_name` VARCHAR(255),
  `significant_category` TINYINT,
  `old_category` TINYINT,
  `price` INT UNSIGNED NOT NULL DEFAULT '1',
  PRIMARY KEY `pk_id`(`category_id`)
);
