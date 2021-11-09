ALTER TABLE `supports_shifts`
CHANGE `id` `shift_opening_time_id` int unsigned NOT NULL AUTO_INCREMENT FIRST,
CHANGE `opening_time` `time` datetime NULL AFTER `support_id`,
DROP `closing_time`,
DROP `closing_status`,
RENAME TO `shift_opening_time`;