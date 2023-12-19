create table if not exists workbench.`workspace_developer_relation`(
   `developer_id` bigint not null comment 'DeveloperID',
   `workspace_id` char(6) not null comment 'WorkspaceID',
    `role` tinyint not null comment 'role  1-Admin 2-Developer 3-Reviewer',

    `created_by` bigint default 0 not null comment 'creator',
    `created_at` datetime default CURRENT_TIMESTAMP not null comment 'create time',
    `updated_by` bigint default 0 not null comment 'The last person to update the data',
    `updated_at` datetime default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP not null comment 'update time',
    `deleted_by` bigint default 0 comment 'Deleting people',
    `deleted_at` datetime default null comment 'Delete time',
    `is_deleted` tinyint default 0 not null comment 'Logical deletion 0-not deleted 1-deleted',
    INDEX (`is_deleted`),

    PRIMARY KEY (`developer_id`,`workspace_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci comment 'workspace_developer';