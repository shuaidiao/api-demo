/*
 Navicat Premium Data Transfer

 Source Server         : 本机测试
 Source Server Type    : MySQL
 Source Server Version : 50722
 Source Host           : localhost:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 50722
 File Encoding         : 65001

 Date: 31/07/2019 16:53:03
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'aaa', '2019-07-31 14:28:11', '2019-07-31 14:28:13');
INSERT INTO `user` VALUES (2, 'bbb', '2019-07-31 14:28:17', '2019-07-31 14:28:19');
INSERT INTO `user` VALUES (3, 'ccc', '2019-07-31 14:28:23', '2019-07-31 14:28:25');
INSERT INTO `user` VALUES (10, 'ddd', '2019-07-31 16:42:31', '2019-07-31 16:42:31');

SET FOREIGN_KEY_CHECKS = 1;
