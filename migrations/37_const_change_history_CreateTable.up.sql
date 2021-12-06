
CREATE TABLE IF NOT EXISTS `const_change_history` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `date` DATE NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `val` VARCHAR(255) NOT NULL,
  `val_type` VARCHAR(255) NOT NULL,
  PRIMARY KEY `pk_id`(`id`),
  CONSTRAINT `name_date` UNIQUE (`name`, `date`)
)
