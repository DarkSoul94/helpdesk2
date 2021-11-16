ALTER TABLE `support`
DROP FOREIGN KEY `support_ibfk_2`,
ADD `id` int unsigned NOT NULL AUTO_INCREMENT UNIQUE FIRST,
CHANGE `status_id` `accept_ticket` tinyint NOT NULL DEFAULT '0' AFTER `support_id`;

UPDATE `support`
SET `accept_ticket` = DEFAULT;

ALTER TABLE `support`
ADD INDEX `support_id` (`support_id`),
ADD PRIMARY KEY `PRIMARY` (`id`),
DROP INDEX `PRIMARY`,
DROP INDEX `status_id`,
DROP INDEX `id`;