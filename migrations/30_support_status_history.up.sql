ALTER TABLE `support_status_history`
ADD `shift_id` int unsigned NULL,
ADD FOREIGN KEY (`shift_id`) REFERENCES `supports_shifts` (`id`);

UPDATE support_status_history AS History
INNER JOIN supports_shifts AS Shifts ON DATE(Shifts.opening_time) = DATE(History.curr_support_status_time)
SET History.shift_id = IF(History.support_id = Shifts.support_id, Shifts.id, NULL)
WHERE History.support_id = Shifts.support_id;