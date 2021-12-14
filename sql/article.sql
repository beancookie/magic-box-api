/*
 Navicat Premium Data Transfer

 Source Server         : 10.144.201.235
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : 10.144.201.235:4408
 Source Schema         : magic_box

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 03/12/2021 18:20:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_on` int(11) NULL DEFAULT NULL,
  `modified_on` int(11) NULL DEFAULT NULL,
  `deleted_on` tinyint(4) NOT NULL DEFAULT 0,
  `title` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `brief_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `collect_count` int(11) NULL DEFAULT NULL,
  `cover_image` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `comment_count` int(11) NULL DEFAULT NULL,
  `user_id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `user_name` varchar(200) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `tag_id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `platform` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
