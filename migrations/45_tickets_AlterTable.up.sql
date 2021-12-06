
ALTER TABLE `tickets` ADD COLUMN `need_resolve` TINYINT NOT NULL DEFAULT(0);

UPDATE tickets SET need_resolve = 1
WHERE resolved_user_id IS NULL 
AND EXISTS (
    SELECT * FROM category_section 
    WHERE tickets.section_id = category_section.section_id 
    AND category_section.need_approval = 1
    );
