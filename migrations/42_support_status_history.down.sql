ALTER TABLE `support_status_history`
CHANGE `id` `support_status_history_id` int unsigned NOT NULL AUTO_INCREMENT FIRST,
ADD `prev_support_status_time` datetime NULL AFTER `support_id`,
CHANGE `select_time` `curr_support_status_time` datetime NULL AFTER `prev_support_status_time`,
ADD `prev_support_status_id` int unsigned NULL AFTER `curr_support_status_time`,
CHANGE `status_id` `curr_support_status_id` int unsigned NULL AFTER `prev_support_status_id`,
CHANGE `duration` `difference` int unsigned NULL AFTER `curr_support_status_id`;

ALTER TABLE `support_status_history`
ADD CONSTRAINT `prev_support_status_id` FOREIGN KEY (`prev_support_status_id`) REFERENCES `support_status` (`support_status_id`);