create table if not exists workbench.`workspaces`(
     `id`    char(6) not null comment 'workspace ID' primary key,
    `name`  varchar(20) not null comment 'workspace Name',
    `label` varchar(255) default '' not null comment 'label',
    `logo`  varchar(255) default '' not null comment 'logo',
    `description` varchar(1023) default '' not null comment 'description',

    `created_by` bigint default 0 not null comment 'creator',
    `created_at` datetime default CURRENT_TIMESTAMP not null comment 'creation time',
    `updated_by` bigint default 0 not null comment 'The last person to update the data',
    `updated_at` datetime default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP not null comment 'update time',
    `deleted_by` bigint default 0 comment 'Deleting people',
    `deleted_at` datetime default null comment 'Delete time',
    `is_deleted` tinyint default 0 not null comment 'Logical deletion 0-not deleted 1-deleted',
    INDEX (`is_deleted`)

    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci comment 'workspace model';