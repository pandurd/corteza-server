package codegen

import (
	"bytes"
	"fmt"
	"github.com/cortezaproject/corteza-server/pkg/cli"
	"go/format"
	"io"
	"os"
	"regexp"
	"strings"
	"text/template"
)

func goTemplate(dst string, tpl *template.Template, payload interface{}) error {
	var output io.WriteCloser
	buf := bytes.Buffer{}

	if err := tpl.Execute(&buf, payload); err != nil {
		return err
	}

	fmtsrc, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "fmt warn: %v", err)
		fmtsrc = buf.Bytes()
	}

	if dst == "" || dst == "-" {
		output = os.Stdout
	} else {
		if output, err = os.Create(dst); err != nil {
			return err
		}

		defer output.Close()
	}

	if _, err := output.Write(fmtsrc); err != nil {
		return err
	}

	return nil
}

func WritePlainTo(tpl *template.Template, payload interface{}, tplName, dst string) {
	var output io.WriteCloser
	buf := bytes.Buffer{}

	if err := tpl.ExecuteTemplate(&buf, tplName, payload); err != nil {
		cli.HandleError(err)
	} else {
		if dst == "" || dst == "-" {
			output = os.Stdout
		} else {
			// cli.HandleError(os.Remove(dst))
			if output, err = os.Create(dst); err != nil {
				cli.HandleError(err)
			}

			defer output.Close()
		}

		if _, err := output.Write(buf.Bytes()); err != nil {
			cli.HandleError(err)
		}
	}
}

func camelCase(pp ...string) (out string) {
	for i, p := range pp {
		if i > 0 && len(p) > 1 {
			p = strings.ToUpper(p[:1]) + p[1:]
		}

		out = out + p
	}

	return out
}

// convets to underscore
func cc2underscore(cc string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	u := matchFirstCap.ReplaceAllString(cc, "${1}_${2}")
	u = matchAllCap.ReplaceAllString(u, "${1}_${2}")
	return strings.ToLower(u)
}
