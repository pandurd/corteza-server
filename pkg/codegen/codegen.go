package codegen

import (
	"github.com/cortezaproject/corteza-server/pkg/cli"
	"text/template"
)

type (
	definitions struct {
		App string

		RestAPI RestAPI
		Actions []*actionsDef
		Events  []*eventsDef
		Types   []*typesDef
	}

	RestAPI struct{}
)

func Proc() {
	var (
		err error

		def = &definitions{}

		tpls = template.New("").Funcs(map[string]interface{}{
			"camelCase": camelCase,
		})
	)

	tpls = template.Must(tpls.ParseGlob("pkg/codegen/assets/*.tpl"))

	if def.Actions, err = procActions(); err != nil {
		cli.HandleError(err)
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
}
