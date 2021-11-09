ALTER TABLE `shift_opening_time`
CHANGE `shift_opening_time_id` `id` int unsigned NOT NULL AUTO_INCREMENT FIRST,
CHANGE `time` `opening_time` datetime NULL AFTER `support_id`,
ADD `closing_time` datetime NULL,
ADD `closing_status` tinyint NULL DEFAULT '0',
RENAME TO `supports_shifts`;