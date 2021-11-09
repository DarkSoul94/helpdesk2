
CREATE TABLE IF NOT EXISTS `filials` (
  `filial_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `region_id` INT UNSIGNED,
  `filial` VARCHAR(255),
  `ip` VARCHAR(255),
  PRIMARY KEY `pk_id`(`filial_id`),
  CONSTRAINT FOREIGN KEY (`region_id`) REFERENCES `regions` (`region_id`) ON DELETE CASCADE ON UPDATE CASCADE
)
