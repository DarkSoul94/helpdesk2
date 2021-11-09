
CREATE TABLE IF NOT EXISTS `support` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `support_id` INT UNSIGNED,
  `accept_ticket` TINYINT DEFAULT 0,
  `priority` TINYINT DEFAULT 0,
  PRIMARY KEY `pk_id`(`id`),
  CONSTRAINT FOREIGN KEY (`support_id`) REFERENCES `users`(`user_id`)
)
