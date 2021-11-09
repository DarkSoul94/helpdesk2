
CREATE TABLE IF NOT EXISTS `ticket_status` (
  `ticket_status_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `ticket_status_name` VARCHAR(255),
  `not_display` TINYINT NULL DEFAULT false,
  PRIMARY KEY `pk_id`(`ticket_status_id`)
);

INSERT INTO ticket_status SET
ticket_status_name = "Новый",
not_display = true;

INSERT INTO ticket_status SET
ticket_status_name = "В ожидании",
not_display = true;

INSERT INTO ticket_status SET
ticket_status_name = "В ожидании согласования",
not_display = true;

INSERT INTO ticket_status SET
ticket_status_name = "В работе",
not_display = false;

INSERT INTO ticket_status SET
ticket_status_name = "В процессе реализации",
not_display = false;

INSERT INTO ticket_status SET
ticket_status_name = "Отправлен на доработку",
not_display = false;

INSERT INTO ticket_status SET
ticket_status_name = "Отложен",
not_display = false;

INSERT INTO ticket_status SET
ticket_status_name = "Отклонен",
not_display = false;

INSERT INTO ticket_status SET
ticket_status_name = "Выполнен",
not_display = false;