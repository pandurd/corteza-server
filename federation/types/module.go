package types

import (
	"encoding/json"

	composeTypes "github.com/cortezaproject/corteza-server/compose/types"
)

type (
	FederatedModule struct {
		composeTypes.Module    `json:"m"`
		composeTypes.Namespace `json:"ns"`
	}

	FederatedModuleFieldSet composeTypes.ModuleFieldSet
	FederatedModuleSet      []*FederatedModule
)

func (md FederatedModule) MarshalJSON() ([]byte, error) {
	// marshal fields
	fields := []map[string]interface{}{}

	for _, f := range md.Module.Fields {
		field := map[string]interface{}{
			"id": f.ID,
			"k":  f.Kind,
			"n":  f.Name,
			"l":  f.Label,
			"o":  f.Options,
		}

		fields = append(fields, field)
	}

	m := map[string]interface{}{
		"m": map[string]interface{}{
			"id": md.Module.ID,
			"hn": md.Module.Handle,
			"nm": md.Module.Name,
			"mt": md.Module.Meta,
			"f":  fields,
		},
		"ns": map[string]interface{}{
			"id": md.Namespace.ID,
			"nm": md.Namespace.Name,
			"sl": md.Namespace.Slug,
			"en": md.Namespace.Enabled,
			"st": md.Namespace.Meta.Subtitle,
			"ds": md.Namespace.Meta.Description,
		},
	}

	return json.Marshal(m)
}

// // Resource returns a system resource ID for this type
// func (m Module) PermissionResource() permissions.Resource {
// 	return ModulePermissionResource.AppendID(m.ID)
// }

// // FindByHandle finds module by it's handle
// func (set ModuleSet) FindByHandle(handle string) *Module {
// 	for i := range set {
// 		if set[i].Handle == handle {
// 			return set[i]
// 		}
// 	}

// 	return nil
// }
