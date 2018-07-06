package crm

/*
	Hello! This file is auto-generated from `docs/src/spec.json`.

	For development:
	In order to update the generated files, edit this file under the location,
	add your struct fields, imports, API definitions and whatever you want, and:

	1. run [spec](https://github.com/titpetric/spec) in the same folder,
	2. run `./_gen.php` in this folder.

	You may edit `types.go`, `types.util.go` or `types_test.go` to
	implement your API calls, helper functions and tests. The file `types.go`
	is only generated the first time, and will not be overwritten if it exists.
*/

type (
	// Types
	Types struct {
		changed []string
	}
)

/* Constructors */
func (Types) New() *Types {
	return &Types{}
}

/* Getters/setters */
