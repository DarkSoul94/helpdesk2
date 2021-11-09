
CREATE TABLE IF NOT EXISTS `shifts_schedule` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `support_id` INT UNSIGNED NOT NULL,
  `office_id` INT UNSIGNED NULL,
  `date` DATE NOT NULL,
  `start_time` TIME NOT NULL,
  `end_time` TIME NOT NULL,
  `vacation` TINYINT NOT NULL DEFAULT(false),
  `sick_leave` TINYINT NOT NULL DEFAULT(false),
  PRIMARY KEY `pk_id`(`id`),
  CONSTRAINT FOREIGN KEY (`support_id`) REFERENCES `users`(`user_id`),
  CONSTRAINT FOREIGN KEY (`office_id`) REFERENCES `offices`(`id`)
);

ALTER TABLE `shifts_schedule`
ADD UNIQUE `support_date` (`support_id`, `date`);