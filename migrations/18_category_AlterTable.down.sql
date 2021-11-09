ALTER TABLE `category`
CHANGE `price` `price` int unsigned NOT NULL DEFAULT '1' AFTER `old_category`;