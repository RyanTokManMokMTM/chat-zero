
CREATE TABLE `user` (
 `id` bigint not null AUTO_INCREMENT,
 `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `del_state` tinyint NOT NULL NOT NULL DEFAULT 0,
 `version` bigint NOT NULL DEFAULT 0,
 `email` varchar(64) CHARACTER SET utf8mb4 collate utf8_general_ci NOT NULL,
 `password` varchar(255) CHARACTER SET utf8mb4 collate utf8_general_ci NOT NULL,
 `name` varchar(255) CHARACTER SET utf8mb4 collate utf8_general_ci NOT NULL,
 `set` tinyint(1) NOT NULL DEFAULT 0,
 `avatar` varchar(255) CHARACTER SET utf8mb4 collate utf8_general_ci NOT NULL DEFAULT '',
 `info` varchar(255) CHARACTER SET utf8mb4 collate utf8_general_ci NOT NULL DEFAULT '',
 primary key(`id`),
 UNIQUE KEY `idx_email`(`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;