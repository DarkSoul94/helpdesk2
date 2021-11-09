CREATE TABLE IF NOT EXISTS `consts` (
  `name` VARCHAR(255) NOT NULL,
  `data`  TEXT NOT NULL,
  `data_type` VARCHAR(255) NOT NULL,
  `table_name` VARCHAR(255)
);

INSERT INTO `consts` (`name`, `data`, `data_type`)
VALUES ('banner', '', 'string');
