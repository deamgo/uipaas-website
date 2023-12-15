CREATE TABLE `developer` (
                             `id` varchar(255) NOT NULL COMMENT '主键',
                             `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户名',
                             `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
                             `email` varchar(255) NOT NULL COMMENT '邮箱',
                             `password` varchar(255) NOT NULL COMMENT '密码',
                             `create_at` datetime DEFAULT NULL COMMENT '创建时间',
                             `status` int DEFAULT NULL COMMENT '状态',
                             `update_at` datetime DEFAULT NULL COMMENT '更新时间',
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;