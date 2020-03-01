/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80018
 Source Host           : localhost:3306
 Source Schema         : kronos

 Target Server Type    : MySQL
 Target Server Version : 80018
 File Encoding         : 65001

 Date: 01/03/2020 18:21:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for kr_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `kr_admin_role`;
CREATE TABLE `kr_admin_role` (
  `admin_id` bigint(20) unsigned NOT NULL,
  `roles_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`admin_id`,`roles_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of kr_admin_role
-- ----------------------------
BEGIN;
INSERT INTO `kr_admin_role` VALUES (1, 1);
INSERT INTO `kr_admin_role` VALUES (1, 2);
COMMIT;

-- ----------------------------
-- Table structure for kr_admins
-- ----------------------------
DROP TABLE IF EXISTS `kr_admins`;
CREATE TABLE `kr_admins` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `username` char(50) COLLATE utf8mb4_bin NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `last_login_ip` int(1) NOT NULL,
  `is_super` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_kr_admins_username` (`username`),
  KEY `idx_kr_admins_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of kr_admins
-- ----------------------------
BEGIN;
INSERT INTO `kr_admins` VALUES (1, '2020-03-01 11:40:50', '2020-03-01 18:20:21', NULL, 'admin', '$2a$10$IPLwWXs2rIwcGewKBWcGJ.3zqTqNfXTUeIEb1rZb533xN1P.K31Ou', 0, 1);
COMMIT;

-- ----------------------------
-- Table structure for kr_articles
-- ----------------------------
DROP TABLE IF EXISTS `kr_articles`;
CREATE TABLE `kr_articles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `title` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `keyword` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `description` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `thumb` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_kr_articles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for kr_casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `kr_casbin_rule`;
CREATE TABLE `kr_casbin_rule` (
  `p_type` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v1` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v2` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v3` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of kr_casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `kr_casbin_rule` VALUES ('p', '测试', '/admin/admins/lists', 'get', '', '', '');
INSERT INTO `kr_casbin_rule` VALUES ('p', '测试', '/admin/role/lists', 'get', '', '', '');
INSERT INTO `kr_casbin_rule` VALUES ('p', '测试', '/admin/admins/edit', 'get', '', '', '');
INSERT INTO `kr_casbin_rule` VALUES ('p', '测试', '/admin/permission/lists', 'get', '', '', '');
INSERT INTO `kr_casbin_rule` VALUES ('p', '测试', '/admin/config/index', 'get', '', '', '');
INSERT INTO `kr_casbin_rule` VALUES ('p', '测试', '/admin/article/index', 'get', '', '', '');
INSERT INTO `kr_casbin_rule` VALUES ('g', 'admin', '管理员', '', '', '', '');
INSERT INTO `kr_casbin_rule` VALUES ('g', 'admin', '测试', '', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for kr_permissions
-- ----------------------------
DROP TABLE IF EXISTS `kr_permissions`;
CREATE TABLE `kr_permissions` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL,
  `description` char(64) COLLATE utf8mb4_bin DEFAULT NULL,
  `slug` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL,
  `http_path` text COLLATE utf8mb4_bin,
  `method` char(10) COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_kr_permissions_title` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of kr_permissions
-- ----------------------------
BEGIN;
INSERT INTO `kr_permissions` VALUES (1, '管理员', '1', '1', '/admin/admins/lists', 'get');
INSERT INTO `kr_permissions` VALUES (2, '角色', '1', '1', '/admin/role/lists', 'get');
INSERT INTO `kr_permissions` VALUES (3, '管理员编辑', '1', '1', '/admin/admins/edit', 'get');
INSERT INTO `kr_permissions` VALUES (4, '权限', '1', '1', '/admin/permission/lists', 'get');
INSERT INTO `kr_permissions` VALUES (5, '系统配置', '1', '1', '/admin/config/index', 'get');
INSERT INTO `kr_permissions` VALUES (6, '文章', '1', '1', '/admin/article/index', 'get');
COMMIT;

-- ----------------------------
-- Table structure for kr_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `kr_role_menu`;
CREATE TABLE `kr_role_menu` (
  `roles_id` bigint(20) unsigned NOT NULL,
  `permissions_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`roles_id`,`permissions_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of kr_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `kr_role_menu` VALUES (2, 1);
INSERT INTO `kr_role_menu` VALUES (2, 2);
INSERT INTO `kr_role_menu` VALUES (2, 3);
INSERT INTO `kr_role_menu` VALUES (2, 4);
INSERT INTO `kr_role_menu` VALUES (2, 5);
INSERT INTO `kr_role_menu` VALUES (2, 6);
COMMIT;

-- ----------------------------
-- Table structure for kr_roles
-- ----------------------------
DROP TABLE IF EXISTS `kr_roles`;
CREATE TABLE `kr_roles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL,
  `description` char(64) COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_kr_roles_title` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of kr_roles
-- ----------------------------
BEGIN;
INSERT INTO `kr_roles` VALUES (1, '管理员', '');
INSERT INTO `kr_roles` VALUES (2, '测试', '123123');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
