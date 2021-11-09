
CREATE TABLE IF NOT EXISTS `ticket_status_history` (
  `ticket_status_history_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `ticket_id` INT UNSIGNED,
  `changed_user_id` INT UNSIGNED,
  `prev_status_time` DATETIME,
  `curr_status_time` DATETIME,
  `prev_status_id` INT UNSIGNED,
  `curr_status_id` INT UNSIGNED,
  `difference` INT UNSIGNED,
  PRIMARY KEY `pk_id`(`ticket_status_history_id`),
  CONSTRAINT FOREIGN KEY (`ticket_id`) REFERENCES `tickets`(`ticket_id`),
  CONSTRAINT FOREIGN KEY (`prev_status_id`) REFERENCES `ticket_status`(`ticket_status_id`),
  CONSTRAINT FOREIGN KEY (`curr_status_id`) REFERENCES `ticket_status`(`ticket_status_id`),
  CONSTRAINT FOREIGN KEY (`changed_user_id`) REFERENCES `users`(`user_id`)
)
