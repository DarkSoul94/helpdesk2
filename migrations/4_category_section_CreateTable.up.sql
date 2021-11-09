
CREATE TABLE IF NOT EXISTS `category_section` (
  `section_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `category_id` INT UNSIGNED,
  `category_section_name` VARCHAR(255),
  `significant_category_section` TINYINT,
  `old_category_section` TINYINT,
  `need_approval` TINYINT,
  PRIMARY KEY `pk_id`(`section_id`),
  CONSTRAINT FOREIGN KEY (`category_id`) REFERENCES `category`(`category_id`)
);
