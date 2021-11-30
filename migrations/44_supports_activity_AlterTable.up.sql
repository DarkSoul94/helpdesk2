
ALTER TABLE `supports_activity`
ADD `reassignment` tinyint unsigned NOT NULL DEFAULT false,
ADD UNIQUE `activity` (`support_id`, `ticket_id`);
