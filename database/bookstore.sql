/*
 Navicat Premium Data Transfer

 Source Server         : mairadb
 Source Server Type    : MySQL
 Source Server Version : 100703
 Source Host           : localhost:3306
 Source Schema         : db1

 Target Server Type    : MySQL
 Target Server Version : 100703
 File Encoding         : 65001

 Date: 01/05/2022 13:49:41
*/
CREATE DATABASE IF NOT EXISTS `bookstore_server`;
USE `bookstore_server`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_books
-- ----------------------------
DROP TABLE IF EXISTS `tb_books`;
CREATE TABLE `tb_books` (
                            `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                            `created_at` datetime(3) DEFAULT NULL,
                            `updated_at` datetime(3) DEFAULT NULL,
                            `deleted_at` datetime(3) DEFAULT NULL,
                            `title` varchar(256) NOT NULL,
                            `price` double NOT NULL,
                            `is_sell` tinyint(1) NOT NULL DEFAULT 0,
                            `number` bigint(20) unsigned NOT NULL DEFAULT 0,
                            `author` varchar(256) NOT NULL,
                            `publish_date` varchar(256) NOT NULL,
                            `category` varchar(256) NOT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_books
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for tb_cart_books
-- ----------------------------
DROP TABLE IF EXISTS `tb_cart_books`;
CREATE TABLE `tb_cart_books` (
                                 `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                 `created_at` datetime(3) DEFAULT NULL,
                                 `updated_at` datetime(3) DEFAULT NULL,
                                 `deleted_at` datetime(3) DEFAULT NULL,
                                 `cart_id` bigint(20) unsigned DEFAULT NULL,
                                 `book_id` bigint(20) unsigned NOT NULL,
                                 `price` double DEFAULT NULL,
                                 `number` bigint(20) unsigned DEFAULT NULL,
                                 `unit_price` double DEFAULT NULL,
                                 PRIMARY KEY (`id`,`book_id`),
                                 KEY `fk_tb_carts_books` (`cart_id`),
                                 CONSTRAINT `fk_tb_carts_books` FOREIGN KEY (`cart_id`) REFERENCES `tb_carts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_cart_books
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for tb_carts
-- ----------------------------
DROP TABLE IF EXISTS `tb_carts`;
CREATE TABLE `tb_carts` (
                            `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                            `created_at` datetime(3) DEFAULT NULL,
                            `updated_at` datetime(3) DEFAULT NULL,
                            `deleted_at` datetime(3) DEFAULT NULL,
                            `user_id` bigint(20) unsigned DEFAULT NULL,
                            `cart_price` double DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `fk_tb_users_cart` (`user_id`),
                            CONSTRAINT `fk_tb_users_cart` FOREIGN KEY (`user_id`) REFERENCES `tb_users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_carts
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for tb_order_books
-- ----------------------------
DROP TABLE IF EXISTS `tb_order_books`;
CREATE TABLE `tb_order_books` (
                                  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                  `created_at` datetime(3) DEFAULT NULL,
                                  `updated_at` datetime(3) DEFAULT NULL,
                                  `deleted_at` datetime(3) DEFAULT NULL,
                                  `order_id` bigint(20) unsigned DEFAULT NULL,
                                  `book_id` bigint(20) unsigned NOT NULL,
                                  `number` bigint(20) unsigned DEFAULT NULL,
                                  `unit_price` double DEFAULT NULL,
                                  PRIMARY KEY (`id`,`book_id`),
                                  KEY `fk_tb_orders_books` (`order_id`),
                                  CONSTRAINT `fk_tb_orders_books` FOREIGN KEY (`order_id`) REFERENCES `tb_orders` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_order_books
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for tb_orders
-- ----------------------------
DROP TABLE IF EXISTS `tb_orders`;
CREATE TABLE `tb_orders` (
                             `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                             `created_at` datetime(3) DEFAULT NULL,
                             `updated_at` datetime(3) DEFAULT NULL,
                             `deleted_at` datetime(3) DEFAULT NULL,
                             `user_id` bigint(20) unsigned NOT NULL,
                             `order_price` double DEFAULT NULL,
                             `is_approved` tinyint(1) NOT NULL DEFAULT 0,
                             `status` varchar(256) NOT NULL,
                             PRIMARY KEY (`id`),
                             KEY `fk_tb_users_orders` (`user_id`),
                             CONSTRAINT `fk_tb_users_orders` FOREIGN KEY (`user_id`) REFERENCES `tb_users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_orders
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for tb_users
-- ----------------------------
DROP TABLE IF EXISTS `tb_users`;
CREATE TABLE `tb_users` (
                            `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                            `created_at` datetime(3) DEFAULT NULL,
                            `updated_at` datetime(3) DEFAULT NULL,
                            `deleted_at` datetime(3) DEFAULT NULL,
                            `username` varchar(256) NOT NULL,
                            `password` varchar(256) NOT NULL,
                            `role` varchar(256) NOT NULL DEFAULT 'general',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_users
-- ----------------------------
BEGIN;
INSERT INTO `tb_users` (`id`, `username`, `password`, `role`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'admin', '$2a$10$Fv9BWzqsiQ.JuuGdcXdvN.Fx3ml.dVR47W22GoJMWQAlm9wHQIMVe', 'admin', '2021-04-18 15:40:33', '2021-04-18 15:40:33', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
