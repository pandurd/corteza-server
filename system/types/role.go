package types

import (
	"github.com/cortezaproject/corteza-server/store"
	"time"

	"github.com/cortezaproject/corteza-server/pkg/permissions"
	"github.com/cortezaproject/corteza-server/pkg/rh"
)

type (
	// Role - An organisation may have many roles. Roles may have many channels available. Access to channels may be shared between roles.
	Role struct {
		ID         uint64     `json:"roleID,string"`
		Name       string     `json:"name"`
		Handle     string     `json:"handle"`
		CreatedAt  time.Time  `json:"createdAt,omitempty"`
		UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
		ArchivedAt *time.Time `json:"archivedAt,omitempty"`
		DeletedAt  *time.Time `json:"deletedAt,omitempty"`
	}

	RoleFilter struct {
		RoleID   []uint64 `json:"roleID"`
		MemberID uint64   `json:"memberID"`

		Query string `json:"query"`

		Handle string `json:"handle"`
		Name   string `json:"name"`

		Deleted  rh.FilterState `json:"deleted"`
		Archived rh.FilterState `json:"archived"`

		// Check fn is called by store backend for each resource found function can
		// modify the resource and return false if store should not return it
		//
		// Store then loads additional resources to satisfy the paging parameters
		Check func(*Role) (bool, error) `json:"-"`

		// Standard helpers for paging and sorting
		store.Sorting
		store.Paging
	}

	RoleMetrics struct {
		Total         uint   `json:"total"`
		Valid         uint   `json:"valid"`
		Deleted       uint   `json:"deleted"`
		Archived      uint   `json:"archived"`
		DailyCreated  []uint `json:"dailyCreated"`
		DailyDeleted  []uint `json:"dailyDeleted"`
		DailyUpdated  []uint `json:"dailyUpdated"`
		DailyArchived []uint `json:"dailyArchived"`
	}
)

// Resource returns a resource ID for this type
func (r *Role) PermissionResource() permissions.Resource {
	return RolePermissionResource.AppendID(r.ID)
}

func (r *Role) DynamicRoles(userID uint64) []uint64 {
	return nil
}

// FindByHandle finds role by it's handle
func (set RoleSet) FindByHandle(handle string) *Role {
	for i := range set {
		if set[i].Handle == handle {
			return set[i]
		}
	}

	return nil
}
