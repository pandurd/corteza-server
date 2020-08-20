package store

// This file is auto-generated.
//
// Template:    pkg/codegen/assets/store_base.gen.go.tpl
// Definitions: {{ .Source }}
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//  - {{ .Source }}

import (
	"context"
{{- range .Import }}
	{{ normalizeImport . }}
{{- end }}
)

type (
	{{- $Types := .Types }}
	{{- $Fields := .Fields }}

	{{ pubIdent .Types.Plural }} interface {
	{{- if not .Search.Disable }}
		Search{{ pubIdent $Types.Plural }}(ctx context.Context, f {{ $Types.GoFilterType }}) ({{ $Types.GoSetType }}, {{ $Types.GoFilterType }}, error)
	{{- end }}
	{{- range .Lookups }}
		Lookup{{ pubIdent $Types.Singular }}By{{ pubIdent .Suffix }}(ctx context.Context{{- range $field := .Fields }}, {{ cc2underscore $field }} {{ ($field | $Fields.Find).Type  }}{{- end }}) (*{{ $Types.GoType }}, error)
	{{- end }}
		Create{{ pubIdent $Types.Singular }}(ctx context.Context, rr ... *{{ $Types.GoType }}) error
		Update{{ pubIdent $Types.Singular }}(ctx context.Context, rr ... *{{ $Types.GoType }}) error
		Partial{{ pubIdent $Types.Singular }}Update(ctx context.Context, onlyColumns []string, rr ... *{{ $Types.GoType }}) error
		Remove{{ pubIdent $Types.Singular }}(ctx context.Context, rr ... *{{ $Types.GoType }}) error
		Remove{{ pubIdent $Types.Singular }}By{{ template "primaryKeySuffix" $Fields }}(ctx context.Context {{ template "primaryKeyArgsIn" $Fields }}) error

		Truncate{{ pubIdent $Types.Plural }}(ctx context.Context) error

		// Extra functions
	{{- range .Extra }}
		{{ .Name }}(ctx context.Context{{ range .Args }}, {{ .Name }} {{ .Type }}{{ end }}) ({{ join ", " .Return }})
	{{- end }}

	}
)

{{- if not .Search.Disable }}
// Search{{ pubIdent $.Types.Plural }} returns all matching {{ $.Types.Plural }} from store
func Search{{ pubIdent $Types.Plural }}(ctx context.Context, s {{ pubIdent $Types.Plural }}, f {{ $Types.GoFilterType }}) ({{ $Types.GoSetType }}, {{ $Types.GoFilterType }}, error) {
	return s.Search{{ pubIdent $Types.Plural }}(ctx, f)
}
{{- end }}

{{- range .Lookups }}
// Lookup{{ pubIdent $.Types.Singular }}By{{ pubIdent .Suffix }} {{ comment .Description true -}}
func Lookup{{ pubIdent $Types.Singular }}By{{ pubIdent .Suffix }}(ctx context.Context, s {{ pubIdent $Types.Plural }}{{- range $field := .Fields }}, {{ cc2underscore $field }} {{ ($field | $Fields.Find).Type  }}{{- end }}) (*{{ $Types.GoType }}, error) {
    return s.Lookup{{ pubIdent $Types.Singular }}By{{ pubIdent .Suffix }}(ctx{{- range $field := .Fields }}, {{ cc2underscore $field }}{{- end }})
}
{{- end }}

// Create{{ pubIdent $.Types.Singular }} creates one or more {{ $.Types.Plural }} in store
func Create{{ pubIdent $Types.Singular }}(ctx context.Context, s {{ pubIdent $Types.Plural }}, rr ... *{{ $Types.GoType }}) error {
	return s.Create{{ pubIdent $Types.Singular }}(ctx, rr... )
}

// Update{{ pubIdent $.Types.Singular }} updates one or more (existing) {{ $.Types.Plural }} in store
func Update{{ pubIdent $Types.Singular }}(ctx context.Context, s {{ pubIdent $Types.Plural }}, rr ... *{{ $Types.GoType }}) error {
	return s.Update{{ pubIdent $Types.Singular }}(ctx, rr... )
}

// Partial{{ pubIdent $.Types.Singular }}Update updates one or more existing {{ $.Types.Plural }} in store
func Partial{{ pubIdent $Types.Singular }}Update(ctx context.Context, s {{ pubIdent $Types.Plural }}, onlyColumns []string, rr ... *{{ $Types.GoType }}) error {
	return s.Partial{{ pubIdent $Types.Singular }}Update(ctx, onlyColumns, rr...)
}

// Remove{{ pubIdent $.Types.Singular }} removes one or more {{ $.Types.Plural }} from store
func Remove{{ pubIdent $Types.Singular }}(ctx context.Context, s {{ pubIdent $Types.Plural }}, rr ... *{{ $Types.GoType }}) error {
	return s.Remove{{ pubIdent $Types.Singular }}(ctx, rr...)
}

// Remove{{ pubIdent $.Types.Singular }}By{{ template "primaryKeySuffix" $.Fields }} removes {{ $.Types.Singular }} from store
func Remove{{ pubIdent $Types.Singular }}By{{ template "primaryKeySuffix" $Fields }}(ctx context.Context, s {{ pubIdent $Types.Plural }} {{ template "primaryKeyArgsIn" $Fields }}) error {
	return s.Remove{{ pubIdent $Types.Singular }}By{{ template "primaryKeySuffix" $Fields }}(ctx{{ template "primaryKeyArgsOut" $Fields }})
}

// Truncate{{ pubIdent $.Types.Plural }} removes all {{ $.Types.Plural }} from store
func Truncate{{ pubIdent $Types.Plural }}(ctx context.Context, s {{ pubIdent $Types.Plural }}) error {
	return s.Truncate{{ pubIdent $Types.Plural }}(ctx)
}

{{ range .Extra }}
func {{ .Name }}(ctx context.Context, s {{ pubIdent $Types.Plural }}{{ range .Args }}, {{ .Name }} {{ .Type }}{{ end }}) ({{ join ", " .Return }}) {
	return s.{{ .Name }}(ctx{{ range .Args }}, {{ .Name }}{{ end }})
}
{{ end }}
