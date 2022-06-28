/** 2021.10.20 */

DROP TABLE `cmf_agt_bank`;

CREATE TABLE `cmf_agt_bank_0`
(
    `id`             int(11) NOT NULL AUTO_INCREMENT,
    `name`           varchar(255) NOT NULL DEFAULT '' COMMENT '开户人姓名',
    `alipay_account` varchar(255) NOT NULL DEFAULT '',
    `bank`           varchar(255) NOT NULL DEFAULT '',
    `branch`         varchar(255) NOT NULL DEFAULT '',
    `bank_account`   varchar(255) NOT NULL DEFAULT '',
    `alipay_name`    varchar(255) NOT NULL DEFAULT '',
    `btc_account`    varchar(100) NOT NULL DEFAULT '' COMMENT '虚拟币账号',
    `pay_type`       tinyint(3) unsigned DEFAULT '0' COMMENT '0 银行卡 1支付宝，默认选择类型',
    `user_id`        int(11) NOT NULL,
    `status`         tinyint(3) unsigned DEFAULT '0' COMMENT '0 允许修改 1不允许修改',
    `remark`         varchar(255) NOT NULL DEFAULT '',
    `create_at`      datetime              DEFAULT CURRENT_TIMESTAMP,
    `update_at`      datetime              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `tablenum`       int(11) DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=106 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='代理提款信息表';

CREATE TABLE `cmf_agt_bank_1`
(
    `id`             int(11) NOT NULL AUTO_INCREMENT,
    `name`           varchar(255) NOT NULL DEFAULT '' COMMENT '开户人姓名',
    `alipay_account` varchar(255) NOT NULL DEFAULT '',
    `bank`           varchar(255) NOT NULL DEFAULT '',
    `branch`         varchar(255) NOT NULL DEFAULT '',
    `bank_account`   varchar(255) NOT NULL DEFAULT '',
    `alipay_name`    varchar(255) NOT NULL DEFAULT '',
    `btc_account`    varchar(100) NOT NULL DEFAULT '' COMMENT '虚拟币账号',
    `pay_type`       tinyint(3) unsigned DEFAULT '0' COMMENT '0 银行卡 1支付宝，默认选择类型',
    `user_id`        int(11) NOT NULL,
    `status`         tinyint(3) unsigned DEFAULT '0' COMMENT '0 允许修改 1不允许修改',
    `remark`         varchar(255) NOT NULL DEFAULT '',
    `create_at`      datetime              DEFAULT CURRENT_TIMESTAMP,
    `update_at`      datetime              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `tablenum`       int(11) DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=106 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='代理提款信息表';

CREATE TABLE `cmf_agt_bank_2`
(
    `id`             int(11) NOT NULL AUTO_INCREMENT,
    `name`           varchar(255) NOT NULL DEFAULT '' COMMENT '开户人姓名',
    `alipay_account` varchar(255) NOT NULL DEFAULT '',
    `bank`           varchar(255) NOT NULL DEFAULT '',
    `branch`         varchar(255) NOT NULL DEFAULT '',
    `bank_account`   varchar(255) NOT NULL DEFAULT '',
    `alipay_name`    varchar(255) NOT NULL DEFAULT '',
    `btc_account`    varchar(100) NOT NULL DEFAULT '' COMMENT '虚拟币账号',
    `pay_type`       tinyint(3) unsigned DEFAULT '0' COMMENT '0 银行卡 1支付宝，默认选择类型',
    `user_id`        int(11) NOT NULL,
    `status`         tinyint(3) unsigned DEFAULT '0' COMMENT '0 允许修改 1不允许修改',
    `remark`         varchar(255) NOT NULL DEFAULT '',
    `create_at`      datetime              DEFAULT CURRENT_TIMESTAMP,
    `update_at`      datetime              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `tablenum`       int(11) DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=106 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='代理提款信息表';


/** 2021.10.21 */
CREATE TABLE `cmf_activity`
(
    `id`        int(12) NOT NULL AUTO_INCREMENT,
    `tag`       varchar(50)  NOT NULL DEFAULT '0' COMMENT '标签',
    `des`       varchar(200) NOT NULL DEFAULT '' COMMENT '描述',
    `url`       varchar(200) NOT NULL DEFAULT '' COMMENT '链接',
    `thumb`     varchar(200) NOT NULL DEFAULT '' COMMENT '图片链接',
    `orderno`   int(11) NOT NULL DEFAULT '0' COMMENT '序号',
    `disabled`  tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否隐藏，0显示，1隐藏',
    `create_at` datetime              DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) COMMENT '活动';


/** 2021.10.25 */
ALTER TABLE `cmf_users`
    ADD COLUMN `hardware_id` varchar(100) NOT NULL DEFAULT '' COMMENT '硬件ID',
ADD INDEX(`hardware_id`);


/** 2021.11.01 */
ALTER TABLE `cmf_users`
    ADD COLUMN `hardware_id` varchar(100) NOT NULL DEFAULT '' COMMENT '硬件ID',
ADD INDEX(`user_email`),
ADD INDEX(`mobile`);

ALTER TABLE `cmf_users`
    ADD INDEX(`user_type`);

ALTER TABLE `cmf_users_vip`
    ADD INDEX(`uid`);

ALTER TABLE `cmf_album`
    ADD COLUMN `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0精选1频道';

/** 2021.11.02 */
ALTER TABLE `cmf_users_coinrecord`
    ADD INDEX(`addtime`);
ALTER TABLE `cmf_users_vip`
    ADD INDEX(`endtime`);
ALTER TABLE `cmf_users`
    ADD INDEX(`agt_level_id`),
ADD INDEX(`iszombiep`),
ADD INDEX(`iszombie`),
ADD INDEX(`issuper`),
ADD INDEX(`source`),
ADD INDEX(`create_time`),
ADD INDEX(`ishot`);

/** 2021.11.04 */
ALTER TABLE `cmf_users`
    ADD COLUMN `news_vip` tinyint(1) NOT NULL DEFAULT 0 COMMENT '线下(动态)VIP';

CREATE TABLE `cmf_news_type`
(
    `id`             int(11) NOT NULL AUTO_INCREMENT,
    `type_name`      varchar(100) NOT NULL,
    `seq_no`         int(11) NOT NULL DEFAULT '999' COMMENT '排序正序',
    `need_vip`       tinyint(1) NOT NULL DEFAULT '0' COMMENT '需要VIP',
    `need_video_vip` tinyint(1) NOT NULL DEFAULT '0' COMMENT '需要点播VIP',
    `discount`       tinyint(4) NOT NULL DEFAULT '100' COMMENT 'VIP打折',
    `create_at`      datetime DEFAULT CURRENT_TIMESTAMP,
    `update_at`      datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `disabled`       tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否隐藏',
    PRIMARY KEY (`id`) USING BTREE
) COMMENT='动态分类表';

ALTER TABLE `cmf_news`
    ADD COLUMN `type_id` int(11) NOT NULL DEFAULT 0 COMMENT '类型ID';


ALTER TABLE `cmf_agt_level`
    ADD COLUMN `extract_news_vip` int(11) NOT NULL DEFAULT '0' COMMENT '动态开通VIP提成',
ADD COLUMN `percent_news_vip` int(11) NOT NULL DEFAULT '0' COMMENT '动态开通VIP扣点';

ALTER TABLE `cmf_news_buy_0`
    ADD COLUMN `is_free` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是免费';
ALTER TABLE `cmf_news_buy_1`
    ADD COLUMN `is_free` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是免费';
ALTER TABLE `cmf_news_buy_2`
    ADD COLUMN `is_free` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是免费';

/** 2021.11.23 */
ALTER TABLE `cmf_users`
    ADD COLUMN `video_reply` tinyint(1) NOT NULL DEFAULT 0 COMMENT '视频评论权限1开启0关闭';
ALTER TABLE `cmf_charge_rules`
    ADD COLUMN `video_reply` tinyint(1) NOT NULL DEFAULT 0 COMMENT '视频评论权限';
ADD COLUMN `news_vip` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否赠送动态会员';

/** 2021.11.25 */
ALTER TABLE `cmf_video`
DROP
INDEX `idx_isVip`,
DROP
INDEX `idx_point`,
DROP
INDEX `idx_type`,
DROP
INDEX `idx_vpt`,
ADD INDEX `idx_vpt`(`type_id`, `is_vip`, `point`),
ADD INDEX `idx_vip`(`is_vip`, `status`, `disabled`, `create_at`),
ADD INDEX `idx_point`(`status`, `disabled`, `point`, `create_at`),
ADD INDEX `idx_point_vip`(`is_vip`, `status`, `disabled`, `point`, `create_at`),
ADD INDEX `idx_type`(`type_id`, `is_vip`, `status`, `disabled`),
ADD INDEX `idx_status`(`status`, `disabled`, `create_at`),
ADD INDEX `idx_type_status`(`type_id`, `status`, `disabled`, `create_at`);

ALTER TABLE `thinkcmf`.`cmf_video_count`
    ADD INDEX(`see_count`),
ADD INDEX(`collection_count`),
ADD INDEX(`like_count`);


/** 2021.11.28 */
CREATE TABLE `cmf_video_order`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `video_id`  int(11) NOT NULL,
    `type`      int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'G点预告:10000,重磅热推:10001,人气排行:10002,小编私藏:10003,今日上新:10004',
    `seq_no`    int(11) NOT NULL DEFAULT '999',
    `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `position`  tinyint(3) NOT NULL DEFAULT '0' COMMENT '位置，填1~8',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_type_vid` (`video_id`,`type`),
    KEY         `type` (`type`),
    KEY         `idx_vid` (`video_id`)
) COMMENT='首页视频排序表';

/** 2021.11.29 */
ALTER TABLE `cmf_users_charge_admin`
    CHANGE COLUMN `coin` `detail` varchar (200) NOT NULL DEFAULT '' COMMENT '产品详情';

/** 2021.12.1 */
ALTER TABLE `cmf_chat_room`
    ADD INDEX(`roomid`);
ALTER TABLE `cmf_chat_user_room`
    ADD INDEX(`uid`);

/** 2021.12.2 */
ALTER TABLE `cmf_users`
    ADD INDEX `dx_t_source`(`create_time`, `source`);

/** 2021.12.3 */
CREATE TABLE `cmf_log_statistics`
(
    `id`                  int(11) NOT NULL AUTO_INCREMENT,
    `date`                date           NOT NULL,
    `register`            int(11) NOT NULL DEFAULT '0' COMMENT '注册人数',
    `promoter`            int(11) NOT NULL DEFAULT '0' COMMENT '推广人数',
    `vip`                 int(11) NOT NULL DEFAULT '0' COMMENT '新增VIP',
    `promoter_vip`        int(11) NOT NULL DEFAULT '0' COMMENT '新增推广VIP',
    `money`               decimal(11, 2) NOT NULL DEFAULT '0.00' COMMENT '业绩',
    `promoter_money`      decimal(11, 2) NOT NULL DEFAULT '0.00' COMMENT '代理业绩',
    `user_money`          decimal(11, 2) NOT NULL DEFAULT '0.00' COMMENT '普通用户业绩',
    `percentage`          decimal(11, 2) NOT NULL DEFAULT '0.00' COMMENT '盈利',
    `promoter_percentage` decimal(11, 2) NOT NULL DEFAULT '0.00' COMMENT '代理盈利',
    `user_percentage`     decimal(11, 2) NOT NULL DEFAULT '0.00' COMMENT '普通用户盈利',
    `create_at`           datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at`           datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE,
    KEY                   `date` (`date`) USING BTREE
) COMMENT='统计记录';

ALTER TABLE `cmf_payment_charge_rules`
    ADD COLUMN `chance` tinyint(4) NOT NULL DEFAULT 0 COMMENT '概率0~100';

ALTER TABLE `cmf_users_voterecord`
    ADD COLUMN `before` decimal(20, 2) NOT NULL DEFAULT 0 COMMENT '行为前余额',
ADD COLUMN `after` decimal(20, 2) NOT NULL DEFAULT 0 COMMENT '行为后余额';

CREATE TABLE `cmf_agt_control`
(
    `id`              int(11) NOT NULL AUTO_INCREMENT,
    `agent_id`        int(11) NOT NULL DEFAULT '0',
    `agent_parent_id` int(11) DEFAULT '0' COMMENT '父级代理id',
    `user_id`         int(11) NOT NULL DEFAULT '0',
    `is_child`        tinyint(4) DEFAULT '0' COMMENT '是否为子代理',
    `type`            tinyint(3) NOT NULL DEFAULT '0' COMMENT '1首充,2续费,3视频购买,4线下购买,5礼物,7线下VIP',
    `money`           decimal(11, 2) NOT NULL DEFAULT '0.00',
    `per`             int(11) NOT NULL DEFAULT '0',
    `percentage`      decimal(11, 2) NOT NULL DEFAULT '0.00',
    `agent_level`     int(11) NOT NULL DEFAULT '0' COMMENT '代理等级',
    `remark`          varchar(500)            DEFAULT NULL,
    `create_at`       datetime                DEFAULT CURRENT_TIMESTAMP,
    `update_at`       datetime                DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE,
    KEY               `agent_id` (`agent_id`) USING BTREE,
    KEY               `user_id` (`user_id`) USING BTREE
) COMMENT='代理控制记录表';


/** 2021.12.6 */
ALTER TABLE `cmf_payment_charge_rules`
    ADD COLUMN `channel` tinyint(4) NOT NULL DEFAULT 0 COMMENT '渠道 0全渠道(默认) 1 PC 2 安卓 3 苹果';

CREATE TABLE `cmf_log_statistics_user_online`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `date`      date       NOT NULL,
    `time`      varchar(5) NOT NULL,
    `count`     int(11) NOT NULL DEFAULT '0' COMMENT '人数',
    `create_at` datetime   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE,
    KEY         `date` (`date`) USING BTREE
) COMMENT='实时用户数量记录';

/** 2021.12.16 */
ALTER TABLE `cmf_menu`
    MODIFY COLUMN `action` char (30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '操作名称';

/** 2022.01.10 */
ALTER TABLE `cmf_news`
    ADD COLUMN `hot` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '热度',
ADD COLUMN `collection_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '收藏数',
ADD COLUMN `view_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '查看数';

ALTER TABLE `cmf_news`
    ADD INDEX(`collection_count`),
ADD INDEX(`hot`),
ADD INDEX(`create_at`);


/** 2022.01.12 */
CREATE TABLE `cmf_news_tag`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `name`      varchar(100) NOT NULL COMMENT '名字',
    `disabled`  tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0显示 1不显示',
    `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `seq_no`    int(11) NOT NULL DEFAULT '0' COMMENT '排序',
    PRIMARY KEY (`id`) USING BTREE
) COMMENT='动态标签表';

ALTER TABLE `cmf_news`
    ADD COLUMN `tag` varchar(500) NOT NULL DEFAULT '' COMMENT '标签',
ADD COLUMN `key1st` varchar(200) NOT NULL DEFAULT '',
ADD COLUMN `key2nd` varchar(200) NOT NULL DEFAULT '',
ADD COLUMN `key3rd` varchar(200) NOT NULL DEFAULT '',
ADD COLUMN `key4th` varchar(200) NOT NULL DEFAULT '',
ADD COLUMN `key5th` varchar(200) NOT NULL DEFAULT '',
ADD COLUMN `key6th` varchar(200) NOT NULL DEFAULT '';

/** 2022.01.13 */
CREATE TABLE `cmf_news_denounce`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `uid`       int(11) NOT NULL,
    `news_id`   int(11) NOT NULL,
    `content`   VARCHAR(500) NOT NULL DEFAULT '',
    `create_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `status`    tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态0未处理1已经处理',
    `type`      tinyint(4) NOT NULL DEFAULT 0 COMMENT '类型，1虚假资源，2货不对板，3服务敷衍，4欺诈收费',
    PRIMARY KEY (`id`) USING BTREE,
    KEY         `uid` (`uid`) USING BTREE,
    KEY         `news_id` (`news_id`) USING BTREE
) COMMENT='动态举报表';


/** 2022.04.17 */
ALTER TABLE `cmf_charge_rules`
    ADD COLUMN `raw_price` decimal(11, 2) NOT NULL DEFAULT 0 COMMENT '原价',
ADD COLUMN `buy_discount` tinyint(4) UNSIGNED NOT NULL DEFAULT 100 COMMENT '购片折扣，100不打折，0免费',
ADD COLUMN `unlimited_download` tinyint(1) NOT NULL DEFAULT 0 COMMENT '无限下载，1是0否',
ADD COLUMN `cdn` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'cdn权限',
ADD COLUMN `level` tinyint(4) NOT NULL DEFAULT 0 COMMENT '标注等级';

ALTER TABLE `cmf_users_vip`
    ADD COLUMN `level` tinyint(4) NOT NULL DEFAULT 0 COMMENT 'VIP等级';


/** 2022.04.27 */
CREATE TABLE `cmf_order_success_rate_statistics`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `date`      date       NOT NULL COMMENT '日期',
    `hour`      tinyint(4) NOT NULL COMMENT '小时',
    `order`     int(11) NOT NULL COMMENT '订单数',
    `success`   int(11) NOT NULL COMMENT '成功数',
    `rate`      varchar(7) NOT NULL DEFAULT '-' COMMENT '成功率',
    `create_at` datetime            DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime            DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE,
    KEY         `hour` (`date`,`hour`)
) COMMENT='订单成功率统计表';


/** 2022.04.28 */
ALTER TABLE `cmf_agt_level`
    ADD COLUMN `card_discount` tinyint(4) NOT NULL DEFAULT 100 COMMENT '卡密折扣,0~100';

ALTER TABLE `cmf_users_info_0`
    ADD COLUMN `card_money` decimal(11, 2) NOT NULL DEFAULT 0 COMMENT '卡密余额';
ALTER TABLE `cmf_users_info_1`
    ADD COLUMN `card_money` decimal(11, 2) NOT NULL DEFAULT 0 COMMENT '卡密余额';
ALTER TABLE `cmf_users_info_2`
    ADD COLUMN `card_money` decimal(11, 2) NOT NULL DEFAULT 0 COMMENT '卡密余额';

/** 2022.05.04 */
CREATE TABLE `cmf_card`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT,
    `agent_id`    int(11) NOT NULL COMMENT '代理ID',
    `card_no`     varchar(32) NOT NULL COMMENT '卡密序列号',
    `goods_id`    int(11) NOT NULL DEFAULT '0' COMMENT '商品ID',
    `card_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '卡密状态 0：未使用，1：已使用',
    `usage_time`  datetime DEFAULT NULL COMMENT '使用时间',
    `create_at`   datetime DEFAULT CURRENT_TIMESTAMP,
    `update_at`   datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_card_no` (`card_no`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='卡密表';

CREATE TABLE `cmf_card_log`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `user_id`   int(11) NOT NULL,
    `card_no`   varchar(32) NOT NULL COMMENT '卡密序列号',
    `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE,
    KEY         `idx_user_id` (`user_id`) USING BTREE,
    KEY         `idx_card_no` (`card_no`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='卡密使用记录';

/** 2022.04.29 */
CREATE TABLE `cmf_agt_statistic`
(
    `agent_id`            int(11) NOT NULL COMMENT '代理ID',
    `personal_sales`      decimal(11, 2) NOT NULL DEFAULT '0.00' COMMENT '个人总业绩',
    `sub_sales`           decimal(11, 2) NOT NULL DEFAULT '0.00' COMMENT '子代理总业绩',
    `sales`               decimal(11, 2) NOT NULL DEFAULT '0.00' COMMENT '总业绩',
    `sub_count`           int(11) NOT NULL DEFAULT '0' COMMENT '子代理总数',
    `first_charge_profit` decimal(11, 2) NOT NULL DEFAULT '0.00' COMMENT '首充总盈利',
    `renewal_profit`      decimal(11, 2) NOT NULL DEFAULT '0.00' COMMENT '续费总盈利',
    `sub_commission`      decimal(11, 2) NOT NULL DEFAULT '0.00' COMMENT '子代理总提成',
    `profit`              decimal(11, 2) NOT NULL DEFAULT '0.00' COMMENT '盈利总计',
    `user_count`          int(11) NOT NULL DEFAULT '0.00' COMMENT '总用户数',
    `first_charge_count`  int(11) NOT NULL DEFAULT '0.00' COMMENT '总首充数',
    `renewal_count`       int(11) NOT NULL DEFAULT '0.00' COMMENT '总续费数',
    `vip_count`           int(11) NOT NULL DEFAULT '0.00' COMMENT '总vip数',
    `create_at`           datetime                DEFAULT CURRENT_TIMESTAMP,
    `update_at`           datetime                DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`agent_id`) USING BTREE
) COMMENT='代理统计表';


ALTER TABLE `cmf_users_agent`
    ADD INDEX(`one_uid`),
ADD INDEX(`two_uid`);


/** 2022.05.01 */
ALTER TABLE `cmf_users`
    ADD COLUMN `auth_code` varchar(255) NULL DEFAULT NULL COMMENT '二次验证密钥',
ADD COLUMN `is_auth` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否开启二次验证';

ALTER TABLE `cmf_users_live`
    MODIFY COLUMN `pull` varchar (500) NOT NULL DEFAULT '' COMMENT '播流地址',
    ADD COLUMN `isgather` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是采集';

/** 2022.05.04 */
ALTER TABLE `cmf_users`
    MODIFY COLUMN `avatar` varchar (255) NOT NULL DEFAULT '' COMMENT '用户头像',
    MODIFY COLUMN `avatar_thumb` varchar (255) NOT NULL DEFAULT '' COMMENT '用户头像缩略图';

/** 2022.05.05 */
ALTER TABLE `cmf_agt_drow`
    ADD COLUMN `detail` varchar(500) DEFAULT '' COMMENT '提现详情';


/** 2022.05.08 */
ALTER TABLE `cmf_gift`
    ADD COLUMN `screen_effect` varchar(255) NOT NULL DEFAULT '' COMMENT '屏幕效果';

ALTER TABLE `cmf_agt_statistic`
    ADD COLUMN `live_first_charge_profit` decimal(11, 2) NOT NULL DEFAULT 0.00 COMMENT '直播首充总盈利',
ADD COLUMN `live_renewal_profit` decimal(11, 2) NOT NULL DEFAULT 0.00 COMMENT '直播续费总盈利',
ADD COLUMN `live_vip_count` int(11) NOT NULL DEFAULT 0 COMMENT '直播总vip数',
ADD COLUMN `live_first_charge_count` int(11) NOT NULL DEFAULT 0 COMMENT '直播总首充数',
ADD COLUMN `live_renewal_count` int(11) NOT NULL DEFAULT 0 COMMENT '直播总续费数';

ALTER TABLE `cmf_agt_drow`
    ADD COLUMN `live_first_charge` decimal(18, 2) NOT NULL DEFAULT 0.00 COMMENT '直播首充提成',
ADD COLUMN `live_charge` decimal(18, 2) NOT NULL DEFAULT 0.00 COMMENT '直播续费提成';

ALTER TABLE `cmf_agt_level`
    ADD COLUMN `extract_live_vip_first` int(11) NOT NULL DEFAULT 0 COMMENT '直播首充提成',
ADD COLUMN `extract_live_vip_renew` int(11) NOT NULL DEFAULT 0 COMMENT '直播续费提成',
ADD COLUMN `percent_live_vip_first` int(11) NOT NULL DEFAULT 0 COMMENT '直播首充扣点',
ADD COLUMN `percent_live_vip_renew` int(11) NOT NULL DEFAULT 0 COMMENT '直播续费扣点';


ALTER TABLE `cmf_charge_rules`
    MODIFY COLUMN `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '1.钻石、2.VIP、3.LIVE VIP';

CREATE TABLE `cmf_users_live_vip`
(
    `id`      int(10) NOT NULL AUTO_INCREMENT,
    `uid`     int(10) NOT NULL DEFAULT '0' COMMENT '用户ID',
    `addtime` int(10) NOT NULL DEFAULT '0' COMMENT '添加时间',
    `endtime` int(10) NOT NULL DEFAULT '0' COMMENT '到期时间',
    `level`   tinyint(4) NOT NULL DEFAULT '0' COMMENT 'VIP等级',
    PRIMARY KEY (`id`) USING BTREE,
    KEY       `uid` (`uid`),
    KEY       `endtime` (`endtime`)
) COMMENT "用户直播VIP表";

ALTER TABLE `cmf_vip_log_0`
    ADD INDEX(`uid`);
ALTER TABLE `cmf_vip_log_1`
    ADD INDEX(`uid`);
ALTER TABLE `cmf_vip_log_2`
    ADD INDEX(`uid`);

CREATE TABLE `cmf_live_vip_log_0`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `uid`       int(11) NOT NULL,
    `name`      varchar(255)   NOT NULL DEFAULT '' COMMENT '名称',
    `price`     decimal(11, 2) NOT NULL DEFAULT '0.00' COMMENT '价格',
    `create_at` datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `update_at` datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE,
    KEY         `uid` (`uid`)
) COMMENT '直播VIP记录表';

CREATE TABLE `cmf_live_vip_log_1`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `uid`       int(11) NOT NULL,
    `name`      varchar(255)   NOT NULL DEFAULT '' COMMENT '名称',
    `price`     decimal(11, 2) NOT NULL DEFAULT '0.00' COMMENT '价格',
    `create_at` datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `update_at` datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE,
    KEY         `uid` (`uid`)
) COMMENT '直播VIP记录表';

CREATE TABLE `cmf_live_vip_log_2`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `uid`       int(11) NOT NULL,
    `name`      varchar(255)   NOT NULL DEFAULT '' COMMENT '名称',
    `price`     decimal(11, 2) NOT NULL DEFAULT '0.00' COMMENT '价格',
    `create_at` datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `update_at` datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE,
    KEY         `uid` (`uid`)
) COMMENT '直播VIP记录表';


ALTER TABLE `cmf_users_live`
    ADD COLUMN `isvip` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是VIP房间';


/** 2022.05.14 */
CREATE TABLE `cmf_gift_log`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
    `room_id`    int(11) NOT NULL DEFAULT '0' COMMENT '房间ID',
    `gift_id`    int(11) NOT NULL DEFAULT '0' COMMENT '礼物ID',
    `gift_count` int(11) NOT NULL DEFAULT '0' COMMENT '礼物数量',
    `gift_price` int(11) NOT NULL DEFAULT '0' COMMENT '礼物金额',
    `is_bag`     tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是背包',
    `create_at`  datetime DEFAULT CURRENT_TIMESTAMP,
    `update_at`  datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) COMMENT '礼物赠送记录';


ALTER TABLE `cmf_gift_log`
    ADD INDEX(`user_id`),
ADD INDEX(`room_id`),
ADD INDEX(`create_at`);

ALTER TABLE `cmf_video`
    ADD COLUMN `video_sort` tinyint(1) NOT NULL DEFAULT '1' COMMENT '视频类型；1：视频，2：短视频';

/** 2022.05.17 */
ALTER TABLE `cmf_ads`
    ADD COLUMN `ref_id` int(11) NOT NULL DEFAULT '0' COMMENT '引用资源id，资源类型见广告分类';

CREATE TABLE `cmf_video_link`
(
    `id`            int(11) unsigned NOT NULL AUTO_INCREMENT,
    `ref_id`      int(11) NOT NULL DEFAULT '0' COMMENT '链接引用ID',
    `video_id`      int(11) NOT NULL DEFAULT '0' COMMENT '视频ID',
    `uid`      int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
    `name`          varchar(50)  NOT NULL DEFAULT '' COMMENT '名称',
    `des`           varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
    `thumbnail_url` varchar(125) NOT NULL DEFAULT '' COMMENT '缩略图',
    `link_type`     tinyint(1) NOT NULL COMMENT '链接类型 1：试看短视频，2：预告短视频',
    `create_at`     datetime              DEFAULT CURRENT_TIMESTAMP,
    `update_at`     datetime              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) COMMENT '视频链接表';
/** 2022.05.18 */
ALTER TABLE `cmf_users_agent`
ADD INDEX(`one_uid`, `two_uid`);

/** 2022.05.18 */
ALTER TABLE `cmf_users_info_0` 
ADD INDEX(`votes`);
ALTER TABLE `cmf_users_info_1` 
ADD INDEX(`votes`);
ALTER TABLE `cmf_users_info_2` 
ADD INDEX(`votes`);

ALTER TABLE `cmf_users_live` 
ADD COLUMN `seq_no` int NOT NULL DEFAULT 999 COMMENT '排序',
ADD INDEX(`seq_no`);

ALTER TABLE `cmf_order` 
ADD INDEX(`create_at`);

/** 2022.05.31 */
ALTER TABLE `cmf_users_live`
ADD COLUMN `pull_flv` varchar(500) NOT NULL DEFAULT '' COMMENT '播流地址flv版本';

CREATE TABLE `cmf_video_comment_count`
(
    `id`            int(11) unsigned NOT NULL AUTO_INCREMENT,
    `comment_id`      int(11) NOT NULL DEFAULT '0' COMMENT '评论ID',
    `reply_count`   int(11) NOT NULL DEFAULT '0' COMMENT '回复数',
    `like_count`   int(11) NOT NULL DEFAULT '0' COMMENT '点赞数',
    `hot`   int(11) NOT NULL DEFAULT '0' COMMENT '热度',
    `create_at`     datetime              DEFAULT CURRENT_TIMESTAMP,
    `update_at`     datetime              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE,
    KEY         `comment_id` (`comment_id`) USING BTREE,
    ADD INDEX(`hot`),
    ADD INDEX(`create_at`)
) COMMENT '视频评论数量表';


/** 2022.06.06 */
ALTER TABLE `cmf_news`
    ADD COLUMN `lat` varchar(100) NOT NULL DEFAULT '' COMMENT '地区中心纬度',
    ADD COLUMN `lon` varchar(100) NOT NULL DEFAULT '' COMMENT '地区中心经度';

CREATE TABLE `cmf_topic`
(
    `id`            int(11) unsigned NOT NULL AUTO_INCREMENT,
    `uid`            int(11) unsigned NOT NULL DEFAULT '0' COMMENT '发起人',
    `name`      varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
    `cover`      varchar(100) NOT NULL DEFAULT '' COMMENT '封面',
    `description`   varchar(200) NOT NULL DEFAULT '' COMMENT '描述',
    `base_count`   int(11) NOT NULL DEFAULT '0' COMMENT '基数',
    `see_count`   int(11) NOT NULL DEFAULT '0' COMMENT '浏览量',
    `news_count`   int(11) NOT NULL DEFAULT '0' COMMENT '讨论量',
    `hot`   int(11) NOT NULL DEFAULT '0' COMMENT '热度',
    `disabled`   tinyint(11) NOT NULL DEFAULT '0' COMMENT '0显示 1不显示',
    `create_at`     datetime              DEFAULT CURRENT_TIMESTAMP,
    `update_at`     datetime              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) COMMENT '话题表';

CREATE TABLE `cmf_topic_news`
(
    `id`            int(11) unsigned NOT NULL AUTO_INCREMENT,
    `topic_id`            int(11) unsigned NOT NULL DEFAULT '0' COMMENT '话题id',
    `news_id`            int(11) unsigned NOT NULL DEFAULT '0' COMMENT '动态id',
    `create_at`     datetime              DEFAULT CURRENT_TIMESTAMP,
    `update_at`     datetime              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) COMMENT '话题动态表';


/** 2022.06.09 */
ALTER TABLE `cmf_users_agent`
ADD COLUMN `agt_level_id` int(11) NOT NULL DEFAULT 0 COMMENT '用户代理等级';
ALTER TABLE `cmf_users_agent`
ADD INDEX `one_uid_3`(`one_uid`, `agt_level_id`);
ALTER TABLE `cmf_agt_income`
ADD INDEX(`agent_parent_id`);

ALTER TABLE `cmf_users_live`
ADD COLUMN `can_gift` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否允许赠送礼物，1允许0禁止';
CREATE TABLE `cmf_gather` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '-1' COMMENT '用户ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `discount` tinyint(4) unsigned NOT NULL DEFAULT '100' COMMENT '合集折扣 100：不打折，0：免费',
  `discount_expire` datetime DEFAULT NULL COMMENT '折扣过期时间',
  `gather_status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '合集状态 0：删除，1：上架中，2：隐藏',
  `gather_type` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '合集类型 1：视频',
  `sort_no` int(11) NOT NULL DEFAULT '1' COMMENT '排序',
  `buy_count` int(11) NOT NULL DEFAULT '0' COMMENT '购买人数',
  `collection_count` int(11) NOT NULL DEFAULT '0' COMMENT '收藏人数',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='资源合集表';

CREATE TABLE `cmf_gather_buy` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `gather_id` int(11) NOT NULL COMMENT '合集ID',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `idx_gather_id` (`gather_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合集购买记录';

CREATE TABLE `cmf_gather_buy_0` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `gather_id` int(11) NOT NULL COMMENT '合集ID',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `idx_gather_id` (`gather_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='合集购买记录';

CREATE TABLE `cmf_gather_buy_1` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `gather_id` int(11) NOT NULL COMMENT '合集ID',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `idx_gather_id` (`gather_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='合集购买记录';

CREATE TABLE `cmf_gather_buy_2` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `gather_id` int(11) NOT NULL COMMENT '合集ID',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `idx_gather_id` (`gather_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合集购买记录';

CREATE TABLE `cmf_gather_collection` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `gather_id` int(11) NOT NULL COMMENT '合集ID',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `idx_gather_id` (`gather_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合集收藏记录';

CREATE TABLE `cmf_gather_collection_0` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `gather_id` int(11) NOT NULL COMMENT '合集ID',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `idx_gather_id` (`gather_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='合集收藏记录';

CREATE TABLE `cmf_gather_collection_1` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `gather_id` int(11) NOT NULL COMMENT '合集ID',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `idx_gather_id` (`gather_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='合集收藏记录';

CREATE TABLE `cmf_gather_collection_2` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `gather_id` int(11) NOT NULL COMMENT '合集ID',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `idx_gather_id` (`gather_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合集收藏记录';

CREATE TABLE `cmf_gather_video` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gather_id` int(11) NOT NULL COMMENT '合集ID',
  `video_id` int(11) NOT NULL COMMENT '视频ID',
  `sort_no` int(11) NOT NULL DEFAULT '1' COMMENT '排序',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_gather_id` (`gather_id`) USING BTREE,
  KEY `idx_video_id` (`video_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='视频合集关联表';

ALTER TABLE `cmf_news`
    ADD COLUMN `angle_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '角标类型：1：独家，2：出品，3：限免，4：广告';

ALTER TABLE `cmf_news`
    ADD COLUMN `address` varchar(100) NOT NULL DEFAULT '' COMMENT '详细地址',
    ADD COLUMN `touid` int(11) NOT NULL DEFAULT 0 COMMENT '被转发者id',
    ADD COLUMN `news_id` int(11) NOT NULL DEFAULT 0 COMMENT '被转发动态id';

CREATE TABLE `cmf_topic_video`
(
    `id`               int(11) unsigned NOT NULL AUTO_INCREMENT,
    `topic_id`         int(11) unsigned NOT NULL DEFAULT '0' COMMENT '话题id',
    `video_id`         int(11) unsigned NOT NULL DEFAULT '0' COMMENT '视频id',
    `create_at`        datetime              DEFAULT CURRENT_TIMESTAMP,
    `update_at`        datetime              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) COMMENT '话题视频表';

CREATE TABLE `cmf_news_user`
(
    `id`        int(11) unsigned NOT NULL AUTO_INCREMENT,
    `news_id`   int(11) unsigned NOT NULL DEFAULT '0' COMMENT '动态id',
    `uid`       int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
    `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) COMMENT '动态@用户表';

CREATE TABLE `cmf_news_comment`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `news_id`   int(11) NOT NULL COMMENT '动态id',
    `uid`       int(11) NOT NULL COMMENT '用户id',
    `touid`     int(11) NOT NULL DEFAULT 0 COMMENT '被回复的用户id',
    `reply_id`  int(11) NOT NULL DEFAULT 0 COMMENT '回复的评论id',
    `comment`   varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `disabled`  tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0显示1不显示',
    `status`    tinyint(3) NOT NULL DEFAULT 0 COMMENT '0待审核1审核通过2审核失败',
    `create_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX       `video_id`(`news_id`) USING BTREE,
    INDEX       `reply_id`(`reply_id`) USING BTREE
) COMMENT = '动态评论表';

CREATE TABLE `cmf_news_comment_count`
(
    `id`          int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `comment_id`  int(11) NOT NULL DEFAULT 0 COMMENT '评论ID',
    `reply_count` int(11) NOT NULL DEFAULT 0 COMMENT '回复数',
    `like_count`  int(11) NOT NULL DEFAULT 0 COMMENT '点赞数',
    `hot`         int(11) NOT NULL DEFAULT 0 COMMENT '热度',
    `create_at`   datetime NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at`   datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX         `comment_id`(`comment_id`) USING BTREE,
    INDEX         `hot`(`hot`) USING BTREE,
    INDEX         `create_at`(`create_at`) USING BTREE
) COMMENT = '动态评论数量表';

CREATE TABLE `cmf_news_comments_like`
(
    `id`        int(10) NOT NULL AUTO_INCREMENT,
    `uid`       int(10) NULL DEFAULT 0 COMMENT '用户id',
    `commentid` int(10) NULL DEFAULT 0 COMMENT '评论ID',
    `create_at`   datetime NULL DEFAULT CURRENT_TIMESTAMP,
    `touid`     int(12) NULL DEFAULT 0 COMMENT '被喜欢的评论者id',
    `newsid`   int(12) NULL DEFAULT 0 COMMENT '评论所属动态id',
    PRIMARY KEY (`id`) USING BTREE
) COMMENT = '动态评论点赞表';


/** 2022.06.16 */
ALTER TABLE `cmf_order`
ADD COLUMN `vip_type` tinyint(4) NOT NULL DEFAULT 0 COMMENT 'VIP类型，0无1首充2续费';

/** 2022.06.16 */
ALTER TABLE `cmf_order`
    ADD COLUMN `vip_type` tinyint(4) NOT NULL DEFAULT 0 COMMENT 'VIP类型，0无1首充2续费';

/** 2022.06.18 */
ALTER TABLE `cmf_video_like`
    ADD COLUMN `touid` int(10) NOT NULL DEFAULT 0 COMMENT '被点赞用户';
ALTER TABLE `cmf_video_like_0`
    ADD COLUMN `touid` int(10) NOT NULL DEFAULT 0 COMMENT '被点赞用户';
ALTER TABLE `cmf_video_like_1`
    ADD COLUMN `touid` int(10) NOT NULL DEFAULT 0 COMMENT '被点赞用户';
ALTER TABLE `cmf_video_like_2`
    ADD COLUMN `touid` int(10) NOT NULL DEFAULT 0 COMMENT '被点赞用户';

ALTER TABLE `cmf_news`
    ADD COLUMN `aspect` varchar(100) NOT NULL DEFAULT "" COMMENT '长宽，以-分割';

CREATE TABLE `cmf_users_message_0`
(
    `id`        int(10) NOT NULL AUTO_INCREMENT,
    `uid`       int(10) NULL DEFAULT 0 COMMENT '用户id',
    `mes_type` int(10) NULL DEFAULT 0 COMMENT '消息类型 1：点赞 2：收藏 3：@我 4：收到的评论 5：系统通知',
    `ref_type` int(10) NULL DEFAULT 0 COMMENT '引用类型 1：视频评论，2：动态评论，3：动态，4：视频，5：视频评论回复，6：动态评论回复，7：系统通知',
    `ref_id` int(10) NULL DEFAULT 0 COMMENT '引用id',
    `mes_status` int(1) NULL DEFAULT 0 COMMENT '消息状态 0：未读，1：已读',
    `create_at`   datetime NULL DEFAULT CURRENT_TIMESTAMP,
    `touid`     int(12) NULL DEFAULT 0 COMMENT '相关用户idid',
    PRIMARY KEY (`id`) USING BTREE
) COMMENT = '用户消息表';

CREATE TABLE `cmf_users_message_1`
(
    `id`        int(10) NOT NULL AUTO_INCREMENT,
    `uid`       int(10) NULL DEFAULT 0 COMMENT '用户id',
    `mes_type` int(10) NULL DEFAULT 0 COMMENT '消息类型 1：点赞 2：收藏 3：@我 4：收到的评论 5：系统通知',
    `ref_type` int(10) NULL DEFAULT 0 COMMENT '引用类型 1：视频评论，2：动态评论，3：动态，4：视频，5：视频评论回复，6：动态评论回复，7：系统通知',
    `ref_id` int(10) NULL DEFAULT 0 COMMENT '引用id',
    `mes_status` int(1) NULL DEFAULT 0 COMMENT '消息状态 0：未读，1：已读',
    `create_at`   datetime NULL DEFAULT CURRENT_TIMESTAMP,
    `touid`     int(12) NULL DEFAULT 0 COMMENT '相关用户idid',
    PRIMARY KEY (`id`) USING BTREE
) COMMENT = '用户消息表';

CREATE TABLE `cmf_users_message_2`
(
    `id`        int(10) NOT NULL AUTO_INCREMENT,
    `uid`       int(10) NULL DEFAULT 0 COMMENT '用户id',
    `mes_type` int(10) NULL DEFAULT 0 COMMENT '消息类型 1：点赞 2：收藏 3：@我 4：收到的评论 5：系统通知',
    `ref_type` int(10) NULL DEFAULT 0 COMMENT '引用类型 1：视频评论，2：动态评论，3：动态，4：视频，5：视频评论回复，6：动态评论回复，7：系统通知',
    `ref_id` int(10) NULL DEFAULT 0 COMMENT '引用id',
    `mes_status` int(1) NULL DEFAULT 0 COMMENT '消息状态 0：未读，1：已读',
    `create_at`   datetime NULL DEFAULT CURRENT_TIMESTAMP,
    `touid`     int(12) NULL DEFAULT 0 COMMENT '相关用户idid',
    PRIMARY KEY (`id`) USING BTREE
) COMMENT = '用户消息表';

/** 2022.06.21 */
ALTER TABLE `cmf_users`
    ADD COLUMN `cover_img` varchar (255) NOT NULL DEFAULT '' COMMENT '用户背景封面图';
/** 2022.06.21 */
ALTER TABLE `cmf_pushrecord`
    ADD COLUMN `title` varchar(100) NOT NULL DEFAULT "" COMMENT '标题';

CREATE TABLE `cmf_users_pushrecord`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `pushrecord_id`       int(10) NULL DEFAULT 0 COMMENT '推送id',
    `uid`       int(10) NULL DEFAULT 0 COMMENT '用户id',
    `touid` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL  COMMENT '推送对象',
    `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL  COMMENT '推送内容',
    `adminid` int(11) NOT NULL COMMENT '管理员ID',
    `admin` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '管理员账号',
    `ip` bigint(20) NOT NULL DEFAULT 0 COMMENT '管理员IP地址',
    `addtime` int(11) NOT NULL DEFAULT 0 COMMENT '发送时间',
    `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
    `pushrecord_status` int(1) NULL DEFAULT 0 COMMENT '消息状态 0：未读，1：已读',
    `pushrecord_type`         tinyint(3) NOT NULL DEFAULT 0 COMMENT '0 系统通知 1 活动通知',
    PRIMARY KEY (`id`) USING BTREE
) COMMENT = '用户推送表';

CREATE TABLE `cmf_users_pushrecord_offset`  (
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `uid`       int(10) NULL DEFAULT 0 COMMENT '用户id',
   `pull_id` int(11) NULL DEFAULT 0 COMMENT '最大已拉取推送消息id',
   `update_at`      datetime              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   PRIMARY KEY (`id`) USING BTREE
) COMMENT = '用户推送偏移表';
ALTER TABLE `cmf_pushrecord`
    ADD COLUMN `pushrecord_type`     tinyint(3) NOT NULL DEFAULT 0 COMMENT '0 系统通知 1 活动通知';
ALTER TABLE `cmf_news`
    ADD COLUMN `to_comment` varchar(255) NOT NULL DEFAULT '' COMMENT '转发评论';

ALTER TABLE `cmf_news`
    ADD COLUMN `share_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '转发数';

/** 2022.06.22 */
alter table cmf_news drop COLUMN to_comment

ALTER TABLE `cmf_users_message_0`
    ADD COLUMN `content` varchar(200) NOT NULL DEFAULT '' COMMENT  '内容',
    ADD COLUMN `cover` varchar(200) NOT NULL DEFAULT '' COMMENT  '封面';
ALTER TABLE `cmf_users_message_1`
    ADD COLUMN `content` varchar(200) NOT NULL DEFAULT '' COMMENT  '内容',
    ADD COLUMN `cover` varchar(200) NOT NULL DEFAULT '' COMMENT  '封面';
ALTER TABLE `cmf_users_message_2`
    ADD COLUMN `content` varchar(200) NOT NULL DEFAULT '' COMMENT  '内容',
    ADD COLUMN `cover` varchar(200) NOT NULL DEFAULT '' COMMENT  '封面';

ALTER TABLE `cmf_news`
    ADD COLUMN `country_name` varchar(100)  NOT NULL DEFAULT '' COMMENT '国家',
    ADD COLUMN `region_name` varchar(100)  NOT NULL DEFAULT '' COMMENT '行政区',
    ADD COLUMN `city_name` varchar(100)  NOT NULL DEFAULT '' COMMENT '城市';

/** 2022.06.23 */
alter table cmf_users_message_0 drop COLUMN content
alter table cmf_users_message_0 drop COLUMN cover
alter table cmf_users_message_1 drop COLUMN content
alter table cmf_users_message_1 drop COLUMN cover
alter table cmf_users_message_2 drop COLUMN content
alter table cmf_users_message_2 drop COLUMN cover