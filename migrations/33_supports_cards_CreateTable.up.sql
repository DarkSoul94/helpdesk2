CREATE TABLE IF NOT EXISTS `supports_cards` (
  `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `support_id` int unsigned NOT NULL,
  `internal_number` varchar(255) NOT NULL DEFAULT '',
  `mobile_number` varchar(255) NOT NULL DEFAULT '',
  `birth_date` varchar(255) NULL,
  `is_senior` tinyint NOT NULL DEFAULT '0',
  `senior_id` int unsigned NULL,
  `wager` 	decimal(10,2) NOT NULL DEFAULT '0',
  `comment` text NULL,
  `color` text NULL,
  FOREIGN KEY (`support_id`) REFERENCES `users` (`user_id`),
  FOREIGN KEY (`senior_id`) REFERENCES `users` (`user_id`));