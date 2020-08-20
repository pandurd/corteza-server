package tests

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//

import (
	"context"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/stretchr/testify/require"
	"testing"
)

func testAllGenerated(t *testing.T, s store.StoreInterfaces) {
{{ range . }}
	// Run generated tests for {{ .Types.Base }}
	t.Run({{ printf "%q" .Types.Base }}, func(t *testing.T) {
		require.New(t).NoError(s.Truncate{{ pubIdent .Types.Plural }}(context.Background()))
		test{{ pubIdent .Types.Base }}(t, s)
	})
{{ end }}
}
