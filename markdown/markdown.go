package markdown

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/patrickrand/swagdown/swagger"
)

var SwagdownTemplates = filepath.Join(os.Getenv("GOPATH"), "src/github.com/patrickrand/swagdown/markdown/templates/*")

// just kidding around... will refactor later
var tmpl = template.Must(template.New("outer.md").Funcs(template.FuncMap{
	"ToUpper": strings.ToUpper,
	"Capitalize": func(s string) string {
		if len(s) == 0 {
			return s
		}
		return strings.ToUpper(string(s[0])) + s[1:]
	},
	"FilterParameters": func(params []swagger.Parameter, t string) []swagger.Parameter {
		var results []swagger.Parameter
		for _, param := range params {
			if strings.ToLower(param.In) == t {
				results = append(results, param)
			}
		}
		return results
	},
}).ParseGlob(SwagdownTemplates))

func Render(w io.Writer, r io.Reader) error {
	api, err := swagger.Decode(r)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, api); err != nil {
		return err
	}

	return nil
}
