/** 2022.06.30 */
CREATE TABLE `comic_users`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_login` VARCHAR(60) NOT NULL DEFAULT '' COMMENT '姓名',
    `user_pass` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '用户密码',
    `mobile` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '用户手机号',
    `user_email` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '用户邮箱',
    `user_nice_name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '用户美名',
    `signature` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户个性签名',
    `avatar` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户头像',
    `avatar_thumb` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户头像缩略图',
    `user_type` tinyint(3)  unsigned DEFAULT '0' COMMENT '用户类型，0:普通用户;1:会员;2:admin',
    `source` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '注册来源',
    `private_key` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '密码密钥',
    `hardware_id` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '硬件id',
    `check_email` tinyint(1) unsigned DEFAULT '0' COMMENT '是否绑定和验证邮箱',
    `check_mobile` tinyint(1) unsigned DEFAULT '0' COMMENT '是否绑定手机',
    `user_status` tinyint(3) unsigned DEFAULT '1' COMMENT '用户状态 0：禁用； 1：正常 ；2：未验证',
    `create_at` datetime  DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
)ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='用户信息表';

CREATE TABLE `comic_options`
(
    `option_id` BIGINT NOT NULL AUTO_INCREMENT,
    `option_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '配置名',
    `option_value` LONGTEXT NOT NULL  COMMENT '配置值',
    `autoload` INT unsigned DEFAULT '1' COMMENT '是否自动加载',
    PRIMARY KEY (`option_id`) USING BTREE
)ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='配置表';

CREATE TABLE `comic_users_info_0`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL ,
    `last_login_ip` varchar(16) NOT NULL DEFAULT '' COMMENT '最后登录ip',
    `last_login_time` datetime DEFAULT CURRENT_TIMESTAMP,
    `score` INT NOT NULL COMMENT '用户积分',
    `coin` INT NOT NULL COMMENT '金币',
    `token` VARCHAR(50) NOT NULL COMMENT '授权token',
    `expire_time` INT NOT NULL COMMENT 'token到期时间',
    PRIMARY KEY (`id`) USING BTREE
)ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='用户附加表';

CREATE TABLE `comic_users_info_1`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL ,
    `last_login_ip` varchar(16) NOT NULL DEFAULT '' COMMENT '最后登录ip',
    `last_login_time` datetime DEFAULT CURRENT_TIMESTAMP,
    `score` INT NOT NULL COMMENT '用户积分',
    `coin` INT NOT NULL COMMENT '金币',
    `token` VARCHAR(50) NOT NULL COMMENT '授权token',
    `expire_time` INT NOT NULL COMMENT 'token到期时间',
    PRIMARY KEY (`id`) USING BTREE
)ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='用户附加表';

CREATE TABLE `comic_users_info_2`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL ,
    `last_login_ip` varchar(16) NOT NULL DEFAULT '' COMMENT '最后登录ip',
    `last_login_time` datetime DEFAULT CURRENT_TIMESTAMP,
    `score` INT NOT NULL COMMENT '用户积分',
    `coin` INT NOT NULL COMMENT '金币',
    `token` VARCHAR(50) NOT NULL COMMENT '授权token',
    `expire_time` INT NOT NULL COMMENT 'token到期时间',
    PRIMARY KEY (`id`) USING BTREE
)ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='用户附加表';

CREATE TABLE `comic_users_info_3`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL ,
    `last_login_ip` varchar(16) NOT NULL DEFAULT '' COMMENT '最后登录ip',
    `last_login_time` datetime DEFAULT CURRENT_TIMESTAMP,
    `score` INT NOT NULL COMMENT '用户积分',
    `coin` INT NOT NULL COMMENT '金币',
    `token` VARCHAR(50) NOT NULL COMMENT '授权token',
    `expire_time` INT NOT NULL COMMENT 'token到期时间',
    PRIMARY KEY (`id`) USING BTREE
)ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='用户附加表';