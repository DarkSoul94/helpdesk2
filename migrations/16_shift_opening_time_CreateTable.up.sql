
CREATE TABLE IF NOT EXISTS `shift_opening_time` (
  `shift_opening_time_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `support_id` INT UNSIGNED,
  `time` DATETIME,
  PRIMARY KEY `pk_id`(`shift_opening_time_id`),
  CONSTRAINT FOREIGN KEY (`support_id`) REFERENCES `users`(`user_id`)
)
