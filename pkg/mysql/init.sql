DROP DATABASE IF EXISTS sign_user;
CREATE DATABASE sign_user;
use sign_user;

DROP TABLE IF EXISTS user_0;
CREATE TABLE user_0(
                       id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键',
                       email VARCHAR(30) NOT NULL COMMENT '注册邮箱',
                       password VARCHAR(50) NOT NULL COMMENT '密码',
                       created_at DATETIME(3) COMMENT '创建时间',
                       updated_at DATETIME(3) COMMENT '更新时间',
                       deleted_at DATETIME(3) COMMENT '删除时间',
                       UNIQUE(email),
                       KEY(email,password),
                       KEY(deleted_at)
)ENGINE =InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS user_1;
CREATE TABLE user_1(
                       id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键',
                       email VARCHAR(30) NOT NULL COMMENT '注册邮箱',
                       password VARCHAR(50) NOT NULL COMMENT '密码',
                       created_at DATETIME(3) COMMENT '创建时间',
                       updated_at DATETIME(3) COMMENT '更新时间',
                       deleted_at DATETIME(3) COMMENT '删除时间',
                       UNIQUE(email),
                       KEY(email,password),
                       KEY(deleted_at)
)ENGINE =InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS user_2;
CREATE TABLE user_2(
                       id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键',
                       email VARCHAR(30) NOT NULL COMMENT '注册邮箱',
                       password VARCHAR(50) NOT NULL COMMENT '密码',
                       created_at DATETIME(3) COMMENT '创建时间',
                       updated_at DATETIME(3) COMMENT '更新时间',
                       deleted_at DATETIME(3) COMMENT '删除时间',
                       UNIQUE(email),
                       KEY(email,password),
                       KEY(deleted_at)
)ENGINE =InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS email_id;
CREATE TABLE email_id(
                       id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键',
                       email VARCHAR(30) NOT NULL COMMENT '注册邮箱',
                       KEY(email)
)ENGINE =InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS user_score_0;
CREATE TABLE user_score_0(
                             id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键id',
                             score BIGINT UNSIGNED DEFAULT 0 COMMENT '用户积分数',
                             freezeSub BIGINT UNSIGNED DEFAULT 0 COMMENT '预扣除积分数',
                             created_at DATETIME(3) COMMENT '创建时间',
                             updated_at DATETIME(3) COMMENT '更新时间',
                             deleted_at DATETIME(3) COMMENT '删除时间',
                             KEY(deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS user_score_1;
CREATE TABLE user_score_1(
                             id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键id',
                             score BIGINT UNSIGNED DEFAULT 0 COMMENT '用户积分数',
                             freezeSub BIGINT UNSIGNED DEFAULT 0 COMMENT '预扣除积分数',
                             created_at DATETIME(3) COMMENT '创建时间',
                             updated_at DATETIME(3) COMMENT '更新时间',
                             deleted_at DATETIME(3) COMMENT '删除时间',
                             KEY(deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS user_score_2;
CREATE TABLE user_score_2(
                             id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键id',
                             score BIGINT UNSIGNED DEFAULT 0 COMMENT '用户积分数',
                             freezeSub BIGINT UNSIGNED DEFAULT 0 COMMENT '预扣除积分数',
                             created_at DATETIME(3) COMMENT '创建时间',
                             updated_at DATETIME(3) COMMENT '更新时间',
                             deleted_at DATETIME(3) COMMENT '删除时间',
                             KEY(deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS user_group_0;
CREATE TABLE user_group_0(
                             id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键id',
                             join_groups VARCHAR(100) DEFAULT '' COMMENT '用户加入群组列表',
                             created_at DATETIME(3) COMMENT '创建时间',
                             updated_at DATETIME(3) COMMENT '更新时间',
                             deleted_at DATETIME(3) COMMENT '删除时间',
                             KEY(deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS user_group_1;
CREATE TABLE user_group_1(
                             id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键id',
                             join_groups VARCHAR(100) DEFAULT '' COMMENT '用户加入群组列表',
                             created_at DATETIME(3) COMMENT '创建时间',
                             updated_at DATETIME(3) COMMENT '更新时间',
                             deleted_at DATETIME(3) COMMENT '删除时间',
                             KEY(deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS user_group_2;
CREATE TABLE user_group_2(
                             id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键id',
                             join_groups VARCHAR(100) DEFAULT '' COMMENT '用户加入群组列表',
                             created_at DATETIME(3) COMMENT '创建时间',
                             updated_at DATETIME(3) COMMENT '更新时间',
                             deleted_at DATETIME(3) COMMENT '删除时间',
                             KEY(deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS group_info;
CREATE TABLE group_info(
                          id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '群组信息',
                          name VARCHAR(50) DEFAULT '' COMMENT '群组名',
                          owner BIGINT UNSIGNED NOT NULL  COMMENT '创建者id',
                          places VARCHAR(100) DEFAULT NULL COMMENT '签到地点坐标列表',
                          sign_in TIME DEFAULT NULL COMMENT '签到时间',
                          sign_out TIME DEFAULT NULL COMMENT '签退时间',
                          score INT UNSIGNED DEFAULT 10 COMMENT '奖励分数',
                          count INT UNSIGNED DEFAULT 0 COMMENT '群组人数',
                          created_at DATETIME(3) COMMENT '创建时间',
                          updated_at DATETIME(3) COMMENT '更新时间',
                          deleted_at DATETIME(3) COMMENT '删除时间',
                          KEY(deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;


DROP DATABASE IF EXISTS sign_info;
CREATE DATABASE sign_info;
USE sign_info;

DROP TABLE IF EXISTS sign_record_0;
CREATE TABLE sign_record_0(
                              id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键id',
                              uid BIGINT UNSIGNED NOT NULL COMMENT '用户id',
                              gid BIGINT UNSIGNED NOT NULL COMMENT '群组id',
                              sign_date VARCHAR(20)  NOT NULL COMMENT '签到日期',
                              signin_time TIME DEFAULT NULL COMMENT '签到时间',
                              signout_time TIME DEFAULT NULL COMMENT '签退时间',
                              signin_places VARCHAR(30) DEFAULT '' COMMENT '签到坐标',
                              signout_places VARCHAR(30) DEFAULT '' COMMENT '签退坐标',
                              KEY(uid),
                              KEY(gid)
)ENGINE=InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS sign_record_1;
CREATE TABLE sign_record_1(
                              id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键id',
                              uid BIGINT UNSIGNED NOT NULL COMMENT '用户id',
                              gid BIGINT UNSIGNED NOT NULL COMMENT '群组id',
                              sign_date VARCHAR(20)  NOT NULL COMMENT '签到日期',
                              signin_time TIME DEFAULT NULL COMMENT '签到时间',
                              signout_time TIME DEFAULT NULL COMMENT '签退时间',
                              signin_places VARCHAR(30) DEFAULT '' COMMENT '签到坐标',
                              signout_places VARCHAR(30) DEFAULT '' COMMENT '签退坐标',
                              KEY(uid),
                              KEY(gid)
)ENGINE=InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS sign_record_2;
CREATE TABLE sign_record_2(
                              id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键id',
                              uid BIGINT UNSIGNED NOT NULL COMMENT '用户id',
                              gid BIGINT UNSIGNED NOT NULL COMMENT '群组id',
                              sign_date VARCHAR(20)  NOT NULL COMMENT '签到日期',
                              signin_time TIME DEFAULT NULL COMMENT '签到时间',
                              signout_time TIME DEFAULT NULL COMMENT '签退时间',
                              signin_places VARCHAR(30) DEFAULT '' COMMENT '签到坐标',
                              signout_places VARCHAR(30) DEFAULT '' COMMENT '签退坐标',
                              KEY(uid),
                              KEY(gid)
)ENGINE=InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS sign_record_3;
CREATE TABLE sign_record_3(
                              id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键id',
                              uid BIGINT UNSIGNED NOT NULL COMMENT '用户id',
                              gid BIGINT UNSIGNED NOT NULL COMMENT '群组id',
                              sign_date VARCHAR(20)  NOT NULL COMMENT '签到日期',
                              signin_time TIME DEFAULT NULL COMMENT '签到时间',
                              signout_time TIME DEFAULT NULL COMMENT '签退时间',
                              signin_places VARCHAR(30) DEFAULT '' COMMENT '签到坐标',
                              signout_places VARCHAR(30) DEFAULT '' COMMENT '签退坐标',
                              KEY(uid),
                              KEY(gid)
)ENGINE=InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS sign_record_4;
CREATE TABLE sign_record_4(
                              id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键id',
                              uid BIGINT UNSIGNED NOT NULL COMMENT '用户id',
                              gid BIGINT UNSIGNED NOT NULL COMMENT '群组id',
                              sign_date VARCHAR(20)  NOT NULL COMMENT '签到日期',
                              signin_time TIME DEFAULT NULL COMMENT '签到时间',
                              signout_time TIME DEFAULT NULL COMMENT '签退时间',
                              signin_places VARCHAR(30) DEFAULT '' COMMENT '签到坐标',
                              signout_places VARCHAR(30) DEFAULT '' COMMENT '签退坐标',
                              KEY(uid),
                              KEY(gid)
)ENGINE=InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS sign_month_0;
CREATE TABLE sign_month_0(
                             id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键id',
                             uid BIGINT UNSIGNED NOT NULL COMMENT '用户id',
                             gid BIGINT UNSIGNED NOT NULL COMMENT '群组id',
                             month DATE DEFAULT NULL COMMENT '月份',
                             bitmap INT4 DEFAULT 0 COMMENT '签到bitmap',
                             created_at DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                             updated_at DATETIME(3) DEFAULT NULL COMMENT '更新时间',
                             deleted_at DATETIME(3) DEFAULT NULL COMMENT '删除时间',
                             UNIQUE (uid,gid,month),
                             KEY(gid,month),
                             KEY (deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET =utf8mb4 COLLATE =utf8mb4_bin;

DROP TABLE IF EXISTS sign_month_1;
CREATE TABLE sign_month_1(
                             id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键id',
                             uid BIGINT UNSIGNED NOT NULL COMMENT '用户id',
                             gid BIGINT UNSIGNED NOT NULL COMMENT '群组id',
                             month DATE DEFAULT NULL COMMENT '月份',
                             bitmap INT4 DEFAULT 0 COMMENT '签到bitmap',
                             created_at DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                             updated_at DATETIME(3) DEFAULT NULL COMMENT '更新时间',
                             deleted_at DATETIME(3) DEFAULT NULL COMMENT '删除时间',
                             UNIQUE (uid,gid,month),
                             KEY(gid,month),
                             KEY (deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET =utf8mb4 COLLATE =utf8mb4_bin;

DROP TABLE IF EXISTS sign_month_2;
CREATE TABLE sign_month_2(
                             id BIGINT UNSIGNED PRIMARY KEY NOT NULL COMMENT '主键id',
                             uid BIGINT UNSIGNED NOT NULL COMMENT '用户id',
                             gid BIGINT UNSIGNED NOT NULL COMMENT '群组id',
                             month DATE DEFAULT NULL COMMENT '月份',
                             bitmap INT4 DEFAULT 0 COMMENT '签到bitmap',
                             created_at DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                             updated_at DATETIME(3) DEFAULT NULL COMMENT '更新时间',
                             deleted_at DATETIME(3) DEFAULT NULL COMMENT '删除时间',
                             UNIQUE (uid,gid,month),
                             KEY(gid,month),
                             KEY (deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET =utf8mb4 COLLATE =utf8mb4_bin;


DROP DATABASE IF EXISTS sign_choose;
CREATE DATABASE sign_choose;
USE sign_choose;

DROP TABLE IF EXISTS sign_activity;
CREATE TABLE sign_activity(
                              id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
                              gid BIGINT UNSIGNED NOT NULL COMMENT '群组id',
                              start_time DATETIME(3) DEFAULT NULL COMMENT '活动开始时间',
                              end_time DATETIME(3) DEFAULT NULL COMMENT '活动结束时间',
                              prizes VARCHAR(100) DEFAULT '' COMMENT '活动奖品信息',
                              prizesTmp VARCHAR(100) DEFAULT '' COMMENT '活动奖品信息的中间状态',
                              cost BIGINT DEFAULT 100 COMMENT '活动抽奖积分数',
                              created_at DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                              updated_at DATETIME(3) DEFAULT NULL COMMENT '更新时间',
                              deleted_at DATETIME(3) DEFAULT NULL COMMENT '删除时间',
                              KEY(deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

DROP TABLE IF EXISTS sign_prizes;
CREATE TABLE sign_prizes(
                            id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
                            name VARCHAR(50) DEFAULT NULL COMMENT '奖品名称',
                            gid BIGINT UNSIGNED NOT NULL COMMENT '群组id',
                            created_at DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                            updated_at DATETIME(3) DEFAULT NULL COMMENT '更新时间',
                            deleted_at DATETIME(3) DEFAULT NULL COMMENT '删除时间',
                            KEY(deleted_at)
)ENGINE=InnoDB,DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

DROP DATABASE IF EXISTS sign_order;
CREATE DATABASE sign_order;
USE sign_order;

DROP TABLE IF EXISTS user_addr_0;
CREATE TABLE user_addr_0(
                            id BIGINT UNSIGNED PRIMARY KEY COMMENT '主键id',
                            uid BIGINT UNSIGNED NOT NULL COMMENT '用户id',
                            addr VARCHAR(50) NOT NULL COMMENT '收货地址',
                            created_at DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                            updated_at DATETIME(3) DEFAULT NULL COMMENT '更新时间',
                            deleted_at DATETIME(3) DEFAULT NULL COMMENT '删除时间',
                            KEY(uid),
                            KEY(deleted_at)
)ENGINE =InnoDb DEFAULT CHARSET=utf8mb4 COLLATE =utf8mb4_bin;

DROP TABLE IF EXISTS user_addr_1;
CREATE TABLE user_addr_1(
                            id BIGINT UNSIGNED PRIMARY KEY COMMENT '主键id',
                            uid BIGINT UNSIGNED NOT NULL COMMENT '用户id',
                            addr VARCHAR(50) NOT NULL COMMENT '收货地址',
                            created_at DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                            updated_at DATETIME(3) DEFAULT NULL COMMENT '更新时间',
                            deleted_at DATETIME(3) DEFAULT NULL COMMENT '删除时间',
                            KEY(uid),
                            KEY(deleted_at)
)ENGINE =InnoDb DEFAULT CHARSET=utf8mb4 COLLATE =utf8mb4_bin;

DROP TABLE IF EXISTS user_addr_2;
CREATE TABLE user_addr_2(
                            id BIGINT UNSIGNED PRIMARY KEY COMMENT '主键id',
                            uid BIGINT UNSIGNED NOT NULL COMMENT '用户id',
                            addr VARCHAR(50) NOT NULL COMMENT '收货地址',
                            created_at DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                            updated_at DATETIME(3) DEFAULT NULL COMMENT '更新时间',
                            deleted_at DATETIME(3) DEFAULT NULL COMMENT '删除时间',
                            KEY(uid),
                            KEY(deleted_at)
)ENGINE =InnoDb DEFAULT CHARSET=utf8mb4 COLLATE =utf8mb4_bin;

DROP TABLE IF EXISTS sign_record_0;
CREATE TABLE choose_record_0(
                                id BIGINT UNSIGNED PRIMARY KEY COMMENT '主键id',
                                uid BIGINT UNSIGNED NOT NULL COMMENT '用户id',
                                pid INT UNSIGNED NOT NULL COMMENT '奖品id',
                                getTime DATETIME(3) DEFAULT NULL COMMENT '中间时间',
                                status ENUM('未发货','已发货') DEFAULT NULL COMMENT '发货状态',
                                created_at DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                                updated_at DATETIME(3) DEFAULT NULL COMMENT '更新时间',
                                deleted_at DATETIME(3) DEFAULT NULL COMMENT '删除时间',
                                KEY (uid),
                                KEY (pid),
                                KEY (deleted_at)
)ENGINE =InnoDB,DEFAULT CHARSET =utf8mb4 COLLATE =utf8mb4_bin;

DROP TABLE IF EXISTS sign_record_1;
CREATE TABLE choose_record_1(
                                id BIGINT UNSIGNED PRIMARY KEY COMMENT '主键id',
                                uid BIGINT UNSIGNED NOT NULL COMMENT '用户id',
                                pid INT UNSIGNED NOT NULL COMMENT '奖品id',
                                getTime DATETIME(3) DEFAULT NULL COMMENT '中间时间',
                                status ENUM('未发货','已发货') DEFAULT NULL COMMENT '发货状态',
                                created_at DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                                updated_at DATETIME(3) DEFAULT NULL COMMENT '更新时间',
                                deleted_at DATETIME(3) DEFAULT NULL COMMENT '删除时间',
                                KEY (uid),
                                KEY (pid),
                                KEY (deleted_at)
)ENGINE =InnoDB,DEFAULT CHARSET =utf8mb4 COLLATE =utf8mb4_bin;

DROP TABLE IF EXISTS sign_record_2;
CREATE TABLE choose_record_2(
                                id BIGINT UNSIGNED PRIMARY KEY COMMENT '主键id',
                                uid BIGINT UNSIGNED NOT NULL COMMENT '用户id',
                                pid INT UNSIGNED NOT NULL COMMENT '奖品id',
                                getTime DATETIME(3) DEFAULT NULL COMMENT '中间时间',
                                status ENUM('未发货','已发货') DEFAULT NULL COMMENT '发货状态',
                                created_at DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                                updated_at DATETIME(3) DEFAULT NULL COMMENT '更新时间',
                                deleted_at DATETIME(3) DEFAULT NULL COMMENT '删除时间',
                                KEY (uid),
                                KEY (pid),
                                KEY (deleted_at)
)ENGINE =InnoDB,DEFAULT CHARSET =utf8mb4 COLLATE =utf8mb4_bin;

grant all privileges on *.* to 'yogen'@'%';
