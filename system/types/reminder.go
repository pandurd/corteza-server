package types

import (
	"github.com/cortezaproject/corteza-server/store"
	"time"

	"github.com/jmoiron/sqlx/types"
)

type (
	Reminder struct {
		ID          uint64         `json:"reminderID,string" db:"id"`
		Resource    string         `json:"resource" db:"resource"`
		Payload     types.JSONText `json:"payload" db:"payload"`
		SnoozeCount uint           `json:"snoozeCount" db:"snooze_count"`

		AssignedTo uint64    `json:"assignedTo,string" db:"assigned_to"`
		AssignedBy uint64    `json:"assignedBy,string" db:"assigned_by"`
		AssignedAt time.Time `json:"assignedAt" db:"assigned_at"`

		DismissedBy uint64     `json:"dismissedBy,string" db:"dismissed_by"`
		DismissedAt *time.Time `json:"dismissedAt" db:"dismissed_at"`

		RemindAt *time.Time `json:"remindAt" db:"remind_at"`

		CreatedAt time.Time  `json:"createdAt,omitempty" db:"created_at"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty" db:"updated_at"`
		DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
	}

	ReminderFilter struct {
		ReminderID       []uint64   `json:"reminderID"`
		Resource         string     `json:"resource"`
		AssignedTo       uint64     `json:"assignedTo,uint64"`
		ScheduledFrom    *time.Time `json:"scheduledFrom"`
		ScheduledUntil   *time.Time `json:"scheduledUntil"`
		ExcludeDismissed bool       `json:"excludeDismissed"`
		ScheduledOnly    bool       `json:"scheduledOnly"`

		// Check fn is called by store backend for each resource found function can
		// modify the resource and return false if store should not return it
		//
		// Store then loads additional resources to satisfy the paging parameters
		Check func(*Reminder) (bool, error) `json:"-"`

		// Standard helpers for paging and sorting
		store.Sorting
		store.Paging
	}
)
