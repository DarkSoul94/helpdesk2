CREATE TABLE IF NOT EXISTS `user_groups` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `permissions` JSON NOT NULL,
  PRIMARY KEY `pk_id`(`id`)
);

INSERT INTO `user_groups` SET
`name` = "Regular user",
`permissions` = "{}";

INSERT INTO `user_groups` SET
`name` = "Admin",
`permissions` = "{}";
