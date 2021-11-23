ALTER TABLE `ticket_status_history`
CHANGE `id` `ticket_status_history_id` int unsigned NOT NULL AUTO_INCREMENT FIRST,
ADD `prev_status_time` datetime NULL AFTER `changed_user_id`,
CHANGE `select_time` `curr_status_time` datetime NULL AFTER `prev_status_time`,
ADD `prev_status_id` int unsigned NULL AFTER `curr_status_time`,
CHANGE `status_id` `curr_status_id` int unsigned NULL AFTER `prev_status_id`,
CHANGE `duration` `difference` int unsigned NULL AFTER `curr_status_id`;

ALTER TABLE `ticket_status_history`
ADD CONSTRAINT `ticket_status_history_ibfk_2` FOREIGN KEY (`prev_status_id`) REFERENCES `ticket_status` (`ticket_status_id`);
