
CREATE TABLE IF NOT EXISTS `approval_bindings` (
  `group_id` INT UNSIGNED NOT NULL,
  `section_id` INT UNSIGNED NOT NULL,
  CONSTRAINT FOREIGN KEY (`group_id`) REFERENCES `user_groups`(`group_id`),
  CONSTRAINT FOREIGN KEY (`section_id`) REFERENCES `category_section`(`section_id`),
  CONSTRAINT `approval` UNIQUE (`group_id`, `section_id`)
)
