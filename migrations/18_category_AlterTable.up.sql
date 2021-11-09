ALTER TABLE `category`
CHANGE `price` `price` decimal(10,2) NOT NULL DEFAULT '1' AFTER `old_category`;