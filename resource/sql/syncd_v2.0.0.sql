# ************************************************************
# Copyright 2019 syncd Author. All Rights Reserved.
# Use of this source code is governed by a MIT-style
# license that can be found in the LICENSE file.
# ************************************************************

create database `core` default charset utf8mb4;

use `core`;

DROP TABLE IF EXISTS `tb_core_user`;

CREATE TABLE `tb_core_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) unsigned NOT NULL DEFAULT '0',
  `username` varchar(20) NOT NULL DEFAULT '',
  `password` char(32) NOT NULL DEFAULT '',
  `salt` char(10) NOT NULL DEFAULT '',
  `truename` varchar(10) NOT NULL DEFAULT '',
  `mobile` varchar(20) NOT NULL DEFAULT '',
  `email` varchar(500) NOT NULL DEFAULT '',
  `status` int(11) unsigned NOT NULL DEFAULT '0',
  `last_login_time` int(11) unsigned NOT NULL DEFAULT '0',
  `last_login_ip` varchar(50) NOT NULL DEFAULT '',
  `ctime` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_username` (`username`),
  KEY `idx_email` (`email`(20))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `tb_core_user` (`id`, `role_id`, `username`, `password`, `salt`, `truename`, `mobile`, `email`, `status`, `last_login_time`, `last_login_ip`, `ctime`)
VALUES
	(1,1,'syncd','c2a8d572366f7cf7bfc8b485f41bfd1b','u0EMxuE6qh','Syncd','','',1,0,'',0);

DROP TABLE IF EXISTS `tb_core_user_role`;

CREATE TABLE `tb_core_user_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `privilege` varchar(2000) NOT NULL DEFAULT '',
  `ctime` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `tb_core_user_role` (`id`, `name`, `privilege`, `ctime`)
VALUES
	(1,'管理员','2001,2002,2003,2004,2100,2101,2102,2201,2202,2203,2204,2205,2206,2207,3001,3002,3004,3003,3101,3102,3103,3104,4001,4002,4003,4004,4101,4102,4103,4104,1001,1002,1006,1003,1004,1005',0);

DROP TABLE IF EXISTS `tb_core_user_token`;

CREATE TABLE `tb_core_user_token` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `token` varchar(100) NOT NULL DEFAULT '',
  `expire` int(11) unsigned NOT NULL DEFAULT '0',
  `ctime` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;