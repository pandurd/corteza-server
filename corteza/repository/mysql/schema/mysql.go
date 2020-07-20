package schema

// MySQL specific prefixes, sql
// templates, functions and other helpers

import (
	"fmt"
	. "github.com/cortezaproject/corteza-server/corteza/repository/internal/provisioner"
	_ "github.com/go-sql-driver/mysql"
)

type (
	// Holds table structure
	tableColumn struct {
		Field   string  `db:"Field"`
		Type    string  `db:"Type"`
		Null    string  `db:"Null"`
		Key     string  `db:"Key"`
		Default *string `db:"Default"`
		Extra   string  `db:"Extra"`
	}
)

// Engine, charset are used on every mysql table
const (
	pfxCreateTable = `ENGINE=InnoDB DEFAULT CHARSET=utf8`
	sqlTableExists = `SELECT COUNT(*) FROM information_schema.TABLES WHERE (TABLE_SCHEMA = ?) AND (TABLE_NAME = ?)`
	fmtDropColumn  = `ALTER TABLE %s DROP COLUMN %s`
	fmtAddColumn   = `ALTER TABLE %s ADD COLUMN %s %s`
)

// utility to simplify table creation
func table(name, create string, ee ...Executor) Executor {
	return Do(
		Label("provisioning mysql database table "+name),
		IfElse(
			tableMissing(name),
			Do(
				func(s *Provisioner) error {
					if err := rawSqlExec(create); err != nil {
						return err
					}

					s.Log("created\n")
					return nil
				},
			),
			Do(ee...),
		),
	)
}

// Returns Tester fn that will
// verify if table is present or missing
func tableMissing(table string) Tester {
	return func(s *Provisioner) (bool, error) {
		// @todo implement
		var count int
		if err := db.Get(&count, sqlTableExists, dbName, table); err != nil {
			return false, err
		} else {
			return count == 0, nil
		}
	}
}

// Returns Executor fn that removes column (if exists) from a table
func dropColumn(table, column string) Executor {
	return func(s *Provisioner) error {
		if tt, err := getTableColumns(table); err != nil || getColumn(tt, column) == nil {
			return err
		}

		if _, err := db.Exec(fmt.Sprintf(fmtDropColumn, table, column)); err != nil {
			return err
		}

		s.Log("column %s.%s dropped\n", table, column)
		return nil
	}
}

// Returns Executor fn that adds column
func addColumn(table, column, spec string) Executor {
	return func(s *Provisioner) error {
		if tt, err := getTableColumns(table); err != nil || getColumn(tt, column) != nil {
			return err
		}

		if _, err := db.Exec(fmt.Sprintf(fmtAddColumn, table, column, spec)); err != nil {
			return err
		}

		s.Log("column %s.%s added\n", table, column)
		return nil
	}
}

// Returns all table's columns
func getTableColumns(name string) ([]*tableColumn, error) {
	tt := make([]*tableColumn, 0)

	if err := db.Select(&tt, "DESCRIBE "+name); err != nil {
		return nil, err
	}

	return tt, nil
}

// Searches for a column by it's name in the list of columns
func getColumn(tt []*tableColumn, name string) *tableColumn {
	for _, t := range tt {
		if t.Field == name {
			return t
		}
	}

	return nil
}

// Executes one or more SQL commands
func rawSqlExec(ss ...string) error {
	for _, sql := range ss {
		if _, err := db.Exec(sql); err != nil {
			return err
		}
	}

	return nil
}

// Returns Executor fn that calls rawSqlExec
func execSql(ss ...string) Executor {
	return func(s *Provisioner) error {
		return rawSqlExec(ss...)
	}
}
