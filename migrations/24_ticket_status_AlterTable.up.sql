ALTER TABLE `ticket_status`
ADD `sort_priority` int unsigned NULL;

UPDATE `ticket_status` SET
`sort_priority` = '1'
WHERE `ticket_status_id` = '4';

UPDATE `ticket_status` SET
`sort_priority` = '1'
WHERE `ticket_status_id` = '6';

UPDATE `ticket_status` SET
`sort_priority` = '2'
WHERE `ticket_status_id` = '2';

UPDATE `ticket_status` SET
`sort_priority` = '2'
WHERE `ticket_status_id` = '3';

UPDATE `ticket_status` SET
`sort_priority` = '3'
WHERE `ticket_status_id` NOT IN ('2', '3', '4', '6');