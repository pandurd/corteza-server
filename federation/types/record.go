package types

import (
	"time"

	composeTypes "github.com/cortezaproject/corteza-server/compose/types"
)

type (
	FederatedRecord struct {
		ID       uint64 `json:"id,string"`
		ModuleID uint64 `json:"mid,string"`

		Values composeTypes.RecordValueSet `json:"v"`

		NamespaceID uint64 `json:"ns,string"`

		OwnedBy   uint64     `json:"o,string"`
		CreatedAt time.Time  `json:"ca"`
		CreatedBy uint64     `json:"cb,string"`
		UpdatedAt *time.Time `json:"ua"`
		UpdatedBy uint64     `json:"ub,string"`
		DeletedAt *time.Time `json:"da"`
		DeletedBy uint64     `json:"db,string"`
	}

	FederatedRecordSet []*FederatedRecord
)
