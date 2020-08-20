package bulk

// This file is auto-generated.
//
// Template:    pkg/codegen/assets/store_bulk.gen.go.tpl
// Definitions: {{ .Source }}
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
// Definitions file that controls how this file is generated:

import (
	"context"
{{- range $import := .Import }}
    {{ normalizeImport $import }}
{{- end }}
)


{{ if not $.Search.Disable }}
{{ $struct := printf "search%s" ( pubIdent $.Types.Plural ) }}
type (
	{{ $struct }} struct {
		done chan struct{}
		err  error
		set  {{ $.Types.GoSetType }}
		filter {{ .Types.GoFilterType  }}
		rfilter {{ .Types.GoFilterType  }}
	}
)

// {{ pubIdent $struct }} returns all matching rows
//
// This function calls convert{{ pubIdent $.Types.Singular }}Filter with the given
func {{ pubIdent $struct }}(filter {{ .Types.GoFilterType  }}) *{{ $struct }} {
	return &{{ $struct }}{
		filter: filter,
		done: make(chan struct{}, 1),
	}
}

// Do executes {{ $struct }} job
func (j *{{ $struct }}) Do(ctx context.Context, s storeInterface) error {
	j.set, j.rfilter, j.err = s.{{ pubIdent $struct }}(ctx, j.filter)
	j.done <- struct{}{}
	return j.err
}

// Collect results of ( $.Types.Singular ) Search
//
// Note: this function blocks until job is done
func (j {{ $struct }}) Collect() ({{ $.Types.GoSetType }}, {{ .Types.GoFilterType  }}, error) {
	<-j.done // block until job gets done
	return j.set, j.rfilter, j.err
}

{{ template "Push" $struct }}
{{ template "GetError" $struct }}
{{ end }}

{{/* ************************************************************************************************************** */}}
{{/* ************************************************************************************************************** */}}

{{- range $lookup := $.Lookups }}

{{ $struct := printf "lookup%sBy%s" ( pubIdent $.Types.Singular ) ( pubIdent $lookup.Suffix ) }}
// {{ $struct }}


type (
	{{ $struct }} struct {
		done chan struct{}
		err  error
		res  *{{ $.Types.GoType }}
	{{- range $lookup.Fields }}
		arg{{ pubIdent . }} {{ (. | $.Fields.Find).Type  }}
	{{- end }}
	}
)


// {{ pubIdent $struct }} {{ comment $lookup.Description true -}}
func {{ pubIdent $struct }}({{- range $lookup.Fields }}{{ cc2underscore . }} {{ (. | $.Fields.Find).Type  }}, {{- end }}) *{{ $struct }} {
	return &{{ $struct }}{
	{{- range $lookup.Fields }}
		arg{{ pubIdent . }}: {{ cc2underscore . }},
	{{- end }}
		done: make(chan struct{}, 1),
	}
}

// Do executes {{ $struct }} job
func (j *{{ $struct }}) Do(ctx context.Context, s storeInterface) error {
	j.res, j.err = s.{{ pubIdent $struct }}(
		ctx,
		{{- range $lookup.Fields }}
		j.arg{{ pubIdent . }},
		{{ end }}
	)

	j.done <- struct{}{}
	return j.err
}

// Collect results of ( $.Types.Singular ) lookup
//
// Note: this function blocks until job is done
func (j *{{ $struct }}) Collect() (*{{ $.Types.GoType }}, error) {
	<-j.done // block until job gets done
	return j.res, j.err
}

{{ template "Push" $struct }}
{{ template "GetError" $struct }}
{{ end }}

{{/* ************************************************************************************************************** */}}
{{/* ************************************************************************************************************** */}}


{{ template "Job" ( dict "prefix" "create" "Types" .Types ) }}
{{ template "Job" ( dict "prefix" "update" "Types" .Types ) }}
{{ template "Job" ( dict "prefix" "remove" "Types" .Types ) }}


{{/* ************************************************************************************************************** */}}
{{/* ************************************************************************************************************** */}}


{{- define "Job" -}}
{{ $struct := printf "%s%s" .prefix ( pubIdent $.Types.Singular ) }}
type (
	{{ $struct }} struct {
	    done chan struct{}
		err  error
		res  *{{ .Types.GoType }}
    }
)

// {{ pubIdent $struct }} creates a new {{ .prefix }} job for {{ pubIdent .Types.Singular }} that can be pushed to store's transaction handler
func {{ pubIdent $struct }}(res *{{ .Types.GoType }}) *{{ $struct }} {
    return &{{ $struct }}{res: res, done: make(chan struct{}, 1)}
}

// Do Executes {{ $struct }} job
//
func (j *{{ $struct }}) Do(ctx context.Context, s storeInterface) error {
	j.err = s.{{ pubIdent $struct }}(ctx, j.res)
	j.done <- struct{}{}
	return j.err
}

{{ template "Push" $struct }}
{{ template "GetError" $struct }}
{{- end -}}

{{- define "GetError" -}}
// GetError returns job error (if any)
//
// Note: this function blocks until job is done
func (j {{ . }}) GetError() error {
	<-j.done // block until job gets done
	return j.err
}
{{- end -}}

{{- define "Push" -}}
// Push accepts transaction channel and returns struct
//
// A small helper that allows more fluid development
func (j *{{ . }}) Push(tx chan <- Job) *{{ . }} {
	tx <- j
	return j
}
{{- end -}}
