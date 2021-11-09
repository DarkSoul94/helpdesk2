
CREATE TABLE IF NOT EXISTS `users` (
  `user_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(255) NOT NULL,
  `user_name` VARCHAR(255) NOT NULL,
  `group_id` INT UNSIGNED NOT NULL DEFAULT '1',
  `department` VARCHAR(255),
  PRIMARY KEY `pk_id`(`user_id`),
  CONSTRAINT FOREIGN KEY (`group_id`) REFERENCES `user_groups`(`group_id`)
);

INSERT users(user_name, email) VALUES("Система распределения", " ");
INSERT users(user_name, email) VALUES("Сотрудник тех. поддержки", " ");
INSERT users(user_name, email) VALUES("Сотрудник бэк. офиса", " ");