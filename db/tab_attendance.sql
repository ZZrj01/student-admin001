/*
Navicat MySQL Data Transfer

Source Server         : student
Source Server Version : 50729
Source Host           : localhost:3306
Source Database       : student

Target Server Type    : MYSQL
Target Server Version : 50729
File Encoding         : 65001

Date: 2021-12-13 15:55:32
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `tab_attendance`
-- ----------------------------
DROP TABLE IF EXISTS `tab_attendance`;
CREATE TABLE `tab_attendance` (
  `create_at` varchar(255) DEFAULT NULL COMMENT '开始时间',
  `finish_at` varchar(255) DEFAULT NULL COMMENT '结束时间',
  `session` varchar(255) DEFAULT NULL COMMENT '节次',
  `classnum` varchar(255) DEFAULT NULL COMMENT '班级人数',
  `id` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tab_attendance
-- ----------------------------
INSERT INTO `tab_attendance` VALUES ('2021-12-13', '', '7-8', '', '709e4623371b4d58aac06004fb478124');
INSERT INTO `tab_attendance` VALUES ('2021-12-13', '', '9-10', '', 'e0915db840ac4b0f885dc42153583ce7');
INSERT INTO `tab_attendance` VALUES ('2021-12-13', '', '3-4', '', 'e5ae2cefc8fc4cb99f5edc38df92a115');
