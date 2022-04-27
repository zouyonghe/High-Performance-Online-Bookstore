/*
 Navicat Premium Data Transfer

 Source Server         : mairadb
 Source Server Type    : MySQL
 Source Server Version : 100703
 Source Host           : localhost:3306
 Source Schema         : bookstore_server

 Target Server Type    : MySQL
 Target Server Version : 100703
 File Encoding         : 65001

 Date: 24/04/2022 23:27:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_books
-- ----------------------------
DROP TABLE IF EXISTS `tb_books`;
CREATE TABLE `tb_books`(
                           `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                           `title`       varchar(255)        NOT NULL,
                           `author`      varchar(255)        NOT NULL,
                           `price`       double              NOT NULL,
                           `category`    varchar(255)        NOT NULL,
                           `isSell`      tinyint(1)          NOT NULL DEFAULT 0,
                           `number`      int(11) unsigned    NOT NULL,
                           `publishDate` datetime            NOT NULL DEFAULT current_timestamp(),
                           `createdAt`   timestamp           NULL     DEFAULT NULL,
                           `updatedAt`   timestamp           NULL     DEFAULT NULL,
                           `deletedAt`   timestamp           NULL     DEFAULT NULL,
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `username` (`title`),
                           KEY `idx_tb_books_deletedAt` (`deletedAt`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of tb_books
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for tb_users
-- ----------------------------
DROP TABLE IF EXISTS `tb_users`;
CREATE TABLE `tb_users`
(
    `id`        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `username`  varchar(255)        NOT NULL,
    `password`  varchar(255)        NOT NULL,
    `role`      varchar(255)        NOT NULL,
    `createdAt` timestamp           NULL DEFAULT NULL,
    `updatedAt` timestamp           NULL DEFAULT NULL,
    `deletedAt` timestamp           NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`),
    KEY `idx_tb_users_deletedAt` (`deletedAt`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of tb_users
-- ----------------------------
BEGIN;
INSERT INTO `tb_users` (`id`, `username`, `password`, `role`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (1, 'admin', '$2a$10$Fv9BWzqsiQ.JuuGdcXdvN.Fx3ml.dVR47W22GoJMWQAlm9wHQIMVe', 'admin', '2021-04-18 15:40:33', '2021-04-18 15:40:33', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
