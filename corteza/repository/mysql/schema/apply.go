package schema

import (
	"context"
	"fmt"
	. "github.com/cortezaproject/corteza-server/corteza/repository/internal/provisioner"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"regexp"
)

// Provisioning is split to 3 standard parts - system, messaging and compose

var (
	// package wide variable for db connection
	//
	// this simplifies our provision utils
	// that us the variable directly
	db *sqlx.DB

	// name of the database
	dbName string
)

func connect(ctx context.Context, dsn string) (err error) {
	mm := regexp.MustCompile(`/([^?]+)`).FindStringSubmatch(dsn)
	if len(mm) == 0 {
		return fmt.Errorf("could not find database name in the dsn %q", dsn)
	} else {
		dbName = mm[1]
	}

	if db, err = sqlx.ConnectContext(ctx, "mysql", dsn); err != nil {
		return
	}

	return
}

func ProvisionSystem(ctx context.Context, dsn string, p Printer) error {
	if db == nil {
		if err := connect(ctx, dsn); err != nil {
			return err
		}
	}

	return New(p).Run(
		table(
			sysUser,
			sysUserCreateSQL,
			// using this for example right now
			//dropColumn(sysUser, "rel_organisation"),
			//dropColumn(sysUser, "rel_user_id"),
			//addColumn(sysRole, "created_by", "BIGINT UNSIGNED NOT NULL DEFAULT 0"),
			//addColumn(sysRole, "updated_by", "BIGINT UNSIGNED NOT NULL DEFAULT 0"),
			//addColumn(sysRole, "deleted_by", "BIGINT UNSIGNED NOT NULL DEFAULT 0"),
		),
		table(sysCredentials, sysCredentialsCreateSQL),
		table(
			sysRole,
			sysRoleCreateSQL,
			//addColumn(sysRole, "owned_by", "BIGINT UNSIGNED NOT NULL DEFAULT 0"),
			//addColumn(sysRole, "created_by", "BIGINT UNSIGNED NOT NULL DEFAULT 0"),
			//addColumn(sysRole, "updated_by", "BIGINT UNSIGNED NOT NULL DEFAULT 0"),
			//addColumn(sysRole, "deleted_by", "BIGINT UNSIGNED NOT NULL DEFAULT 0"),
		),
		table(sysRoleMember, sysRoleMemberCreateSQL),
		table(sysActionlog, sysActionlogCreateSQL),
		table(sysApplication, sysApplicationCreateSQL),
		table(sysAttachment, sysAttachmentCreateSQL),
		table(sysOrganisation, sysOrganisationCreateSQL),
		table(sysReminder, sysReminderCreateSQL),
		table(sysSettings, sysSettingsCreateSQL),
		table(sysPermissionRules, sysPermissionRulesCreateSQL),
	)
}

func ProvisionCompose(ctx context.Context, dsn string, p Printer) error {
	if db == nil {
		if err := connect(ctx, dsn); err != nil {
			return err
		}
	}

	return New(p).Run(
		table(composeChart, composeChartCreateSQL),
		table(composeModule, composeModuleCreateSQL),
		table(composeModuleField, composeModuleFieldCreateSQL),
		table(composeNamespace, composeNamespaceCreateSQL),
		table(composePage, composePageCreateSQL),
		table(composeRecord, composeRecordCreateSQL),
		table(composeRecordValue, composeRecordValueCreateSQL),
		table(composeAttachment, composeAttachmentCreateSQL),
		table(composeSettings, composeSettingsCreateSQL),
		table(composePermissionRules, composePermissionRulesCreateSQL),
	)
}

func ProvisionMessaging(ctx context.Context, dsn string, p Printer) error {
	if db == nil {
		if err := connect(ctx, dsn); err != nil {
			return err
		}
	}

	return New(p).Run(
		table(messagingChannel, messagingChannelCreateSQL),
		table(messagingChannelMember, messagingChannelMemberCreateSQL),
		table(messagingMention, messagingMentionCreateSQL),
		table(messagingMessage, messagingMessageCreateSQL),
		table(messagingMessageAttachment, messagingMessageAttachmentCreateSQL),
		table(messagingMessageFlag, messagingMessageFlagCreateSQL),
		table(messagingUnread, messagingUnreadCreateSQL),
		table(messagingAttachment, messagingAttachmentCreateSQL),
		table(messagingSettings, messagingSettingsCreateSQL),
		table(messagingPermissionRules, messagingPermissionRulesCreateSQL),
	)
}
