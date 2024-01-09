/*
 Navicat Premium Data Transfer

 Source Server         : db
 Source Server Type    : MySQL
 Source Server Version : 80028 (8.0.28)
 Source Host           : localhost:3306
 Source Schema         : workbench

 Target Server Type    : MySQL
 Target Server Version : 80028 (8.0.28)
 File Encoding         : 65001

 Date: 02/01/2024 11:45:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for application
-- ----------------------------
DROP TABLE IF EXISTS `application`;
CREATE TABLE `application` (
                               `id` varchar(25) NOT NULL,
                               `workspace_id` char(6) NOT NULL COMMENT 'workspaceID',
                               `name` varchar(50) NOT NULL COMMENT 'app name',
                               `description` varchar(255) DEFAULT NULL COMMENT 'app description',
                               `icon` varchar(255) DEFAULT NULL COMMENT 'app icon',
                               `status` tinyint NOT NULL COMMENT 'app status 0-Draft 1-Published',
                               `created_by` bigint NOT NULL DEFAULT '0',
                               `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
                               `updated_by` bigint NOT NULL DEFAULT '0',
                               `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
                               `deleted_by` bigint DEFAULT '0',
                               `deleted_at` datetime DEFAULT NULL ,
                               `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT 'Logical deletion  0-NotDeleted 1-Deleted',
                               PRIMARY KEY (`id`),
                               KEY `is_deleted_index` (`is_deleted`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for developer
-- ----------------------------
DROP TABLE IF EXISTS `developer`;
CREATE TABLE `developer` (
                             `id` varchar(255) NOT NULL COMMENT 'PRIMARY KEY',
                             `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'username',
                             `avatar` varchar(255) DEFAULT NULL COMMENT 'avatar',
                             `email` varchar(255) NOT NULL COMMENT 'email',
                             `password` varchar(255) NOT NULL COMMENT 'password',
                             `create_at` datetime DEFAULT NULL ,
                             `status` int DEFAULT NULL ,
                             `update_at` datetime DEFAULT NULL,
                             PRIMARY KEY (`id`) USING BTREE,
                             UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for workspace_developer_relation
-- ----------------------------
DROP TABLE IF EXISTS `workspace_developer_relation`;
CREATE TABLE `workspace_developer_relation` (
                                                `developer_id` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'ID',
                                                `workspace_id` char(6) NOT NULL COMMENT 'workspace_id',
                                                `email` varchar(50) NOT NULL COMMENT 'email',
                                                `role` tinyint NOT NULL COMMENT 'role  1-Admin 2-Developer 3-Reviewer',
                                                `status` tinyint DEFAULT NULL COMMENT 'Invitation status 0-Pending 1-Accept',
                                                `created_by` bigint NOT NULL DEFAULT '0' ,
                                                `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
                                                `updated_by` bigint NOT NULL DEFAULT '0' ,
                                                `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                                `deleted_by` bigint DEFAULT '0' ,
                                                `deleted_at` datetime DEFAULT NULL ,
                                                `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT 'Logical deletion  0-NotDeleted 1-Deleted',
                                                KEY `is_deleted` (`is_deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Developer and workspace relationship data model';

-- ----------------------------
-- Table structure for workspaces
-- ----------------------------
DROP TABLE IF EXISTS `workspaces`;
CREATE TABLE `workspaces` (
                              `id` char(6) NOT NULL COMMENT 'ID',
                              `name` varchar(20) NOT NULL COMMENT 'name',
                              `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'label',
                              `logo` varchar(255)  COMMENT 'logo path',
                              `description` varchar(1023) NOT NULL DEFAULT '' COMMENT 'description',
                              `created_by` bigint NOT NULL DEFAULT '0' ,
                              `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              `updated_by` bigint NOT NULL DEFAULT '0' ,
                              `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
                              `deleted_by` bigint DEFAULT '0' ,
                              `deleted_at` datetime DEFAULT NULL ,
                              `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT 'Logical deletion  0-NotDeleted 1-Deleted',
                              PRIMARY KEY (`id`),
                              KEY `is_deleted` (`is_deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Workspace data model';

SET FOREIGN_KEY_CHECKS = 1;
