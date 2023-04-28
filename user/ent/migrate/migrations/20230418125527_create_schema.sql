-- Create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `first_name` varchar(100) NOT NULL, `last_name` varchar(100) NOT NULL, `age` bigint NOT NULL DEFAULT 20, `address` varchar(300) NOT NULL, `email` varchar(255) NULL, `group_id` bigint NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
