
-- Create syntax for TABLE 'server_disks'
-- 磁盘路径 关联表
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
-- 服务器用户信息 关联表
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
  `server_id` varchar(36) NOT NULL DEFAULT '' COMMENT '生成的唯一id字符串',
  `tag` varchar(100) NOT NULL DEFAULT '' COMMENT '名称/标签',
  `env` varchar(30) DEFAULT '' COMMENT '环境',
  `engine_room` varchar(30) NOT NULL DEFAULT '' COMMENT '机房',
  `core` int(4) unsigned NOT NULL COMMENT '核心数',
  `memory` int(4) unsigned NOT NULL COMMENT '内存数(G）',
  `intranet_ip` varchar(30) DEFAULT '' COMMENT '内网ip',
  `extranet_ip` varchar(30) DEFAULT '' COMMENT '外网ip',
  `remark` varchar(255) DEFAULT '',
  `is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '逻辑删除字段',
  `creator_id` bigint(20) unsigned DEFAULT NULL COMMENT '创建人',
  `updater_id` bigint(20) unsigned COMMENT '修改人',
  `created_at` bigint(20) unsigned NOT NULL,
  `updated_at` bigint(20) unsigned DEFAULT NULL ,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Create syntax for TABLE 'service'
CREATE TABLE `services` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `tag` varchar(100) NOT NULL DEFAULT '' COMMENT '名称/标签',
  `application` varchar(100) NOT NULL DEFAULT '' COMMENT '应用',
  `url` varchar(200) DEFAULT '' COMMENT 'url',
  `version` varchar(50) NOT NULL DEFAULT '' COMMENT 'version',
  `dependency` varchar(200) DEFAULT '' COMMENT '依赖',
  `remark` varchar(255) DEFAULT '',
  `creator_id` bigint(20) unsigned DEFAULT NULL COMMENT '创建人',
  `updater_id` bigint(20) unsigned DEFAULT NULL COMMENT '最后操作人',
  `is_deleted` tinyint(1) unsigned NOT NULL,
  `created_at` bigint(20) unsigned NOT NULL,
  `updated_at` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Create syntax for TABLE 'server_servings'
CREATE TABLE `server_servings` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `serving_id` bigint(20) unsigned NOT NULL COMMENT '服务id',
  `server_id` bigint(20) unsigned DEFAULT NULL COMMENT '服务器id',
  `env` varchar(30) NOT NULL DEFAULT '' COMMENT '环境-服务器',
  `application` varchar(100) NOT NULL DEFAULT '' COMMENT '应用',
  `engine_room` varchar(30) NOT NULL DEFAULT '' COMMENT '机房',
  `intranet_ip` varchar(30) DEFAULT '' COMMENT '内网ip',
  `extranet_ip` varchar(30) DEFAULT '' COMMENT '外网ip',
  `creator_id` bigint(20) unsigned DEFAULT NULL COMMENT '创建人',
  `updater_id` bigint(20) unsigned DEFAULT NULL COMMENT '最后操作人',
  `is_deleted` tinyint(1) unsigned NOT NULL,
  `created_at` bigint(20) unsigned NOT NULL,
  `updated_at` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



