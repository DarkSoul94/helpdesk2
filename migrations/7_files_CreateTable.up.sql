
CREATE TABLE IF NOT EXISTS `files` (
  `file_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `ticket_id` INT UNSIGNED,
  `file_name` VARCHAR(255),
  `file_extension` VARCHAR(255),
  `file_date` DATETIME,
  `file_data` LONGBLOB,
  PRIMARY KEY `pk_id`(`file_id`),
  CONSTRAINT FOREIGN KEY (`ticket_id`) REFERENCES `tickets`(`ticket_id`)
)
