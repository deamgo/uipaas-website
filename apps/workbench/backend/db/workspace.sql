create table if not exists workbench.`workspace`(
    `id`    char(6) not null comment '工作空间ID' primary key,
    `name`  varchar(20) not null comment '工作空间名称 访问标识',
    `logo`  varchar(255) not null comment '图标地址',
    `lable` varchar(255) default '' not null comment '短描述',
    `description` varchar(1023) default '' not null comment '长描述',
    `created_by` bigint default 0 not null comment '创建人',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_by` bigint default 0 not null comment '最后一次更新的开发者',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP not null comment '最后一次更新时间',
    `deleted_by` bigint default 0 comment '删除人',
    `deleted_at` datetime default null comment '删除时间',
    `is_deleted` tinyint default 0 not null comment '逻辑删除  0-未删除 1-已删除',
    INDEX (`is_deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci comment '工作空间数据模型';
