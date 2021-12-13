
ALTER TABLE `support_status`
ADD `accept_ticket` tinyint NOT NULL DEFAULT '0';

UPDATE `support_status` SET
`accept_ticket` = true
WHERE `support_status_id` = '1';

UPDATE `support_status` SET
`accept_ticket` = false
WHERE `support_status_id` = '2';

UPDATE `support_status` SET
`accept_ticket` = false
WHERE `support_status_id` = '3';

UPDATE `support_status` SET
`accept_ticket` = false
WHERE `support_status_id` = '4';

