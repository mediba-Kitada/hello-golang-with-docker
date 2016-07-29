-- MySQL dump 10.13  Distrib 5.6.17, for Linux (x86_64)
--
-- Host: localhost    Database: sample
-- ------------------------------------------------------
-- Server version	5.6.17-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- ---
-- Table 'persona'
-- ペルソナ リソース用テーブル
-- ---

DROP TABLE IF EXISTS `persona`;
CREATE TABLE `persona` (
  `id` int(2) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ペルソナID',
  `sex` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT 'ペルソナ性別(0:男/女,1:男,2:女)',
  `age_min` tinyint(3) unsigned NOT NULL COMMENT 'ペルソナ年齢範囲下限',
  `age_max` tinyint(3) unsigned NOT NULL COMMENT 'ペルソナ年齢範囲上限',
  `priority` tinyint(2) UNSIGNED NOT NULL DEFAULT '99' COMMENT 'ペルソナ優先順位',
  `name` varchar(255) COMMENT 'ペルソナ名',
  `public_started` datetime NOT NULL COMMENT 'ペルソナ開始日時',
  `public_ended` datetime NOT NULL COMMENT 'ペルソナ開始日時',
  `delete_flag` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '論理削除フラグ(0:未,1:削除済み)',
  `create_userid` tinyint(3) unsigned NOT NULL COMMENT '作成ユーザーＩＤ',
  `update_userid` tinyint(3) unsigned NOT NULL COMMENT '更新ユーザーＩＤ',
  `created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'ペルソナ';
