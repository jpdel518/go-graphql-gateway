-- Create "groups" table
CREATE TABLE `groups` (`id` bigint NOT NULL AUTO_INCREMENT, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `name` varchar(100) NOT NULL, `description` varchar(100) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
