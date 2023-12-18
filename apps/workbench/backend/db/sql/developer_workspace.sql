create table if not exists workbench.`workspace_developer_relation`(
   `developer_id` bigint not null comment '开发者ID',
   `workspace_id` char(6) not null comment '工作空间ID',
    `role` tinyint not null comment '角色  1-Admin 2-Developer 3-Reviewer',

    `created_by` bigint default 0 not null comment '创建人',
    `created_at` datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_by` bigint default 0 not null comment '最后一次更新的开发者',
    `updated_at` datetime default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP not null comment '最后一次更新时间',
    `deleted_by` bigint default 0 comment '删除人',
    `deleted_at` datetime default null comment '删除时间',
    `is_deleted` tinyint default 0 not null comment '逻辑删除  0-未删除 1-已删除',
    INDEX (`is_deleted`),

    PRIMARY KEY (`developer_id`,`workspace_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci comment '开发者与工作空间关系数据模型';