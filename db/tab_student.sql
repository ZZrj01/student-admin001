/*
Navicat MySQL Data Transfer

Source Server         : student
Source Server Version : 50729
Source Host           : localhost:3306
Source Database       : student

Target Server Type    : MYSQL
Target Server Version : 50729
File Encoding         : 65001

Date: 2021-12-13 16:39:50
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `tab_student`
-- ----------------------------
DROP TABLE IF EXISTS `tab_student`;
CREATE TABLE `tab_student` (
  `id` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `sex` varchar(255) DEFAULT NULL,
  `create_at` varchar(255) DEFAULT NULL,
  `professional` varchar(255) DEFAULT NULL,
  `group` varchar(255) DEFAULT NULL,
  `seat` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tab_student
-- ----------------------------
INSERT INTO `tab_student` VALUES ('5c0332583acb4b95b25be903d545e368', '孙颖莎', '13536593840', '女', '2021-11-08', '体育', 'G1', '4');
INSERT INTO `tab_student` VALUES ('988366f250b74cdbbdb10b02e8c6ccbe', '陈梦', '13500000000', '女', '2021-11-08', '艺术', 'G1', '1');
INSERT INTO `tab_student` VALUES ('b1fe9b565fd04402a94953a0815e9a88', '许昕', '13725654524', '男', '2021-11-08', '体育', 'G1', '3');
INSERT INTO `tab_student` VALUES ('bcec6c6e9171459a800b2543f96ef6d9', '马龙', '13856457874', '男', '2021-11-08', '体育', 'G1', '2');
