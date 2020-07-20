package schema

// Note:
// Create table instructions should always be up-to-date
// and should NOT rely on incremental upgrades

const sysUser = `sys_user`
const sysUserCreateSQL = `
CREATE TABLE ` + sysUser + ` (
  id                         BIGINT UNSIGNED     NOT NULL,
  email                      TEXT                NOT NULL,
  email_confirmed            TINYINT(1)          NOT NULL  DEFAULT '0',
  username                   TEXT                NOT NULL,
  name                       TEXT                NOT NULL,
  handle                     TEXT                NOT NULL,
  kind                       VARCHAR(8)          NOT NULL  DEFAULT '',
  meta                       JSON                NOT NULL,
  rel_organisation           BIGINT UNSIGNED     NOT NULL,
  rel_user_id                BIGINT UNSIGNED     NOT NULL,
  created_at                 DATETIME            NOT NULL  DEFAULT CURRENT_TIMESTAMP,
  updated_at                 DATETIME                      DEFAULT NULL,
  suspended_at               DATETIME                      DEFAULT NULL,
  deleted_at                 DATETIME                      DEFAULT NULL,

  PRIMARY KEY (id)
) ` + pfxCreateTable

const sysCredentials = `sys_credentials`
const sysCredentialsCreateSQL = `
CREATE TABLE ` + sysCredentials + ` (
  id                         BIGINT UNSIGNED     NOT NULL,
  rel_owner                  BIGINT UNSIGNED     NOT NULL,
  label                      TEXT                NOT NULL                                          COMMENT 'something we can differentiate credentials by',
  kind                       VARCHAR(128)        NOT NULL                                          COMMENT 'hash, facebook, gplus, github, linkedin ...',
  credentials                TEXT                NOT NULL                                          COMMENT 'encrypted passwords, secrets, social profile ID',
  meta                       JSON                NOT NULL,
  expires_at                 DATETIME                      DEFAULT NULL,
  created_at                 DATETIME            NOT NULL  DEFAULT CURRENT_TIMESTAMP,
  updated_at                 DATETIME                      DEFAULT NULL,
  deleted_at                 DATETIME                      DEFAULT NULL,
  last_used_at               DATETIME                      DEFAULT NULL,
  
  PRIMARY KEY (id),
  KEY (rel_owner)
) ` + pfxCreateTable

const sysRole = `sys_role`
const sysRoleCreateSQL = `
CREATE TABLE ` + sysRole + ` (
  id                         BIGINT UNSIGNED     NOT NULL,
  name                       TEXT                NOT NULL,
  handle                     TEXT                NOT NULL,
  created_at                 DATETIME            NOT NULL  DEFAULT CURRENT_TIMESTAMP,
  updated_at                 DATETIME                      DEFAULT NULL,
  archived_at                DATETIME                      DEFAULT NULL,
  deleted_at                 DATETIME                      DEFAULT NULL,
  
  PRIMARY KEY (id)
) ` + pfxCreateTable

const sysRoleMember = `sys_role_member`
const sysRoleMemberCreateSQL = `
CREATE TABLE ` + sysRoleMember + ` (
  rel_role                   BIGINT UNSIGNED     NOT NULL,
  rel_user                   BIGINT UNSIGNED     NOT NULL,
  
  PRIMARY KEY (rel_role, rel_user)
) ` + pfxCreateTable
