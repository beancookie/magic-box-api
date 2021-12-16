/*
 Navicat Premium Data Transfer

 Source Server         : 118.24.151.15
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 118.24.151.15:3306
 Source Schema         : magic_box

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 16/12/2021 14:48:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for hot_new
-- ----------------------------
DROP TABLE IF EXISTS `hot_new`;
CREATE TABLE `hot_new`  (
  `id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `onboard_time` int(11) NULL DEFAULT NULL,
  `category` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `raw_hot` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `key_word` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `platform` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `created_on` int(255) NULL DEFAULT NULL,
  `modified_on` int(255) NULL DEFAULT NULL,
  `deleted_on` int(255) NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
