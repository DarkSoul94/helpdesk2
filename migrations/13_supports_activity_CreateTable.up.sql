
CREATE TABLE IF NOT EXISTS `supports_activity` (
  `support_id` INT UNSIGNED,
  `ticket_id` INT UNSIGNED,
  CONSTRAINT FOREIGN KEY (`support_id`) REFERENCES `users`(`user_id`),
  CONSTRAINT FOREIGN KEY (`ticket_id`) REFERENCES `tickets`(`ticket_id`)
)
