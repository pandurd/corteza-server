package rdbms

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/compose/types"
)

func (s Store) convertComposeModuleFieldFilter(f types.ModuleFieldFilter) (query squirrel.SelectBuilder, err error) {
	query = s.QueryComposeModuleFields()

	if len(f.ModuleID) == 0 {
		err = fmt.Errorf("can not search for module fields without module IDs")
		return
	}

	query = query.Where("cmd.rel_module = ?", f.ModuleID)

	return
}
