CREATE TABLE IF NOT EXISTS `support_status_history` (
  `support_status_history_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `support_id` INT UNSIGNED,
  `prev_support_status_time` DATETIME,
  `curr_support_status_time` DATETIME,
  `prev_support_status_id` INT UNSIGNED,
  `curr_support_status_id` INT UNSIGNED,
  `difference` INT UNSIGNED,
  PRIMARY KEY `pk_id`(`support_status_history_id`),
  CONSTRAINT FOREIGN KEY (`prev_support_status_id`) REFERENCES `support_status`(`support_status_id`),
  CONSTRAINT FOREIGN KEY (`curr_support_status_id`) REFERENCES `support_status`(`support_status_id`),
  CONSTRAINT FOREIGN KEY (`support_id`) REFERENCES `users`(`user_id`)
)
