/*
 Navicat Premium Data Transfer

 Source Server         : w
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : localhost:3306
 Source Schema         : pest

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 07/02/2020 19:58:40
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for car_visit_list
-- ----------------------------
DROP TABLE IF EXISTS `car_visit_list`;
CREATE TABLE `car_visit_list`  (
  `uid` bigint(0) NOT NULL AUTO_INCREMENT,
  `car_no` varchar(0) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `car_picture` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `driver_license_picture` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `diiver_permit_picture` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `goods` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `visit_id` bigint(0) NULL DEFAULT NULL,
  `create_time` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for daily_report_list
-- ----------------------------
DROP TABLE IF EXISTS `daily_report_list`;
CREATE TABLE `daily_report_list`  (
  `uid` bigint(0) NOT NULL AUTO_INCREMENT,
  `personnel_id` bigint(0) NULL DEFAULT NULL,
  `report_time` varchar(0) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `symptom` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `hospitalized_flag` tinyint(1) NOT NULL COMMENT '0否1是',
  `temperature` float NULL DEFAULT NULL,
  `touch_people` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for guard_list
-- ----------------------------
DROP TABLE IF EXISTS `guard_list`;
CREATE TABLE `guard_list`  (
  `uid` bigint(0) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `master` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for guarder_list
-- ----------------------------
DROP TABLE IF EXISTS `guarder_list`;
CREATE TABLE `guarder_list`  (
  `uid` bigint(0) NOT NULL AUTO_INCREMENT,
  `username` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名，登陆可用',
  `password` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '经salt加密过的密码',
  `nickname` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `salt` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `state` tinyint(1) NOT NULL COMMENT '用户状态',
  `power` tinyint(0) NOT NULL COMMENT '0管理人员1工作人员',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户头像地址',
  `user_agent` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户登录的user_agent',
  `create_time` bigint(0) NULL DEFAULT NULL COMMENT '注册时间',
  `card_no` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `identity` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = 'W_后台登录用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for house_list
-- ----------------------------
DROP TABLE IF EXISTS `house_list`;
CREATE TABLE `house_list`  (
  `uid` bigint(0) NOT NULL AUTO_INCREMENT,
  `nature` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `street` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `street_no` int(0) NULL DEFAULT NULL,
  `address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `number` int(0) NULL DEFAULT NULL,
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for isolation_list
-- ----------------------------
DROP TABLE IF EXISTS `isolation_list`;
CREATE TABLE `isolation_list`  (
  `uid` bigint(0) NOT NULL AUTO_INCREMENT,
  `house_id` bigint(0) NOT NULL,
  `flag` tinyint(1) NULL DEFAULT NULL,
  `reason` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `start_time` varchar(0) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `end_time` varchar(0) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `real_complete_time` varchar(0) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `unit` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for personnel_list
-- ----------------------------
DROP TABLE IF EXISTS `personnel_list`;
CREATE TABLE `personnel_list`  (
  `uid` bigint(0) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `occupation` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `card_no` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `card_picture` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `face_picture` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sex` tinyint(1) NULL DEFAULT 0 COMMENT '0男1女',
  `nation` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `birthday` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sign_organization` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `limited_date` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `history` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for relation_house_personnel
-- ----------------------------
DROP TABLE IF EXISTS `relation_house_personnel`;
CREATE TABLE `relation_house_personnel`  (
  `uid` bigint(0) NOT NULL AUTO_INCREMENT,
  `house_id` bigint(0) NULL DEFAULT NULL,
  `personnel_id` bigint(0) NULL DEFAULT NULL,
  `holder_flag` tinyint(1) NULL DEFAULT NULL,
  `relation_holder` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for relation_unit_personnel
-- ----------------------------
DROP TABLE IF EXISTS `relation_unit_personnel`;
CREATE TABLE `relation_unit_personnel`  (
  `uid` bigint(0) NOT NULL AUTO_INCREMENT,
  `unit_id` bigint(0) NULL DEFAULT NULL,
  `personnel_id` bigint(0) NULL DEFAULT NULL,
  `position` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for touch_list
-- ----------------------------
DROP TABLE IF EXISTS `touch_list`;
CREATE TABLE `touch_list`  (
  `uid` bigint(0) NOT NULL AUTO_INCREMENT,
  `personnel_id` bigint(0) NULL DEFAULT NULL,
  `way` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `place` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `touch_number` bigint(0) NULL DEFAULT NULL,
  `touch_people` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for unit_list
-- ----------------------------
DROP TABLE IF EXISTS `unit_list`;
CREATE TABLE `unit_list`  (
  `uid` bigint(0) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `house_id` bigint(0) NULL DEFAULT NULL,
  `license_number` bigint(0) NULL DEFAULT NULL,
  `identification_number` bigint(0) NULL DEFAULT NULL,
  `picture` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `kind` int(10) UNSIGNED ZEROFILL NULL DEFAULT NULL COMMENT '类别',
  `scale` int(0) NULL DEFAULT NULL COMMENT '规模人数',
  `tel` int(0) NULL DEFAULT NULL,
  `bank_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `bank_account` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '备注',
  `create_time` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_list
-- ----------------------------
DROP TABLE IF EXISTS `user_list`;
CREATE TABLE `user_list`  (
  `uid` bigint(0) NOT NULL AUTO_INCREMENT,
  `username` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名，登陆可用',
  `password` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '经salt加密过的密码',
  `nickname` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `salt` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `state` tinyint(1) NOT NULL COMMENT '用户状态',
  `power` tinyint(0) NOT NULL COMMENT '0管理人员1工作人员',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户头像地址',
  `user_agent` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户登录的user_agent',
  `create_time` bigint(0) NULL DEFAULT NULL COMMENT '注册时间',
  `card_no` int(0) NOT NULL,
  `identity` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `guard_id` int(0) NOT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = 'W_后台登录用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for visit_list
-- ----------------------------
DROP TABLE IF EXISTS `visit_list`;
CREATE TABLE `visit_list`  (
  `uid` bigint(0) NOT NULL AUTO_INCREMENT,
  `guarder_id` bigint(0) NULL DEFAULT NULL,
  `time` bigint(0) NULL DEFAULT NULL,
  `goal` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `purpose` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `relation` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for visitor_list
-- ----------------------------
DROP TABLE IF EXISTS `visitor_list`;
CREATE TABLE `visitor_list`  (
  `uid` bigint(0) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `card_no` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `temperature` float NULL DEFAULT NULL,
  `picture` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `contact` int(0) UNSIGNED NULL DEFAULT NULL,
  `result` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `visit_id` bigint(0) NULL DEFAULT NULL,
  `create_time` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- View structure for house_personnel
-- ----------------------------
DROP VIEW IF EXISTS `house_personnel`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `house_personnel` AS select `relation_house_personnel`.`uid` AS `uid`,`relation_house_personnel`.`house_id` AS `house_id`,`relation_house_personnel`.`personnel_id` AS `personnel_id`,`relation_house_personnel`.`holder_flag` AS `holder_flag`,`relation_house_personnel`.`relation_holder` AS `relation_holder`,`house_list`.`nature` AS `nature`,`house_list`.`street` AS `street`,`house_list`.`street_no` AS `street_no`,`house_list`.`address` AS `address`,`house_list`.`number` AS `number`,`house_list`.`comment` AS `comment`,`personnel_list`.`name` AS `name`,`personnel_list`.`occupation` AS `occupation`,`personnel_list`.`card_no` AS `card_no`,`personnel_list`.`card_picture` AS `card_picture`,`personnel_list`.`face_picture` AS `face_picture`,`personnel_list`.`history` AS `history` from ((`relation_house_personnel` join `house_list` on((`relation_house_personnel`.`house_id` = `house_list`.`uid`))) join `personnel_list` on((`relation_house_personnel`.`personnel_id` = `personnel_list`.`uid`)));

-- ----------------------------
-- View structure for house_unit
-- ----------------------------
DROP VIEW IF EXISTS `house_unit`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `house_unit` AS select `unit_list`.`uid` AS `uid`,`unit_list`.`house_id` AS `house_id`,`unit_list`.`name` AS `name`,`unit_list`.`license_number` AS `license_number`,`unit_list`.`identification_number` AS `identification_number`,`unit_list`.`picture` AS `picture`,`unit_list`.`kind` AS `kind`,`unit_list`.`scale` AS `scale`,`unit_list`.`tel` AS `tel`,`unit_list`.`bank_name` AS `bank_name`,`unit_list`.`bank_account` AS `bank_account`,`unit_list`.`comment` AS `unit_comment`,`house_list`.`nature` AS `nature`,`house_list`.`street` AS `street`,`house_list`.`street_no` AS `street_no`,`house_list`.`address` AS `address`,`house_list`.`number` AS `number`,`house_list`.`comment` AS `house_comment` from (`unit_list` left join `house_list` on((`unit_list`.`house_id` = `house_list`.`uid`)));

-- ----------------------------
-- View structure for unit_personnel
-- ----------------------------
DROP VIEW IF EXISTS `unit_personnel`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `unit_personnel` AS select `relation_unit_personnel`.`uid` AS `uid`,`relation_unit_personnel`.`personnel_id` AS `personnel_id`,`relation_unit_personnel`.`unit_id` AS `unit_id`,`relation_unit_personnel`.`position` AS `position`,`personnel_list`.`name` AS `personnel_name`,`personnel_list`.`occupation` AS `occupation`,`personnel_list`.`card_no` AS `card_no`,`personnel_list`.`card_picture` AS `card_picture`,`personnel_list`.`face_picture` AS `face_picture`,`personnel_list`.`history` AS `history`,`house_unit`.`house_id` AS `house_id`,`house_unit`.`name` AS `unit_name`,`house_unit`.`license_number` AS `license_number`,`house_unit`.`identification_number` AS `identification_number`,`house_unit`.`picture` AS `unit_picture`,`house_unit`.`kind` AS `kind`,`house_unit`.`scale` AS `scale`,`house_unit`.`tel` AS `tel`,`house_unit`.`bank_name` AS `bank_name`,`house_unit`.`bank_account` AS `bank_account`,`house_unit`.`nature` AS `nature`,`house_unit`.`street` AS `street`,`house_unit`.`street_no` AS `street_no`,`house_unit`.`address` AS `address`,`house_unit`.`unit_comment` AS `unit_comment`,`house_unit`.`house_comment` AS `house_comment`,`house_unit`.`number` AS `number` from ((`relation_unit_personnel` left join `personnel_list` on((`relation_unit_personnel`.`personnel_id` = `personnel_list`.`uid`))) left join `house_unit` on((`relation_unit_personnel`.`unit_id` = `house_unit`.`uid`)));

SET FOREIGN_KEY_CHECKS = 1;
