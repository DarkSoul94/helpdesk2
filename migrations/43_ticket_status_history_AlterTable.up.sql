ALTER TABLE `ticket_status_history`
DROP FOREIGN KEY `ticket_status_history_ibfk_2`;

ALTER TABLE `ticket_status_history`
DROP `prev_status_time`,
DROP `prev_status_id`,
CHANGE `ticket_status_history_id` `id` int unsigned NOT NULL AUTO_INCREMENT FIRST,
CHANGE `curr_status_time` `select_time` datetime NULL AFTER `changed_user_id`,
CHANGE `curr_status_id` `status_id` int unsigned NULL AFTER `select_time`,
CHANGE `difference` `duration` int unsigned NULL AFTER `status_id`;