/*
Navicat MySQL Data Transfer

Source Server         : root
Source Server Version : 50562
Source Host           : localhost:3306
Source Database       : student

Target Server Type    : MYSQL
Target Server Version : 50562
File Encoding         : 65001

Date: 2021-09-28 18:28:15
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `tab_user`
-- ----------------------------
DROP TABLE IF EXISTS `tab_user`;
CREATE TABLE `tab_user` (
  `id` varchar(32) NOT NULL DEFAULT '',
  `name` varchar(255) DEFAULT NULL,
  `passwd` varchar(32) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `sex` tinyint(4) DEFAULT NULL,
  `create_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tab_user
-- ----------------------------
INSERT INTO `tab_user` VALUES ('21232f297a57a5a743894a0e4a801fc3', 'admin', 'admin', '13800000000', '0', '1632824871');
