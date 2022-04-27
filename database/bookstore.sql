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

 Date: 28/04/2022 01:08:29
*/

CREATE DATABASE IF NOT EXISTS `bookstore_server`;
USE `bookstore_server`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_book_cart
-- ----------------------------
DROP TABLE IF EXISTS `tb_book_cart`;
CREATE TABLE `tb_book_cart` (
  `cart_id` bigint(20) unsigned NOT NULL,
  `book_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`cart_id`,`book_id`),
  KEY `fk_tb_book_cart_book` (`book_id`),
  CONSTRAINT `fk_tb_book_cart_book` FOREIGN KEY (`book_id`) REFERENCES `tb_books` (`id`),
  CONSTRAINT `fk_tb_book_cart_cart` FOREIGN KEY (`cart_id`) REFERENCES `tb_carts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_book_cart
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for tb_book_order
-- ----------------------------
DROP TABLE IF EXISTS `tb_book_order`;
CREATE TABLE `tb_book_order` (
  `order_id` bigint(20) unsigned NOT NULL,
  `book_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`order_id`,`book_id`),
  KEY `fk_tb_book_order_book` (`book_id`),
  CONSTRAINT `fk_tb_book_order_book` FOREIGN KEY (`book_id`) REFERENCES `tb_books` (`id`),
  CONSTRAINT `fk_tb_book_order_order` FOREIGN KEY (`order_id`) REFERENCES `tb_orders` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_book_order
-- ----------------------------
BEGIN;
COMMIT;

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
-- Table structure for tb_carts
-- ----------------------------
DROP TABLE IF EXISTS `tb_carts`;
CREATE TABLE `tb_carts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  `cart_price` double NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_carts
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
  `order_price` double NOT NULL,
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
