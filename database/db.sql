CREATE DATABASE IF NOT EXISTS `bookstore_server` DEFAULT CHARACTER SET utf8;

USE `bookstore_server`;

DROP TABLE IF EXISTS `tb_users`;

CREATE TABLE `tb_users` (
                            `id`  bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                            `username` varchar(255) NOT NULL,
                            `password` varchar(255) NOT NULL,
                            `role`     varchar(255) NOT NULL,
                            `createdAt` timestamp NULL DEFAULT NULL,
                            `updatedAt` timestamp NULL DEFAULT NULL,
                            `deletedAt` timestamp NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `username` (`username`),
                            KEY `idx_tb_users_deletedAt` (`deletedAt`)
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

LOCK TABLES `tb_users` WRITE;

INSERT INTO `tb_users` VALUES (1,'admin','$2a$10$Fv9BWzqsiQ.JuuGdcXdvN.Fx3ml.dVR47W22GoJMWQAlm9wHQIMVe','admin','2021-04-18 15:40:33','2021-04-18 15:40:33',NULL);

UNLOCK TABLES;

DROP TABLE IF EXISTS `tb_shops`;

CREATE TABLE `tb_shops` (
                            `id`  bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                            `name` varchar(255) NOT NULL,
                            `owner` bigint(20) unsigned NOT NULL,
                            `createdAt` timestamp NULL DEFAULT NULL,
                            `updatedAt` timestamp NULL DEFAULT NULL,
                            `deletedAt` timestamp NULL DEFAULT NULL,
                            CONSTRAINT `tb_users` FOREIGN KEY (`owner`) REFERENCES tb_users(`id`),
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `name` (`name`),
                            KEY `idx_tb_shops_deletedAt` (`deletedAt`)
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `tb_books`;

CREATE TABLE `tb_books` (
                            `id`  bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                            `title` varchar(255) NOT NULL,
                            `price` double NOT NULL,
                            `category` varchar(255) NOT NULL,
                            `shopID` bigint(20) unsigned NOT NULL,
                            `isSell` bool NOT NULL DEFAULT FALSE,
                            `number` int(11) unsigned NOT NULL,
                            `publishDate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `createdAt` timestamp NULL DEFAULT NULL,
                            `updatedAt` timestamp NULL DEFAULT NULL,
                            `deletedAt` timestamp NULL DEFAULT NULL,
                            CONSTRAINT `tb_shops` FOREIGN KEY (`shopID`) REFERENCES tb_shops(`id`),
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `username` (`title`),
                            KEY `idx_tb_books_deletedAt` (`deletedAt`)
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
