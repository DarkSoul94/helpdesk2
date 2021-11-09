CREATE TABLE IF NOT EXISTS `support_lateness` (
  `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `date` datetime NOT NULL,
  `support_id` int unsigned NOT NULL,
  `cause` text NOT NULL,
  `decision` tinyint NULL,
  `difference` INT UNSIGNED,
  FOREIGN KEY (`support_id`) REFERENCES `users` (`user_id`)
);