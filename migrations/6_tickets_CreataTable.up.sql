
CREATE TABLE IF NOT EXISTS `tickets` (
  `ticket_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `ticket_date` DATETIME,
  `section_id` INT UNSIGNED,
  `ticket_text` TEXT,
  `ticket_status_id` INT UNSIGNED,
  `filial` VARCHAR(255),
  `ip` VARCHAR(255),
  `ticket_author_id` INT UNSIGNED,
  `support_id` INT UNSIGNED,
  `resolved_user_id` INT UNSIGNED,
  `service_comment` TEXT,
  `ticket_grade` INT UNSIGNED,
  PRIMARY KEY `pk_id`(`ticket_id`),
  CONSTRAINT FOREIGN KEY (`section_id`) REFERENCES `category_section`(`section_id`),
  CONSTRAINT FOREIGN KEY (`ticket_status_id`) REFERENCES `ticket_status`(`ticket_status_id`),
  CONSTRAINT FOREIGN KEY (`ticket_author_id`) REFERENCES `users`(`user_id`),
  CONSTRAINT FOREIGN KEY (`support_id`) REFERENCES `users`(`user_id`),
  CONSTRAINT FOREIGN KEY (`resolved_user_id`) REFERENCES `users`(`user_id`)
)
