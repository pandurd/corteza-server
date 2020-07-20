package schema

// Note:
// Create table instructions should always be up-to-date
// and should NOT rely on incremental upgrades
//
// @todo reformat create statements to match sql_users.go
//       uppercase reserved words, aligned columns (2, 30, 50, 60, 100)

const messagingAttachment = `messaging_attachment`
const messagingAttachmentCreateSQL = `CREATE TABLE ` + messagingAttachment + ` (
  id bigint unsigned NOT NULL,
  rel_user bigint unsigned NOT NULL,
  url varchar(512) DEFAULT NULL,
  preview_url varchar(512) DEFAULT NULL,
  size int unsigned DEFAULT NULL,
  mimetype varchar(255) DEFAULT NULL,
  name text,
  meta json DEFAULT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT NULL,
  deleted_at datetime DEFAULT NULL,
  
  PRIMARY KEY (id)
) ` + pfxCreateTable

const messagingChannel = `messaging_channel`
const messagingChannelCreateSQL = `CREATE TABLE ` + messagingChannel + ` (
  id bigint unsigned NOT NULL,
  name text NOT NULL,
  topic text NOT NULL,
  meta json NOT NULL,
  type enum('private','public','group') DEFAULT NULL,
  membership_policy enum('featured','forced','') NOT NULL DEFAULT '',
  rel_organisation bigint unsigned NOT NULL,
  rel_creator bigint unsigned NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT NULL,
  archived_at datetime DEFAULT NULL,
  deleted_at datetime DEFAULT NULL,
  rel_last_message bigint unsigned NOT NULL DEFAULT '0',
  
  PRIMARY KEY (id)
) ` + pfxCreateTable

const messagingChannelMember = `messaging_channel_member`
const messagingChannelMemberCreateSQL = `CREATE TABLE ` + messagingChannelMember + ` (
  rel_channel bigint unsigned NOT NULL,
  rel_user bigint unsigned NOT NULL,
  type enum('owner','member','invitee') DEFAULT NULL,
  flag enum('pinned','hidden','ignored','') NOT NULL DEFAULT '',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT NULL,
  
  PRIMARY KEY (rel_channel, rel_user)
) ` + pfxCreateTable

const messagingMention = `messaging_mention`
const messagingMentionCreateSQL = `CREATE TABLE ` + messagingMention + ` (
  id bigint unsigned NOT NULL,
  rel_channel bigint unsigned NOT NULL,
  rel_message bigint unsigned NOT NULL,
  rel_user bigint unsigned NOT NULL,
  rel_mentioned_by bigint unsigned NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  
  PRIMARY KEY (id),
  KEY lookup_mentions (rel_mentioned_by)
) ` + pfxCreateTable

const messagingMessage = `messaging_message`
const messagingMessageCreateSQL = `CREATE TABLE ` + messagingMessage + ` (
  id bigint unsigned NOT NULL,
  type mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  message mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  meta json DEFAULT NULL,
  rel_user bigint unsigned NOT NULL,
  rel_channel bigint unsigned NOT NULL,
  reply_to bigint unsigned NOT NULL DEFAULT '0',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT NULL,
  deleted_at datetime DEFAULT NULL,
  replies int unsigned NOT NULL DEFAULT '0',
  
  PRIMARY KEY (id)
) ` + pfxCreateTable

const messagingMessageAttachment = `messaging_message_attachment`
const messagingMessageAttachmentCreateSQL = `CREATE TABLE ` + messagingMessageAttachment + ` (
  rel_message bigint unsigned NOT NULL,
  rel_attachment bigint unsigned NOT NULL,
  
  PRIMARY KEY (rel_message)
) ` + pfxCreateTable

const messagingMessageFlag = `messaging_message_flag`
const messagingMessageFlagCreateSQL = `CREATE TABLE ` + messagingMessageFlag + ` (
  id bigint unsigned NOT NULL,
  rel_channel bigint unsigned NOT NULL,
  rel_message bigint unsigned NOT NULL,
  rel_user bigint unsigned NOT NULL,
  flag text,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  
  PRIMARY KEY (id)
) ` + pfxCreateTable

const messagingPermissionRules = `messaging_permission_rules`
const messagingPermissionRulesCreateSQL = `CREATE TABLE ` + messagingPermissionRules + ` (
  rel_role bigint unsigned NOT NULL,
  resource varchar(128) NOT NULL,
  operation varchar(128) NOT NULL,
  access tinyint(1) NOT NULL,
  
  PRIMARY KEY (rel_role, resource, operation)
) ` + pfxCreateTable

const messagingSettings = `messaging_settings`
const messagingSettingsCreateSQL = `CREATE TABLE ` + messagingSettings + ` (
  rel_owner bigint unsigned NOT NULL DEFAULT '0' COMMENT 'Value owner, 0 for global settings',
  name varchar(200) NOT NULL COMMENT 'Unique set of setting keys',
  value json DEFAULT NULL COMMENT 'Setting value',
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'When was the value updated',
  updated_by bigint unsigned NOT NULL DEFAULT '0' COMMENT 'Who created/updated the value',
  
  PRIMARY KEY (name, rel_owner)
) ` + pfxCreateTable

const messagingUnread = `messaging_unread`
const messagingUnreadCreateSQL = `CREATE TABLE ` + messagingUnread + ` (
  rel_channel bigint unsigned NOT NULL DEFAULT '0',
  rel_reply_to bigint unsigned NOT NULL,
  rel_user bigint unsigned NOT NULL DEFAULT '0',
  count int unsigned NOT NULL DEFAULT '0',
  rel_last_message bigint unsigned NOT NULL DEFAULT '0',
  
  PRIMARY KEY (rel_channel, rel_reply_to, rel_user)
) ` + pfxCreateTable
