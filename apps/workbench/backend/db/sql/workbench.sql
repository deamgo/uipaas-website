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
                               `description` varchar(255) DEFAULT NULL COMMENT 'app 描述',
                               `icon` varchar(255) DEFAULT NULL COMMENT 'app 图标',
                               `status` tinyint NOT NULL COMMENT 'app 状态',
                               `created_by` bigint NOT NULL DEFAULT '0' COMMENT '创建人',
                               `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '最后一次更新的开发者',
                               `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
                               `deleted_by` bigint DEFAULT '0' COMMENT '删除人',
                               `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                               `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '逻辑删除  0-未删除 1-已删除',
                               PRIMARY KEY (`id`),
                               KEY `is_deleted_index` (`is_deleted`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for developer
-- ----------------------------
DROP TABLE IF EXISTS `developer`;
CREATE TABLE `developer` (
                             `id` varchar(255) NOT NULL COMMENT '主键',
                             `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户名',
                             `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
                             `email` varchar(255) NOT NULL COMMENT '邮箱',
                             `password` varchar(255) NOT NULL COMMENT '密码',
                             `create_at` datetime DEFAULT NULL COMMENT '创建时间',
                             `status` int DEFAULT NULL COMMENT '状态',
                             `update_at` datetime DEFAULT NULL COMMENT '更新时间',
                             PRIMARY KEY (`id`) USING BTREE,
                             UNIQUE KEY `username` (`username`),
                             UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for workspace_developer_relation
-- ----------------------------
DROP TABLE IF EXISTS `workspace_developer_relation`;
CREATE TABLE `workspace_developer_relation` (
                                                `developer_id` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '开发者ID',
                                                `workspace_id` char(6) NOT NULL COMMENT '工作空间ID',
                                                `email` varchar(50) NOT NULL COMMENT '开发者邮箱',
                                                `role` tinyint NOT NULL COMMENT '角色  1-Admin 2-Developer 3-Reviewer',
                                                `status` tinyint DEFAULT NULL COMMENT '邀请状态 0-Pending 1-Accept',
                                                `created_by` bigint NOT NULL DEFAULT '0' COMMENT '创建人',
                                                `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                                `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '最后一次更新的开发者',
                                                `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
                                                `deleted_by` bigint DEFAULT '0' COMMENT '删除人',
                                                `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                                                `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '逻辑删除  0-未删除 1-已删除',
                                                KEY `is_deleted` (`is_deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='开发者与工作空间关系数据模型';

-- ----------------------------
-- Table structure for workspaces
-- ----------------------------
DROP TABLE IF EXISTS `workspaces`;
CREATE TABLE `workspaces` (
                              `id` char(6) NOT NULL COMMENT '工作空间ID',
                              `name` varchar(20) NOT NULL COMMENT '工作空间名称 访问标识',
                              `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '短描述',
                              `logo` varchar(255) NOT NULL COMMENT '图标地址',
                              `description` varchar(1023) NOT NULL DEFAULT '' COMMENT '长描述',
                              `created_by` bigint NOT NULL DEFAULT '0' COMMENT '创建人',
                              `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '最后一次更新的开发者',
                              `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
                              `deleted_by` bigint DEFAULT '0' COMMENT '删除人',
                              `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                              `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '逻辑删除  0-未删除 1-已删除',
                              PRIMARY KEY (`id`),
                              KEY `is_deleted` (`is_deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='工作空间数据模型';

SET FOREIGN_KEY_CHECKS = 1;
