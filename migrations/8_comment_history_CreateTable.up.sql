
CREATE TABLE IF NOT EXISTS `comment_history` (
  `comment_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `ticket_id` INT UNSIGNED,
  `comment_date` DATETIME,
  `comment_author_id` INT UNSIGNED,
  `comment_text` TEXT,
  PRIMARY KEY `pk_id`(`comment_id`),
  CONSTRAINT FOREIGN KEY (`ticket_id`) REFERENCES `tickets`(`ticket_id`),
  CONSTRAINT FOREIGN KEY (`comment_author_id`) REFERENCES `users`(`user_id`)
)
