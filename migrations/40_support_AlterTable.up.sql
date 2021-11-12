ALTER TABLE `support`
CHANGE `accept_ticket` `status_id` int unsigned NOT NULL DEFAULT 4 AFTER `support_id`,
ADD FOREIGN KEY (`status_id`) REFERENCES `support_status` (`support_status_id`);
