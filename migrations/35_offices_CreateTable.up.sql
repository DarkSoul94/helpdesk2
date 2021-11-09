
CREATE TABLE IF NOT EXISTS `offices` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `color` VARCHAR(255) NOT NULL,
  `deleted` TINYINT NOT NULL DEFAULT(false),
  PRIMARY KEY `pk_id`(`id`)
);

INSERT INTO `offices` SET
`name` = "Главный офис",
`color` = "#FFFFFF",
`deleted` = false;