package codegen

import (
	"github.com/cortezaproject/corteza-server/pkg/cli"
	"github.com/davecgh/go-spew/spew"
	"strings"
	"text/template"
)

type (
	definitions struct {
		App string

		Rest    []*restDef
		Actions []*actionsDef
		Events  []*eventsDef
		Types   []*typesDef
	}
)

func Proc() {
	var (
		err error

		def = &definitions{}

		tpls = template.New("").Funcs(map[string]interface{}{
			"camelCase":       camelCase,
			"pubIdent":        pubIdent,
			"toLower":         strings.ToLower,
			"normalizeImport": normalizeImport,
		})
	)

	tpls = template.Must(tpls.ParseGlob("pkg/codegen/assets/*.tpl"))

	if def.Actions, err = procActions(); err != nil {
		cli.HandleError(err)
	} else {
		cli.HandleError(genActions(tpls, def.Actions))
	}

	if def.Events, err = procEvents(); err != nil {
		cli.HandleError(err)
	} else {
		cli.HandleError(genEvents(tpls, def.Events))
	}

	if def.Types, err = procTypes(); err != nil {
		cli.HandleError(err)
	} else {
		cli.HandleError(genTypes(tpls, def.Types))
	}

	if def.Rest, err = procRest(); err != nil {
		cli.HandleError(err)
	} else {
		spew.Dump(def.Rest)
		cli.HandleError(genRest(tpls, def.Rest))
	}
}
