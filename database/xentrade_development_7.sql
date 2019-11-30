/*
Navicat MySQL Data Transfer

Source Server         : xentrade
Source Server Version : 50505
Source Host           : 10.70.1.85:3306
Source Database       : xentrade_development

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2018-07-02 22:16:22
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `accounts`
-- ----------------------------
DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `member_id` int(11) DEFAULT NULL,
  `currency` int(11) DEFAULT NULL,
  `balance` decimal(32,16) DEFAULT NULL,
  `locked` decimal(32,16) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `in` decimal(32,16) DEFAULT NULL,
  `out` decimal(32,16) DEFAULT NULL,
  `default_withdraw_fund_source_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of accounts
-- ----------------------------

-- ----------------------------
-- Table structure for `account_versions`
-- ----------------------------
DROP TABLE IF EXISTS `account_versions`;
CREATE TABLE `account_versions` (
  `id` int(11) NOT NULL,
  `member_id` int(11) DEFAULT NULL,
  `account_id` int(11) DEFAULT NULL,
  `reason` int(11) DEFAULT NULL,
  `balance` decimal(32,16) DEFAULT NULL,
  `locked` decimal(32,16) DEFAULT NULL,
  `fee` decimal(32,16) DEFAULT NULL,
  `amount` decimal(32,16) DEFAULT NULL,
  `modifiable_id` int(11) DEFAULT NULL,
  `modifiable_type` varchar(255) DEFAULT NULL,
  `currency` int(11) DEFAULT NULL,
  `fun` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of account_versions
-- ----------------------------

-- ----------------------------
-- Table structure for `api_tokens`
-- ----------------------------
DROP TABLE IF EXISTS `api_tokens`;
CREATE TABLE `api_tokens` (
  `id` int(11) NOT NULL,
  `member_id` int(11) DEFAULT NULL,
  `account_id` int(11) DEFAULT NULL,
  `reason` int(11) DEFAULT NULL,
  `balance` decimal(32,16) DEFAULT NULL,
  `locked` decimal(32,16) DEFAULT NULL,
  `fee` decimal(32,16) DEFAULT NULL,
  `amount` decimal(32,16) DEFAULT NULL,
  `modifiable_id` int(11) DEFAULT NULL,
  `modifiable_type` varchar(255) DEFAULT NULL,
  `currency` int(11) DEFAULT NULL,
  `fun` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of api_tokens
-- ----------------------------

-- ----------------------------
-- Table structure for `assets`
-- ----------------------------
DROP TABLE IF EXISTS `assets`;
CREATE TABLE `assets` (
  `id` int(11) NOT NULL,
  `type` varchar(255) DEFAULT NULL,
  `attachable_id` int(11) DEFAULT NULL,
  `attachable_type` varchar(255) DEFAULT NULL,
  `file` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of assets
-- ----------------------------

-- ----------------------------
-- Table structure for `audit_logs`
-- ----------------------------
DROP TABLE IF EXISTS `audit_logs`;
CREATE TABLE `audit_logs` (
  `id` int(11) NOT NULL,
  `type` varchar(255) DEFAULT NULL,
  `operator_id` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `auditable_id` int(11) DEFAULT NULL,
  `auditable_type` varchar(255) DEFAULT NULL,
  `source_state` varchar(255) DEFAULT NULL,
  `target_state` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of audit_logs
-- ----------------------------

-- ----------------------------
-- Table structure for `authentications`
-- ----------------------------
DROP TABLE IF EXISTS `authentications`;
CREATE TABLE `authentications` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `provider` varchar(255) DEFAULT NULL,
  `uid` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `secret` varchar(255) DEFAULT NULL,
  `member_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of authentications
-- ----------------------------

-- ----------------------------
-- Table structure for `comments`
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` int(11) NOT NULL,
  `content` text,
  `author_id` int(11) DEFAULT NULL,
  `ticket_id` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of comments
-- ----------------------------

-- ----------------------------
-- Table structure for `deposits`
-- ----------------------------
DROP TABLE IF EXISTS `deposits`;
CREATE TABLE `deposits` (
  `id` int(11) NOT NULL,
  `account_id` int(11) DEFAULT NULL,
  `member_id` int(11) DEFAULT NULL,
  `currency` int(11) DEFAULT NULL,
  `amount` decimal(32,16) DEFAULT NULL,
  `fee` decimal(32,16) DEFAULT NULL,
  `fund_uid` varchar(255) DEFAULT NULL,
  `fund_extra` varchar(255) DEFAULT NULL,
  `txid` varchar(255) DEFAULT NULL,
  `state` int(11) DEFAULT NULL,
  `aasm_state` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `done_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `confirmations` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `payment_transaction_id` int(11) DEFAULT NULL,
  `txout` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of deposits
-- ----------------------------

-- ----------------------------
-- Table structure for `documents`
-- ----------------------------
DROP TABLE IF EXISTS `documents`;
CREATE TABLE `documents` (
  `id` int(11) NOT NULL,
  `key` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `body` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `is_auth` tinyint(1) DEFAULT NULL,
  `desc` text,
  `keywords` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of documents
-- ----------------------------

-- ----------------------------
-- Table structure for `document_translations`
-- ----------------------------
DROP TABLE IF EXISTS `document_translations`;
CREATE TABLE `document_translations` (
  `id` int(11) NOT NULL,
  `document_id` int(11) DEFAULT NULL,
  `locale` int(11) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `body` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `desc` text,
  `keywords` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of document_translations
-- ----------------------------

-- ----------------------------
-- Table structure for `fund_sources`
-- ----------------------------
DROP TABLE IF EXISTS `fund_sources`;
CREATE TABLE `fund_sources` (
  `id` int(11) NOT NULL,
  `member_id` int(11) DEFAULT NULL,
  `currency` int(11) DEFAULT NULL,
  `extra` varchar(255) DEFAULT NULL,
  `uid` varchar(255) DEFAULT NULL,
  `is_locked` tinyint(1) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of fund_sources
-- ----------------------------

-- ----------------------------
-- Table structure for `goose_db_version`
-- ----------------------------
DROP TABLE IF EXISTS `goose_db_version`;
CREATE TABLE `goose_db_version` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `version_id` bigint(20) NOT NULL,
  `is_applied` tinyint(1) NOT NULL,
  `tstamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of goose_db_version
-- ----------------------------

-- ----------------------------
-- Table structure for `identities`
-- ----------------------------
DROP TABLE IF EXISTS `identities`;
CREATE TABLE `identities` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(255) DEFAULT NULL,
  `password_digest` varchar(255) DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  `retry_count` int(11) DEFAULT NULL,
  `is_locked` tinyint(1) DEFAULT NULL,
  `locked_at` datetime DEFAULT NULL,
  `token` varchar(255) NOT NULL,
  `last_verify_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of identities
-- ----------------------------

-- ----------------------------
-- Table structure for `id_documents`
-- ----------------------------
DROP TABLE IF EXISTS `id_documents`;
CREATE TABLE `id_documents` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_document_type` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `id_document_number` varchar(255) DEFAULT NULL,
  `member_id` int(11) DEFAULT NULL,
  `address` text,
  `city` varchar(255) DEFAULT NULL,
  `country` varchar(255) DEFAULT NULL,
  `zipcode` varchar(255) DEFAULT NULL,
  `id_bill_type` int(11) DEFAULT NULL,
  `aasm_state` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `birth_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of id_documents
-- ----------------------------

-- ----------------------------
-- Table structure for `members`
-- ----------------------------
DROP TABLE IF EXISTS `members`;
CREATE TABLE `members` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sn` varchar(255) DEFAULT NULL,
  `display_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `identity_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `state` int(11) DEFAULT NULL,
  `activated` tinyint(1) DEFAULT NULL,
  `country_code` int(11) DEFAULT NULL,
  `phone_number` varchar(255) DEFAULT NULL,
  `disabled` tinyint(1) DEFAULT NULL,
  `api_disabled` tinyint(1) DEFAULT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `account_type` int(11) DEFAULT NULL,
  `first_name` varchar(255) DEFAULT NULL,
  `middle_name` varchar(255) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `birth_date` datetime DEFAULT NULL,
  `residential_address` varchar(255) DEFAULT NULL,
  `suburb` varchar(255) DEFAULT NULL,
  `postcode` varchar(255) DEFAULT NULL,
  `country` int(11) DEFAULT NULL,
  `google_activated` tinyint(1) DEFAULT NULL,
  `name_on_card` varchar(255) DEFAULT NULL,
  `card_number` varchar(255) DEFAULT NULL,
  `card_expiry_date` datetime DEFAULT NULL,
  `card_ccv` varchar(255) DEFAULT NULL,
  `bank_account_name` varchar(255) DEFAULT NULL,
  `bank_bsb` varchar(255) DEFAULT NULL,
  `bank_account_number` varchar(255) DEFAULT NULL,
  `bank_pay_id` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of members
-- ----------------------------

-- ----------------------------
-- Table structure for `oauth_access_grants`
-- ----------------------------
DROP TABLE IF EXISTS `oauth_access_grants`;
CREATE TABLE `oauth_access_grants` (
  `id` int(11) NOT NULL,
  `resource_owner_id` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `application_id` int(11) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `expires_in` int(11) DEFAULT NULL,
  `redirect_uri` text,
  `scopes` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of oauth_access_grants
-- ----------------------------

-- ----------------------------
-- Table structure for `oauth_access_tokens`
-- ----------------------------
DROP TABLE IF EXISTS `oauth_access_tokens`;
CREATE TABLE `oauth_access_tokens` (
  `id` int(11) NOT NULL,
  `resource_owner_id` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `application_id` int(11) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `refresh_token` varchar(255) DEFAULT NULL,
  `revoked_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `expires_in` int(11) DEFAULT NULL,
  `scopes` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of oauth_access_tokens
-- ----------------------------

-- ----------------------------
-- Table structure for `oauth_applications`
-- ----------------------------
DROP TABLE IF EXISTS `oauth_applications`;
CREATE TABLE `oauth_applications` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `uid` varchar(255) DEFAULT NULL,
  `secret` varchar(255) DEFAULT NULL,
  `redirect_uri` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of oauth_applications
-- ----------------------------

-- ----------------------------
-- Table structure for `orders`
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `bid` int(11) DEFAULT NULL,
  `ask` int(11) DEFAULT NULL,
  `currency` int(11) DEFAULT NULL,
  `price` decimal(32,16) DEFAULT NULL,
  `volume` decimal(32,16) DEFAULT NULL,
  `origin_volume` decimal(32,16) DEFAULT NULL,
  `state` int(11) DEFAULT NULL,
  `done_at` datetime DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `member_id` int(11) DEFAULT NULL,
  `sn` varchar(255) DEFAULT NULL,
  `source` varchar(255) NOT NULL,
  `ord_type` varchar(255) DEFAULT NULL,
  `locked` decimal(32,16) DEFAULT NULL,
  `origin_locked` decimal(32,16) DEFAULT NULL,
  `funds_received` decimal(32,16) DEFAULT NULL,
  `trades_count` decimal(32,16) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of orders
-- ----------------------------

-- ----------------------------
-- Table structure for `partial_trees`
-- ----------------------------
DROP TABLE IF EXISTS `partial_trees`;
CREATE TABLE `partial_trees` (
  `id` int(11) NOT NULL,
  `proof_id` int(11) DEFAULT NULL,
  `account_id` int(11) DEFAULT NULL,
  `json` text,
  `sum` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of partial_trees
-- ----------------------------

-- ----------------------------
-- Table structure for `payment_address`
-- ----------------------------
DROP TABLE IF EXISTS `payment_address`;
CREATE TABLE `payment_address` (
  `id` int(11) NOT NULL,
  `account_id` int(11) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `currency` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of payment_address
-- ----------------------------

-- ----------------------------
-- Table structure for `payment_transactions`
-- ----------------------------
DROP TABLE IF EXISTS `payment_transactions`;
CREATE TABLE `payment_transactions` (
  `id` int(11) NOT NULL,
  `txid` varchar(255) DEFAULT NULL,
  `amount` decimal(32,16) DEFAULT NULL,
  `confirmations` int(11) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `state` int(11) DEFAULT NULL,
  `aasm_state` varchar(255) DEFAULT NULL,
  `receive_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `done_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `currency` int(11) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `txout` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of payment_transactions
-- ----------------------------

-- ----------------------------
-- Table structure for `proofs`
-- ----------------------------
DROP TABLE IF EXISTS `proofs`;
CREATE TABLE `proofs` (
  `id` int(11) NOT NULL,
  `root` varchar(255) DEFAULT NULL,
  `currency` int(11) DEFAULT NULL,
  `ready` tinyint(1) DEFAULT NULL,
  `sum` varchar(255) DEFAULT NULL,
  `address` text,
  `balance` varchar(30) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of proofs
-- ----------------------------

-- ----------------------------
-- Table structure for `read_marks`
-- ----------------------------
DROP TABLE IF EXISTS `read_marks`;
CREATE TABLE `read_marks` (
  `id` int(11) NOT NULL,
  `read_id` int(11) DEFAULT NULL,
  `readable_id` int(11) DEFAULT NULL,
  `reader_type` int(11) DEFAULT NULL,
  `member_id` int(11) DEFAULT NULL,
  `readable_type` varchar(20) DEFAULT NULL,
  `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of read_marks
-- ----------------------------

-- ----------------------------
-- Table structure for `running_accounts`
-- ----------------------------
DROP TABLE IF EXISTS `running_accounts`;
CREATE TABLE `running_accounts` (
  `id` int(11) NOT NULL,
  `category` int(11) DEFAULT NULL,
  `income` decimal(32,16) DEFAULT NULL,
  `expenses` decimal(32,16) DEFAULT NULL,
  `currency` int(11) DEFAULT NULL,
  `member_id` int(11) DEFAULT NULL,
  `source_id` int(11) DEFAULT NULL,
  `source_type` varchar(255) DEFAULT NULL,
  `note` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of running_accounts
-- ----------------------------

-- ----------------------------
-- Table structure for `schema_migrations`
-- ----------------------------
DROP TABLE IF EXISTS `schema_migrations`;
CREATE TABLE `schema_migrations` (
  `version` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of schema_migrations
-- ----------------------------

-- ----------------------------
-- Table structure for `signup_histories`
-- ----------------------------
DROP TABLE IF EXISTS `signup_histories`;
CREATE TABLE `signup_histories` (
  `id` int(11) NOT NULL,
  `member_id` int(11) DEFAULT NULL,
  `ip` varchar(255) DEFAULT NULL,
  `accept_language` varchar(255) DEFAULT NULL,
  `ua` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of signup_histories
-- ----------------------------

-- ----------------------------
-- Table structure for `simple_captcha_data`
-- ----------------------------
DROP TABLE IF EXISTS `simple_captcha_data`;
CREATE TABLE `simple_captcha_data` (
  `id` int(11) NOT NULL,
  `key` varchar(255) DEFAULT NULL,
  `value` varchar(6) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of simple_captcha_data
-- ----------------------------

-- ----------------------------
-- Table structure for `taggings`
-- ----------------------------
DROP TABLE IF EXISTS `taggings`;
CREATE TABLE `taggings` (
  `id` int(11) NOT NULL,
  `tag_id` int(11) DEFAULT NULL,
  `taggable_id` int(11) DEFAULT NULL,
  `taggable_type` varchar(255) DEFAULT NULL,
  `tagger_id` int(11) DEFAULT NULL,
  `tagger_type` varchar(255) DEFAULT NULL,
  `context` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of taggings
-- ----------------------------

-- ----------------------------
-- Table structure for `tags`
-- ----------------------------
DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tags
-- ----------------------------

-- ----------------------------
-- Table structure for `tickets`
-- ----------------------------
DROP TABLE IF EXISTS `tickets`;
CREATE TABLE `tickets` (
  `id` int(11) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `content` text,
  `aasm_state` varchar(255) DEFAULT NULL,
  `author_id` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tickets
-- ----------------------------

-- ----------------------------
-- Table structure for `tokens`
-- ----------------------------
DROP TABLE IF EXISTS `tokens`;
CREATE TABLE `tokens` (
  `id` int(11) NOT NULL,
  `token` varchar(255) DEFAULT NULL,
  `expire_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `member_id` int(11) DEFAULT NULL,
  `is_used` tinyint(1) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tokens
-- ----------------------------

-- ----------------------------
-- Table structure for `trades`
-- ----------------------------
DROP TABLE IF EXISTS `trades`;
CREATE TABLE `trades` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `price` decimal(32,16) DEFAULT NULL,
  `volume` decimal(32,16) DEFAULT NULL,
  `ask_id` int(11) DEFAULT NULL,
  `bid_id` int(11) DEFAULT NULL,
  `trend` int(11) DEFAULT NULL,
  `currency` int(11) DEFAULT NULL,
  `ask_member_id` int(11) DEFAULT NULL,
  `bid_member_id` int(11) DEFAULT NULL,
  `funds` decimal(32,16) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of trades
-- ----------------------------

-- ----------------------------
-- Table structure for `two_factors`
-- ----------------------------
DROP TABLE IF EXISTS `two_factors`;
CREATE TABLE `two_factors` (
  `id` int(11) NOT NULL,
  `member_id` int(11) DEFAULT NULL,
  `otp_secret` varchar(255) DEFAULT NULL,
  `last_verify_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `activated` tinyint(1) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of two_factors
-- ----------------------------

-- ----------------------------
-- Table structure for `users`
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` varchar(36) NOT NULL,
  `sub` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `givenName` varchar(255) DEFAULT NULL,
  `familyName` varchar(255) DEFAULT NULL,
  `profile` varchar(255) DEFAULT NULL,
  `picture` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `emailVerified` varchar(255) DEFAULT NULL,
  `gender` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `reset_token` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of users
-- ----------------------------

-- ----------------------------
-- Table structure for `versions`
-- ----------------------------
DROP TABLE IF EXISTS `versions`;
CREATE TABLE `versions` (
  `id` int(11) NOT NULL,
  `item_type` varchar(255) DEFAULT NULL,
  `item_id` int(11) DEFAULT NULL,
  `event` varchar(255) DEFAULT NULL,
  `whodunnit` varchar(255) DEFAULT NULL,
  `object` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of versions
-- ----------------------------

-- ----------------------------
-- Table structure for `withdraws`
-- ----------------------------
DROP TABLE IF EXISTS `withdraws`;
CREATE TABLE `withdraws` (
  `id` int(11) NOT NULL,
  `sn` varchar(255) DEFAULT NULL,
  `account_id` int(11) DEFAULT NULL,
  `member_id` int(11) DEFAULT NULL,
  `currency` int(11) DEFAULT NULL,
  `amount` decimal(32,16) DEFAULT NULL,
  `fee` decimal(32,16) DEFAULT NULL,
  `fund_uid` varchar(255) DEFAULT NULL,
  `done_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `txid` varchar(255) DEFAULT NULL,
  `aasm_state` varchar(255) DEFAULT NULL,
  `sum` decimal(32,16) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of withdraws
-- ----------------------------
