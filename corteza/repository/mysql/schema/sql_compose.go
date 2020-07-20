package schema

// Note:
// Create table instructions should always be up-to-date
// and should NOT rely on incremental upgrades
//
// @todo reformat create statements to match sql_users.go
//       uppercase reserved words, aligned columns (2, 30, 50, 60, 100)

const composeAttachment = `compose_attachment`
const composeAttachmentCreateSQL = `CREATE TABLE ` + composeAttachment + ` (
  id bigint unsigned NOT NULL,
  rel_namespace bigint unsigned NOT NULL,
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
  
  PRIMARY KEY (id),
  KEY rel_namespace (rel_namespace)
) ` + pfxCreateTable

const composeChart = `compose_chart`
const composeChartCreateSQL = `CREATE TABLE ` + composeChart + ` (
  id bigint unsigned NOT NULL,
  handle varchar(200) NOT NULL,
  rel_namespace bigint unsigned NOT NULL,
  name varchar(64) NOT NULL COMMENT 'The name of the chart',
  config json NOT NULL COMMENT 'Chart & reporting configuration',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT NULL,
  deleted_at datetime DEFAULT NULL,
  
  PRIMARY KEY (id),
  KEY rel_namespace (rel_namespace)
) ` + pfxCreateTable

const composeModule = `compose_module`
const composeModuleCreateSQL = `CREATE TABLE ` + composeModule + ` (
  id bigint unsigned NOT NULL,
  handle varchar(200) NOT NULL,
  rel_namespace bigint unsigned NOT NULL,
  name varchar(64) NOT NULL COMMENT 'The name of the module',
  json json NOT NULL COMMENT 'List of field definitions for the module',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT NULL,
  deleted_at datetime DEFAULT NULL,
  
  PRIMARY KEY (id),
  KEY rel_namespace (rel_namespace)
) ` + pfxCreateTable

const composeModuleField = `compose_module_field`
const composeModuleFieldCreateSQL = `CREATE TABLE ` + composeModuleField + ` (
  id bigint unsigned NOT NULL,
  rel_module bigint unsigned NOT NULL,
  place tinyint unsigned NOT NULL,
  kind varchar(64) NOT NULL COMMENT 'The type of the form input field',
  options json NOT NULL COMMENT 'Options in JSON format.',
  default_value json DEFAULT NULL COMMENT 'Default value as a record value set.',
  name varchar(64) NOT NULL COMMENT 'The name of the field in the form',
  label varchar(255) NOT NULL COMMENT 'The label of the form input',
  is_private tinyint(1) NOT NULL COMMENT 'Contains personal/sensitive data?',
  is_required tinyint(1) NOT NULL,
  is_visible tinyint(1) NOT NULL,
  is_multi tinyint(1) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT NULL,
  deleted_at datetime DEFAULT NULL,
  
  PRIMARY KEY (id),
  UNIQUE KEY uid_compose_module_field_place (rel_module, place),
  UNIQUE KEY uid_compose_module_field_name (rel_module, name)
) ` + pfxCreateTable

const composeNamespace = `compose_namespace`
const composeNamespaceCreateSQL = `CREATE TABLE ` + composeNamespace + ` (
  id bigint unsigned NOT NULL,
  name varchar(64) NOT NULL COMMENT 'Name',
  slug varchar(64) NOT NULL COMMENT 'URL slug',
  enabled tinyint(1) NOT NULL COMMENT 'Is namespace enabled?',
  meta json NOT NULL COMMENT 'Meta data',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT NULL,
  deleted_at datetime DEFAULT NULL,
  
  PRIMARY KEY (id)
) ` + pfxCreateTable

const composePage = `compose_page`
const composePageCreateSQL = `CREATE TABLE ` + composePage + ` (
  id bigint unsigned NOT NULL COMMENT 'Page ID',
  handle varchar(200) NOT NULL,
  rel_namespace bigint unsigned NOT NULL,
  self_id bigint unsigned NOT NULL COMMENT 'Parent Page ID',
  rel_module bigint unsigned NOT NULL DEFAULT '0',
  title varchar(255) NOT NULL COMMENT 'Title (required)',
  description text NOT NULL COMMENT 'Description',
  blocks json NOT NULL COMMENT 'JSON array of blocks for the page',
  visible tinyint NOT NULL COMMENT 'Is page visible in navigation?',
  weight int NOT NULL COMMENT 'Order for navigation',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT NULL,
  deleted_at datetime DEFAULT NULL,
  
  PRIMARY KEY (id) USING BTREE,
  KEY module_id (rel_module),
  KEY self_id (self_id),
  KEY rel_namespace (rel_namespace)
) ` + pfxCreateTable

const composePermissionRules = `compose_permission_rules`
const composePermissionRulesCreateSQL = `CREATE TABLE ` + composePermissionRules + ` (
  rel_role bigint unsigned NOT NULL,
  resource varchar(128) NOT NULL,
  operation varchar(128) NOT NULL,
  access tinyint(1) NOT NULL,
  
  PRIMARY KEY (rel_role, resource, operation)
) ` + pfxCreateTable

const composeRecord = `compose_record`
const composeRecordCreateSQL = `CREATE TABLE ` + composeRecord + ` (
  id bigint unsigned NOT NULL,
  rel_namespace bigint unsigned NOT NULL,
  module_id bigint unsigned NOT NULL,
  owned_by bigint unsigned NOT NULL DEFAULT '0',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT NULL,
  deleted_at datetime DEFAULT NULL,
  created_by bigint unsigned NOT NULL DEFAULT '0',
  updated_by bigint unsigned NOT NULL DEFAULT '0',
  deleted_by bigint unsigned NOT NULL DEFAULT '0',
  
  PRIMARY KEY (id, module_id),
  KEY user_id (owned_by),
  KEY rel_namespace (rel_namespace)
) ` + pfxCreateTable

const composeRecordValue = `compose_record_value`
const composeRecordValueCreateSQL = `CREATE TABLE ` + composeRecordValue + ` (
  record_id bigint NOT NULL,
  name varchar(64) NOT NULL,
  value longtext,
  ref bigint unsigned NOT NULL DEFAULT '0',
  deleted_at datetime DEFAULT NULL,
  place int unsigned NOT NULL DEFAULT '0',
  
  PRIMARY KEY (record_id, name, place),
  KEY crm_record_value_ref (ref)
) ` + pfxCreateTable

const composeSettings = `compose_settings`
const composeSettingsCreateSQL = `CREATE TABLE ` + composeSettings + ` (
  rel_owner bigint unsigned NOT NULL DEFAULT '0' COMMENT 'Value owner, 0 for global settings',
  name varchar(200) NOT NULL COMMENT 'Unique set of setting keys',
  value json DEFAULT NULL COMMENT 'Setting value',
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'When was the value updated',
  updated_by bigint unsigned NOT NULL DEFAULT '0' COMMENT 'Who created/updated the value',
  
  PRIMARY KEY (name, rel_owner)
) ` + pfxCreateTable
