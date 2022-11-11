CREATE TABLE `task_log` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `task_id` int NOT NULL DEFAULT '0' COMMENT '任务ID',
  `name` varchar(32) NOT NULL COMMENT '任务名称',
  `spec` varchar(64) NOT NULL COMMENT 'crontab',
  `command` varchar(255) NOT NULL COMMENT '命令',
  `timeout` mediumint(9) NOT NULL DEFAULT '0' COMMENT '任务执行超时时间',
  `retry_times` tinyint NOT NULL DEFAULT '0' COMMENT '任务重试次数',
  `start_time` datetime DEFAULT NULL COMMENT '开始执行时间',
  `end_time` datetime DEFAULT NULL COMMENT '执行结束时间',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '10:运行中，20:完成，30:失败',
  `result` mediumtext NOT NULL,
  `created` datetime NOT NULL COMMENT '添加时间',
  `updated` datetime NOT NULL COMMENT '更新时间',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否被删除',
  PRIMARY KEY (`id`),
  KEY (`task_id`),
  KEY (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='任务日志表';

CREATE TABLE `task` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(32) COLLATE utf8mb4_bin NOT NULL COMMENT '任务名称',
  `spec` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT 'crontab',
  `command` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '命令',
  `timeout` mediumint(9) NOT NULL DEFAULT '0' COMMENT '任务执行超时时间',
  `retry_times` tinyint(4) NOT NULL DEFAULT '0' COMMENT '任务重试次数',
  `retry_interval` smallint(6) NOT NULL DEFAULT '0' COMMENT '重试间隔',
  `remark` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '备注',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '10正常20停止',
  `created` datetime NOT NULL COMMENT '添加时间',
  `updated` datetime NOT NULL COMMENT '更新时间',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否被删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='任务表';

INSERT INTO `auth`(`id`, `app_key`, `app_secret`, `created`, `updated`) VALUES (1, 'eddycjy', 'go-programming-tour-book' , NOW(), NOW());

CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL COMMENT '账号',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `salt` varchar(20) NOT NULL COMMENT "加密盐",
  `created` datetime NOT NULL COMMENT '添加时间',
  `updated` datetime NOT NULL COMMENT '更新时间',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否被删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';