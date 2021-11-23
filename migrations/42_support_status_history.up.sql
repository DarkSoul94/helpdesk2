ALTER TABLE `support_status_history`
DROP FOREIGN KEY `support_status_history_ibfk_1`;

ALTER TABLE `support_status_history`
CHANGE `support_status_history_id` `id` int unsigned NOT NULL AUTO_INCREMENT FIRST,
DROP `prev_support_status_time`,
CHANGE `curr_support_status_time` `select_time` datetime NULL AFTER `support_id`,
DROP `prev_support_status_id`,
CHANGE `curr_support_status_id` `status_id` int unsigned NULL AFTER `select_time`,
CHANGE `difference` `duration` int unsigned NULL AFTER `status_id`;