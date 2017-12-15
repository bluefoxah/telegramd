-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: 2017-12-15 14:19:11
-- 服务器版本： 5.6.37
-- PHP Version: 5.6.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `nebulaim`
--

-- --------------------------------------------------------

--
-- 表的结构 `apps`
--

CREATE TABLE `apps` (
  `id` int(11) NOT NULL,
  `api_id` int(11) NOT NULL,
  `api_hash` varchar(256) NOT NULL,
  `title` varchar(128) NOT NULL,
  `short_name` varchar(128) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='apps';

-- --------------------------------------------------------

--
-- 表的结构 `app_configs`
--

CREATE TABLE `app_configs` (
  `app_id` int(11) NOT NULL,
  `config_key` int(11) NOT NULL,
  `config_value` int(11) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` int(11) NOT NULL,
  `updated_at` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `app_ios_push_certs`
--

CREATE TABLE `app_ios_push_certs` (
  `cert_id` int(11) NOT NULL,
  `app_id` int(11) NOT NULL,
  `bundle_id` int(11) NOT NULL,
  `cert_type` int(11) NOT NULL,
  `cert_memo` int(11) NOT NULL,
  `uploaded` int(11) NOT NULL,
  `expired` int(11) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` int(11) NOT NULL,
  `updated_at` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `app_keys`
--

CREATE TABLE `app_keys` (
  `app_id` int(11) NOT NULL,
  `app_key` varchar(256) NOT NULL,
  `app_secret` varchar(256) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` int(11) NOT NULL,
  `refresher` int(11) NOT NULL,
  `refreshed_at` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `auths`
--

CREATE TABLE `auths` (
  `id` int(11) NOT NULL,
  `auth_id` bigint(20) NOT NULL,
  `api_id` int(11) NOT NULL,
  `device_model` varchar(255) NOT NULL DEFAULT '',
  `system_version` varchar(255) NOT NULL DEFAULT '',
  `app_version` varchar(255) NOT NULL DEFAULT '',
  `system_lang_code` varchar(255) NOT NULL DEFAULT '',
  `lang_pack` varchar(255) NOT NULL DEFAULT '',
  `lang_code` varchar(255) NOT NULL DEFAULT '',
  `connection_hash` bigint(20) NOT NULL DEFAULT '0' COMMENT 'initConnection消息hash值',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `auth_keys`
--

CREATE TABLE `auth_keys` (
  `id` int(11) NOT NULL,
  `auth_id` bigint(20) NOT NULL COMMENT 'auth_id',
  `body` varchar(512) NOT NULL COMMENT 'auth_key，原始数据为256的二进制数据，存储时转换成base64格式',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `auth_phone_transactions`
--

CREATE TABLE `auth_phone_transactions` (
  `id` int(11) NOT NULL,
  `transaction_hash` varchar(255) NOT NULL,
  `api_id` int(11) NOT NULL,
  `api_hash` varchar(255) NOT NULL,
  `phone_number` varchar(32) NOT NULL,
  `code` varchar(8) NOT NULL,
  `attempts` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `auth_salts`
--

CREATE TABLE `auth_salts` (
  `id` int(11) NOT NULL,
  `auth_id` bigint(20) NOT NULL,
  `salt` bigint(20) NOT NULL,
  `valid_since` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `auth_users`
--

CREATE TABLE `auth_users` (
  `id` int(11) NOT NULL,
  `auth_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `hash` bigint(20) NOT NULL DEFAULT '0',
  `device_model` varchar(128) NOT NULL DEFAULT '',
  `platform` varchar(64) NOT NULL DEFAULT '',
  `system_version` varchar(64) NOT NULL DEFAULT '',
  `api_id` int(11) NOT NULL DEFAULT '0',
  `app_name` varchar(64) NOT NULL DEFAULT '',
  `app_version` varchar(64) NOT NULL DEFAULT '',
  `date_created` int(11) NOT NULL DEFAULT '0',
  `date_active` int(11) NOT NULL DEFAULT '0',
  `ip` varchar(64) NOT NULL DEFAULT '',
  `country` varchar(64) NOT NULL DEFAULT '',
  `region` varchar(64) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `channels`
--

CREATE TABLE `channels` (
  `id` int(11) NOT NULL,
  `access_hash` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `channel_users`
--

CREATE TABLE `channel_users` (
  `id` int(11) NOT NULL,
  `channel_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `chats`
--

CREATE TABLE `chats` (
  `id` int(11) NOT NULL,
  `creator_user_id` int(11) NOT NULL,
  `create_random_id` bigint(20) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `participant_count` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `title_changer_user_id` int(11) NOT NULL,
  `title_changed_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `title_change_random_id` bigint(20) NOT NULL,
  `avatar_changer_user_id` int(11) NOT NULL,
  `avatar_changed_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `avatar_change_random_id` bigint(20) NOT NULL,
  `is_public` tinyint(1) NOT NULL DEFAULT '0',
  `about` text NOT NULL,
  `topic` varchar(255) DEFAULT '',
  `is_hidden` tinyint(1) DEFAULT '0',
  `version` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `chat_participants`
--

CREATE TABLE `chat_participants` (
  `id` int(11) NOT NULL,
  `chat_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `participant_type` tinyint(4) DEFAULT '0',
  `inviter_user_id` int(11) NOT NULL DEFAULT '0',
  `invited_at` int(11) NOT NULL DEFAULT '0',
  `joined_at` int(11) NOT NULL DEFAULT '0',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `chat_users`
--

CREATE TABLE `chat_users` (
  `id` int(11) NOT NULL,
  `chat_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `participant_type` tinyint(4) DEFAULT '0',
  `inviter_user_id` int(11) NOT NULL DEFAULT '0',
  `invited_at` int(11) NOT NULL DEFAULT '0',
  `joined_at` int(11) NOT NULL DEFAULT '0',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `client_updates_state`
--

CREATE TABLE `client_updates_state` (
  `id` int(11) NOT NULL,
  `auth_key_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `pts` int(11) NOT NULL DEFAULT '0',
  `qts` int(11) NOT NULL DEFAULT '0',
  `seq` int(11) NOT NULL DEFAULT '0',
  `date2` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `devices`
--

CREATE TABLE `devices` (
  `id` int(11) NOT NULL,
  `auth_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `token_type` tinyint(4) NOT NULL,
  `token` varchar(255) NOT NULL,
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `files`
--

CREATE TABLE `files` (
  `id` bigint(20) NOT NULL,
  `creator_user_id` int(11) NOT NULL,
  `file_id` bigint(20) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `file_parts` int(11) NOT NULL,
  `file_size` bigint(20) NOT NULL,
  `md5_checksum` char(33) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `file_parts`
--

CREATE TABLE `file_parts` (
  `id` bigint(20) NOT NULL,
  `creator_user_id` int(11) NOT NULL,
  `file_id` bigint(20) NOT NULL,
  `file_part` int(11) NOT NULL,
  `is_big_file` tinyint(4) NOT NULL,
  `file_total_parts` int(11) NOT NULL,
  `bytes` blob NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `messages`
--

CREATE TABLE `messages` (
  `id` int(11) NOT NULL,
  `sender_user_id` int(11) NOT NULL,
  `peer_type` int(11) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `random_id` bigint(20) NOT NULL,
  `message_type` tinyint(4) NOT NULL DEFAULT '0',
  `message_data` text NOT NULL,
  `date2` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `message_boxes`
--

CREATE TABLE `message_boxes` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `sender_user_id` int(11) NOT NULL,
  `message_box_type` tinyint(4) NOT NULL,
  `peer_type` tinyint(4) NOT NULL DEFAULT '0',
  `peer_id` int(11) NOT NULL,
  `pts` int(11) NOT NULL,
  `message_id` int(11) NOT NULL,
  `media_unread` tinyint(4) NOT NULL DEFAULT '0',
  `state` tinyint(4) NOT NULL,
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `orgs`
--

CREATE TABLE `orgs` (
  `org_id` int(11) NOT NULL,
  `account_name` varchar(64) NOT NULL,
  `passwd` char(32) NOT NULL,
  `org_name` varchar(256) NOT NULL,
  `mail` varchar(64) NOT NULL,
  `mobile` varchar(32) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` int(11) NOT NULL,
  `updated_at` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `photo_datas`
--

CREATE TABLE `photo_datas` (
  `id` int(11) NOT NULL,
  `file_id` bigint(20) NOT NULL,
  `photo_type` tinyint(4) NOT NULL,
  `dc_id` int(11) NOT NULL,
  `volume_id` bigint(20) NOT NULL,
  `local_id` int(11) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `width` int(11) NOT NULL,
  `height` int(11) NOT NULL,
  `file_size` int(11) NOT NULL,
  `bytes` mediumblob NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `reports`
--

CREATE TABLE `reports` (
  `id` int(11) NOT NULL,
  `auth_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `peer_type` int(11) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `reason` tinyint(4) NOT NULL,
  `content` varchar(10000) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `secret_messages`
--

CREATE TABLE `secret_messages` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `peer_type` int(11) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `random_id` bigint(20) NOT NULL,
  `message_content_header` int(11) NOT NULL,
  `message_content_data` blob NOT NULL,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `seq_updates_ngen`
--

CREATE TABLE `seq_updates_ngen` (
  `id` bigint(20) NOT NULL,
  `seq_name` varchar(255) NOT NULL,
  `seq` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `tmp_passwords`
--

CREATE TABLE `tmp_passwords` (
  `id` int(11) NOT NULL,
  `auth_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `password_hash` varchar(512) NOT NULL,
  `period` int(11) NOT NULL,
  `tmp_password` varchar(512) NOT NULL,
  `valid_until` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `first_name` varchar(255) NOT NULL DEFAULT '',
  `last_name` varchar(255) NOT NULL DEFAULT '',
  `username` varchar(255) NOT NULL,
  `phone` varchar(32) NOT NULL,
  `country_code` varchar(2) NOT NULL,
  `bio` varchar(255) NOT NULL,
  `about` varchar(512) NOT NULL DEFAULT '',
  `state` int(11) NOT NULL DEFAULT '0',
  `is_bot` tinyint(1) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `deleted_reason` varchar(500) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `user_contacts`
--

CREATE TABLE `user_contacts` (
  `id` int(11) NOT NULL,
  `owner_user_id` int(11) NOT NULL,
  `contact_user_id` int(11) NOT NULL,
  `is_blocked` tinyint(1) NOT NULL DEFAULT '0',
  `date2` int(11) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `user_dialogs`
--

CREATE TABLE `user_dialogs` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `peer_type` tinyint(4) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `is_pinned` tinyint(1) NOT NULL DEFAULT '0',
  `top_message` int(11) NOT NULL DEFAULT '0',
  `read_inbox_max_id` int(11) NOT NULL DEFAULT '0',
  `read_outbox_max_id` int(11) NOT NULL DEFAULT '0',
  `unread_count` int(11) NOT NULL DEFAULT '0',
  `unread_mentions_count` int(11) NOT NULL DEFAULT '0',
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `user_imported_contacts`
--

CREATE TABLE `user_imported_contacts` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `auth_id` bigint(20) NOT NULL,
  `client_id` bigint(20) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `state` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='手机通信录';

-- --------------------------------------------------------

--
-- 表的结构 `user_notify_settings`
--

CREATE TABLE `user_notify_settings` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `peer_type` tinyint(4) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `show_previews` tinyint(1) NOT NULL DEFAULT '0',
  `silent` tinyint(1) NOT NULL DEFAULT '0',
  `mute_until` int(11) NOT NULL DEFAULT '0',
  `sound` varchar(255) NOT NULL DEFAULT 'default',
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `user_presences`
--

CREATE TABLE `user_presences` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `auth_id` bigint(20) NOT NULL DEFAULT '0',
  `last_seen_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `last_seen_ip` varchar(64) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `user_privacys`
--

CREATE TABLE `user_privacys` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `password` varchar(64) NOT NULL DEFAULT '',
  `recovery_mail` varchar(64) NOT NULL DEFAULT '',
  `status_timestamp` tinyint(4) NOT NULL DEFAULT '0',
  `chat_invite` tinyint(4) NOT NULL DEFAULT '0',
  `phone_call` tinyint(4) NOT NULL DEFAULT '0',
  `ttl` int(11) NOT NULL DEFAULT '0',
  `ttl_created_at` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `user_sequence`
--

CREATE TABLE `user_sequence` (
  `id` bigint(20) NOT NULL,
  `user_id` varchar(255) NOT NULL,
  `seq` bigint(20) NOT NULL DEFAULT '0',
  `header` bigint(20) NOT NULL,
  `data` blob,
  `created_at` bigint(20) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `apps`
--
ALTER TABLE `apps`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `api_id` (`api_id`);

--
-- Indexes for table `app_configs`
--
ALTER TABLE `app_configs`
  ADD PRIMARY KEY (`app_id`);

--
-- Indexes for table `app_ios_push_certs`
--
ALTER TABLE `app_ios_push_certs`
  ADD PRIMARY KEY (`cert_id`);

--
-- Indexes for table `app_keys`
--
ALTER TABLE `app_keys`
  ADD PRIMARY KEY (`app_id`);

--
-- Indexes for table `auths`
--
ALTER TABLE `auths`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_id` (`auth_id`);

--
-- Indexes for table `auth_keys`
--
ALTER TABLE `auth_keys`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_id` (`auth_id`);

--
-- Indexes for table `auth_phone_transactions`
--
ALTER TABLE `auth_phone_transactions`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `transaction_hash` (`transaction_hash`);

--
-- Indexes for table `auth_salts`
--
ALTER TABLE `auth_salts`
  ADD PRIMARY KEY (`id`),
  ADD KEY `auth` (`auth_id`);

--
-- Indexes for table `auth_users`
--
ALTER TABLE `auth_users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_id` (`auth_id`);

--
-- Indexes for table `channels`
--
ALTER TABLE `channels`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `channel_users`
--
ALTER TABLE `channel_users`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `chats`
--
ALTER TABLE `chats`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `title_changer_user_id` (`title_changer_user_id`,`title_change_random_id`),
  ADD UNIQUE KEY `avatar_changer_user_id` (`avatar_changer_user_id`,`avatar_change_random_id`);

--
-- Indexes for table `chat_participants`
--
ALTER TABLE `chat_participants`
  ADD PRIMARY KEY (`id`),
  ADD KEY `chat_id` (`chat_id`);

--
-- Indexes for table `chat_users`
--
ALTER TABLE `chat_users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `chat_id` (`chat_id`);

--
-- Indexes for table `client_updates_state`
--
ALTER TABLE `client_updates_state`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_key_id` (`auth_key_id`,`user_id`),
  ADD KEY `auth_key_id_2` (`auth_key_id`,`user_id`);

--
-- Indexes for table `devices`
--
ALTER TABLE `devices`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `files`
--
ALTER TABLE `files`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `creator_user_id` (`creator_user_id`,`file_id`);

--
-- Indexes for table `file_parts`
--
ALTER TABLE `file_parts`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `file_id_2` (`file_id`,`file_part`);

--
-- Indexes for table `messages`
--
ALTER TABLE `messages`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `message_boxes`
--
ALTER TABLE `message_boxes`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `orgs`
--
ALTER TABLE `orgs`
  ADD PRIMARY KEY (`org_id`),
  ADD UNIQUE KEY `account_name` (`account_name`);

--
-- Indexes for table `photo_datas`
--
ALTER TABLE `photo_datas`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `reports`
--
ALTER TABLE `reports`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `secret_messages`
--
ALTER TABLE `secret_messages`
  ADD PRIMARY KEY (`id`),
  ADD KEY `message_content_header` (`message_content_header`);

--
-- Indexes for table `seq_updates_ngen`
--
ALTER TABLE `seq_updates_ngen`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `seq_name` (`seq_name`);

--
-- Indexes for table `tmp_passwords`
--
ALTER TABLE `tmp_passwords`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `phone` (`phone`);

--
-- Indexes for table `user_contacts`
--
ALTER TABLE `user_contacts`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `owner_user_id` (`owner_user_id`,`contact_user_id`);

--
-- Indexes for table `user_dialogs`
--
ALTER TABLE `user_dialogs`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`peer_type`,`peer_id`);

--
-- Indexes for table `user_imported_contacts`
--
ALTER TABLE `user_imported_contacts`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_notify_settings`
--
ALTER TABLE `user_notify_settings`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`peer_type`,`peer_id`);

--
-- Indexes for table `user_presences`
--
ALTER TABLE `user_presences`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`auth_id`),
  ADD KEY `user_id_2` (`user_id`,`last_seen_at`);

--
-- Indexes for table `user_privacys`
--
ALTER TABLE `user_privacys`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`);

--
-- Indexes for table `user_sequence`
--
ALTER TABLE `user_sequence`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `seq` (`seq`,`user_id`),
  ADD UNIQUE KEY `id` (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `apps`
--
ALTER TABLE `apps`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `app_configs`
--
ALTER TABLE `app_configs`
  MODIFY `app_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `app_ios_push_certs`
--
ALTER TABLE `app_ios_push_certs`
  MODIFY `cert_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `auths`
--
ALTER TABLE `auths`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=77;

--
-- 使用表AUTO_INCREMENT `auth_keys`
--
ALTER TABLE `auth_keys`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=124;

--
-- 使用表AUTO_INCREMENT `auth_phone_transactions`
--
ALTER TABLE `auth_phone_transactions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `auth_salts`
--
ALTER TABLE `auth_salts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=37;

--
-- 使用表AUTO_INCREMENT `auth_users`
--
ALTER TABLE `auth_users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=20;

--
-- 使用表AUTO_INCREMENT `channels`
--
ALTER TABLE `channels`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `channel_users`
--
ALTER TABLE `channel_users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `chats`
--
ALTER TABLE `chats`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=41;

--
-- 使用表AUTO_INCREMENT `chat_participants`
--
ALTER TABLE `chat_participants`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=69;

--
-- 使用表AUTO_INCREMENT `chat_users`
--
ALTER TABLE `chat_users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=77;

--
-- 使用表AUTO_INCREMENT `client_updates_state`
--
ALTER TABLE `client_updates_state`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `devices`
--
ALTER TABLE `devices`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=84;

--
-- 使用表AUTO_INCREMENT `files`
--
ALTER TABLE `files`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;

--
-- 使用表AUTO_INCREMENT `file_parts`
--
ALTER TABLE `file_parts`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=61;

--
-- 使用表AUTO_INCREMENT `messages`
--
ALTER TABLE `messages`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=310;

--
-- 使用表AUTO_INCREMENT `message_boxes`
--
ALTER TABLE `message_boxes`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=685;

--
-- 使用表AUTO_INCREMENT `orgs`
--
ALTER TABLE `orgs`
  MODIFY `org_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `photo_datas`
--
ALTER TABLE `photo_datas`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=77;

--
-- 使用表AUTO_INCREMENT `reports`
--
ALTER TABLE `reports`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `secret_messages`
--
ALTER TABLE `secret_messages`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `seq_updates_ngen`
--
ALTER TABLE `seq_updates_ngen`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- 使用表AUTO_INCREMENT `tmp_passwords`
--
ALTER TABLE `tmp_passwords`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `user_contacts`
--
ALTER TABLE `user_contacts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- 使用表AUTO_INCREMENT `user_dialogs`
--
ALTER TABLE `user_dialogs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=29;

--
-- 使用表AUTO_INCREMENT `user_imported_contacts`
--
ALTER TABLE `user_imported_contacts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `user_notify_settings`
--
ALTER TABLE `user_notify_settings`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `user_presences`
--
ALTER TABLE `user_presences`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `user_privacys`
--
ALTER TABLE `user_privacys`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `user_sequence`
--
ALTER TABLE `user_sequence`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=129;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
