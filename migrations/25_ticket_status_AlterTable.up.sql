ALTER TABLE `ticket_status`
CHANGE `sort_priority` `sort_priority_supp` int unsigned NULL AFTER `not_display`,
ADD `sort_priority_user` int unsigned NULL;

UPDATE `ticket_status` SET
`sort_priority_user` = '2',
`sort_priority_supp` = '1'
WHERE `ticket_status_id` = '4';

UPDATE `ticket_status` SET
`sort_priority_user` = '1',
`sort_priority_supp` = '2'
WHERE `ticket_status_id` = '6';

UPDATE `ticket_status` SET
`sort_priority_user` = '3',
`sort_priority_supp` = '3'
WHERE `ticket_status_id` = '2';

UPDATE `ticket_status` SET
`sort_priority_user` = '3',
`sort_priority_supp` = '3'
WHERE `ticket_status_id` = '3';

UPDATE `ticket_status` SET
`sort_priority_user` = '4',
`sort_priority_supp` = '3'
WHERE `ticket_status_id` NOT IN ('2', '3', '4', '6');