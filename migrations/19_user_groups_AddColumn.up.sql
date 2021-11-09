ALTER TABLE `user_groups`
ADD `full_search` tinyint NULL DEFAULT '0';

UPDATE user_groups SET
full_search = true
WHERE group_name = 'Администратор HelpDesk' OR group_name = 'Сотрудник тех.поддержки';