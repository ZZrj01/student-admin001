/*
Navicat MySQL Data Transfer

Source Server         : student
Source Server Version : 50729
Source Host           : localhost:3306
Source Database       : student

Target Server Type    : MYSQL
Target Server Version : 50729
File Encoding         : 65001

Date: 2021-12-06 09:30:32
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `tab_class`
-- ----------------------------
DROP TABLE IF EXISTS `tab_class`;
CREATE TABLE `tab_class` (
  `id` varchar(32) NOT NULL DEFAULT '',
  `name` varchar(32) DEFAULT NULL,
  `create_at` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tab_class
-- ----------------------------
INSERT INTO `tab_class` VALUES ('9d2fc267066a454bb44fc7d7ed41af45', '机器人班', '2021-11-08 04:30');
INSERT INTO `tab_class` VALUES ('9d609d06c4444ef38f1c1cba3d364701', '软件班', '2021-11-08 04:30');
INSERT INTO `tab_class` VALUES ('a10fd5dd1b5749f6a466d04954cfaeb1', '艺术班', '2021-11-08 04:30');
INSERT INTO `tab_class` VALUES ('be24f042bd884b11b8615d689052b946', '体育班', '2021-11-08 04:30');
