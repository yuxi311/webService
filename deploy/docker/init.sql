DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `name` varchar(255) DEFAULT NULL COMMENT 'User name',
    `username` varchar(255) DEFAULT NULL COMMENT 'User login',
    `password` varchar(255) DEFAULT NULL COMMENT 'User login password',
    `role` int DEFAULT '1' COMMENT 'User role',
    `description` varchar(255) DEFAULT NULL COMMENT 'User description',
    `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created timestamp',
    `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Updated timestamp',
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
