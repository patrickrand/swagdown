package markdown

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/patrickrand/swagdown/swagger"
)

var (
	SwagdownTemplates = filepath.Join(os.Getenv("GOPATH"), "src/github.com/patrickrand/swagdown/markdown/templates/*")
	funcMap           = template.FuncMap{"FilterParameters": swagger.FilterParameters}
	tmpl              = template.Must(template.New("outer.md").Funcs(funcMap).ParseGlob(SwagdownTemplates))
)

func RenderFromJSON(w io.Writer, r io.Reader) error {
	return render(w, r, swagger.DecodeJSON)
}

func RenderFromYAML(w io.Writer, r io.Reader) error {
	return render(w, r, swagger.DecodeYAML)
}

func render(w io.Writer, r io.Reader, decode func(data []byte) (*swagger.API, error)) error {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}

	api, err := decode(buf.Bytes())
	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, api); err != nil {
		return err
	}

	return nil
}
