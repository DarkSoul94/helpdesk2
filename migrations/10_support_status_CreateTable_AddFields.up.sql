CREATE TABLE IF NOT EXISTS `support_status` (
  `support_status_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `support_status_name` VARCHAR(255),
  PRIMARY KEY `pk_id`(`support_status_id`)
);

INSERT support_status(support_status_name) VALUES("Принимаю запросы");
INSERT support_status(support_status_name) VALUES("Перерыв");
INSERT support_status(support_status_name) VALUES("Работа в офисе");
INSERT support_status(support_status_name) VALUES("Не принимаю запросы");
