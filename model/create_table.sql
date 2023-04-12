DROP TABLE IF EXISTS `apartment`;
CREATE TABLE `apartment` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `aid` bigint(20) NOT NULL,
                        `url` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `type` varchar(64) COLLATE utf8mb4_general_ci DEFAULT '',
                        `floor` varchar(64) COLLATE utf8mb4_general_ci DEFAULT '' ,
                        `location` varchar(64) COLLATE utf8mb4_general_ci DEFAULT '',
                        `deposit` varchar(64) COLLATE utf8mb4_general_ci DEFAULT '',
                        `room` varchar(64) COLLATE utf8mb4_general_ci DEFAULT '',
                        `area` varchar(64) COLLATE utf8mb4_general_ci DEFAULT '',
                        `subway` varchar(128) COLLATE utf8mb4_general_ci DEFAULT '',
                        `status` varchar(64) COLLATE utf8mb4_general_ci DEFAULT '',
                        `price` varchar(64) COLLATE utf8mb4_general_ci DEFAULT '',
                        `intro` longtext COLLATE utf8mb4_general_ci NOT NULL ,
                        `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;