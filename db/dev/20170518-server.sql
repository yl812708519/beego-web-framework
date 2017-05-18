-- Create syntax for TABLE 'server_disks'
CREATE TABLE `server_disks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `server_id` bigint(20) unsigned NOT NULL,
  `root_path` varchar(200) NOT NULL DEFAULT '' COMMENT '根路径',
  `size` int(8) unsigned NOT NULL COMMENT '容量',
  `is_deleted` tinyint(1) unsigned NOT NULL,
  `created_at` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create syntax for TABLE 'server_users'
CREATE TABLE `server_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `server_id` bigint(20) unsigned NOT NULL,
  `user_name` varchar(100) NOT NULL DEFAULT '',
  `password` varchar(100) NOT NULL DEFAULT '',
  `is_deleted` tinyint(1) unsigned NOT NULL,
  `created_at` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create syntax for TABLE 'servers'
CREATE TABLE `servers` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `server_id` varchar(20) NOT NULL DEFAULT '' COMMENT '生成的唯一id字符串',
  `tag` varchar(100) NOT NULL DEFAULT '' COMMENT '名称/标签',
  `application` varchar(100) NOT NULL DEFAULT '' COMMENT '应用',
  `env` varchar(30) NOT NULL DEFAULT '' COMMENT '环境',
  `engine_room` varchar(30) NOT NULL DEFAULT '' COMMENT '机房',
  `core` tinyint(4) unsigned NOT NULL COMMENT '核心数',
  `memory` tinyint(4) unsigned NOT NULL COMMENT '内存数(G）',
  `intranet_ip` varchar(30) NOT NULL DEFAULT '' COMMENT '内网ip',
  `extranet_ip` varchar(30) NOT NULL DEFAULT '' COMMENT '外网ip',
  `remark` varchar(255) NOT NULL DEFAULT '',
  `is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '逻辑删除字段',
  `creator_id` bigint(20) unsigned NOT NULL COMMENT '创建人',
  `created_at` bigint(20) unsigned NOT NULL,
  `updated_at` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Create syntax for TABLE 'service'
CREATE TABLE `service` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `tag` varchar(100) NOT NULL DEFAULT '' COMMENT '名称/标签',
  `application` varchar(100) NOT NULL DEFAULT '' COMMENT '应用',
  `server_id` bigint(20) unsigned NOT NULL COMMENT '服务器id',
  `url` varchar(200) NOT NULL DEFAULT '' COMMENT 'url',
  `version` varchar(50) NOT NULL DEFAULT '' COMMENT 'version',
  `dependency` varchar(200) NOT NULL DEFAULT '' COMMENT '依赖',
  `env` varchar(100) NOT NULL DEFAULT '' COMMENT '环境',
  `remark` varchar(255) NOT NULL DEFAULT '',
  `updater_id` bigint(20) unsigned DEFAULT NULL COMMENT '最后操作人',
  `is_deleted` tinyint(1) unsigned NOT NULL,
  `created_at` bigint(20) unsigned NOT NULL,
  `updated_at` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;