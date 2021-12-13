ALTER TABLE `support`
CHANGE `accept_ticket` `status_id` int unsigned NOT NULL DEFAULT '4' AFTER `support_id`;

UPDATE `support`
SET `status_id` = DEFAULT;

ALTER TABLE `support`
DROP `id`,
DROP INDEX `support_id`,
ADD CONSTRAINT `status_id` FOREIGN KEY (`status_id`) REFERENCES `support_status` (`support_status_id`),
ADD PRIMARY KEY `support_id` (`support_id`);
