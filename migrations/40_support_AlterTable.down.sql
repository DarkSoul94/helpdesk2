ALTER TABLE `support`
CHANGE `status_id` `accept_ticket` TINYINT DEFAULT 0 AFTER `support_id`,
DROP INDEX `status_id`;