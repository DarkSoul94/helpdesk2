
CREATE TABLE IF NOT EXISTS `user_groups` (
  `group_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `group_name` VARCHAR(255) NOT NULL,
  `create_ticket` TINYINT,
  `get_all_tickets` TINYINT,
  `see_additional_info` TINYINT,
  `can_resolve_ticket` TINYINT,
  `work_on_tickets` TINYINT,
  `change_settings` TINYINT,
  PRIMARY KEY `pk_id`(`group_id`)
);

INSERT INTO user_groups SET
group_name = "Пользователь", 
create_ticket = true, 
get_all_tickets = false, 
see_additional_info = false, 
can_resolve_ticket = false, 
work_on_tickets = false, 
change_settings = false;

INSERT INTO user_groups SET
group_name = "Администратор HelpDesk", 
create_ticket = true, 
get_all_tickets = true, 
see_additional_info = true, 
can_resolve_ticket = false, 
work_on_tickets = true, 
change_settings = true;

INSERT INTO user_groups SET
group_name = "Сотрудник тех.поддержки", 
create_ticket = false, 
get_all_tickets = false, 
see_additional_info = true, 
can_resolve_ticket = false, 
work_on_tickets = true, 
change_settings = false;

INSERT INTO user_groups SET
group_name = "Сотрудник бэк-офиса", 
create_ticket = true, 
get_all_tickets = false, 
see_additional_info = false, 
can_resolve_ticket = true, 
work_on_tickets = false, 
change_settings = false;