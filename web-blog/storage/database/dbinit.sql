/*
 Navicat MySQL Data Transfer

 Source Server         : 本地数据库
 Source Server Type    : MySQL
 Source Server Version : 50727
 Source Host           : 127.0.0.1:3306
 Source Schema         : webblogdatabase

 Date: 24/12/2021 23:29:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article_xiazhaoxuans
-- ----------------------------
DROP TABLE IF EXISTS `article_xiazhaoxuans`;
CREATE TABLE `article_xiazhaoxuans` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `user_id` int(10) unsigned DEFAULT NULL,
  `content` text,
  `view` int(11) DEFAULT '0',
  `comment` int(11) DEFAULT '0',
  `is_top` tinyint(3) unsigned DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `title` (`title`),
  KEY `user_id` (`user_id`),
  KEY `idx_articles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for users_xiazhaoxuans
-- ----------------------------
DROP TABLE IF EXISTS `users_xiazhaoxuans`;
CREATE TABLE `users_xiazhaoxuans` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `lockstate` tinyint(1) DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `autograph` varchar(255) DEFAULT NULL,
  `introduction` varchar(255) DEFAULT NULL,
  `personal_website` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
