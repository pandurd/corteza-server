package schema

// Note:
// Create table instructions should always be up-to-date
// and should NOT rely on incremental upgrades
//
// @todo reformat create statements to match sql_users.go
//       uppercase reserved words, aligned columns (2, 30, 50, 60, 100)

const sysActionlog = `sys_actionlog`
const sysActionlogCreateSQL = `
CREATE TABLE ` + sysActionlog + ` (
  ts datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  actor_ip_addr varchar(15) NOT NULL,
  actor_id bigint unsigned DEFAULT NULL,
  request_origin varchar(32) NOT NULL,
  request_id varchar(64) NOT NULL,
  resource varchar(128) NOT NULL,
  action varchar(64) NOT NULL,
  error varchar(64) NOT NULL,
  severity smallint NOT NULL,
  description text,
  meta json DEFAULT NULL,

  KEY ts (ts DESC),
  KEY request_origin (request_origin),
  KEY actor_id (actor_id),
  KEY resource (resource),
  KEY action (action)
) ` + pfxCreateTable

const sysApplication = `sys_application`
const sysApplicationCreateSQL = `
CREATE TABLE ` + sysApplication + ` (
  id bigint unsigned NOT NULL,
  rel_owner bigint unsigned NOT NULL,
  name text NOT NULL COMMENT 'something we can differentiate application by',
  enabled tinyint(1) NOT NULL,
  unify json DEFAULT NULL COMMENT 'unify specific settings',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT NULL,
  deleted_at datetime DEFAULT NULL,

  PRIMARY KEY (id)
) ` + pfxCreateTable

const sysAttachment = `sys_attachment`
const sysAttachmentCreateSQL = `
CREATE TABLE ` + sysAttachment + ` (
  id bigint unsigned NOT NULL,
  rel_owner bigint unsigned NOT NULL,
  kind varchar(32) NOT NULL,
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

const sysOrganisation = `sys_organisation`
const sysOrganisationCreateSQL = `
CREATE TABLE ` + sysOrganisation + ` (
 id bigint unsigned NOT NULL,
 fqn text NOT NULL,
 name text NOT NULL,
 created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at datetime DEFAULT NULL,
 archived_at datetime DEFAULT NULL,
 deleted_at datetime DEFAULT NULL,

 PRIMARY KEY (id)
)
` + pfxCreateTable

const sysPermissionRules = `sys_permission_rules`
const sysPermissionRulesCreateSQL = `
CREATE TABLE ` + sysPermissionRules + ` (
  rel_role bigint unsigned NOT NULL,
  resource varchar(128) NOT NULL,
  operation varchar(128) NOT NULL,
  access tinyint(1) NOT NULL,

  PRIMARY KEY (rel_role, resource, operation)
) ` + pfxCreateTable

const sysReminder = `sys_reminder`
const sysReminderCreateSQL = `
CREATE TABLE ` + sysReminder + ` (
  id bigint unsigned NOT NULL,
  resource varchar(128) NOT NULL COMMENT 'Resource, that this reminder is bound to',
  payload json NOT NULL COMMENT 'Payload for this reminder',
  snooze_count int NOT NULL DEFAULT '0' COMMENT 'Number of times this reminder was snoozed',
  assigned_to bigint unsigned NOT NULL DEFAULT '0' COMMENT 'Assignee for this reminder',
  assigned_by bigint unsigned NOT NULL DEFAULT '0' COMMENT 'User that assigned this reminder',
  assigned_at datetime NOT NULL COMMENT 'When the reminder was assigned',
  dismissed_by bigint unsigned NOT NULL DEFAULT '0' COMMENT 'User that dismissed this reminder',
  dismissed_at datetime DEFAULT NULL COMMENT 'Time the reminder was dismissed',
  remind_at datetime DEFAULT NULL COMMENT 'Time the user should be reminded',
  created_by bigint unsigned NOT NULL DEFAULT '0',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_by bigint unsigned NOT NULL DEFAULT '0',
  updated_at datetime DEFAULT NULL,
  deleted_by bigint unsigned NOT NULL DEFAULT '0',
  deleted_at datetime DEFAULT NULL,

  PRIMARY KEY (id)
) ` + pfxCreateTable

const sysSettings = `sys_settings`
const sysSettingsCreateSQL = `
CREATE TABLE ` + sysSettings + ` (
  rel_owner bigint unsigned NOT NULL DEFAULT '0' COMMENT 'Value owner, 0 for global settings',
  name varchar(200) NOT NULL COMMENT 'Unique set of setting keys',
  value json DEFAULT NULL COMMENT 'Setting value',
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'When was the value updated',
  updated_by bigint unsigned NOT NULL DEFAULT '0' COMMENT 'Who created/updated the value',

  PRIMARY KEY (name, rel_owner)
) ` + pfxCreateTable
