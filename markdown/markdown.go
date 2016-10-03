package markdown

import (
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/patrickrand/swagdown/swagger"
)

var (
	TemplateFilename = filepath.Join(os.Getenv("GOPATH"), "src/github.com/patrickrand/swagdown/markdown/templates/api.md")
	Template         = template.Must(template.ParseFiles(TemplateFilename))
)

func Render(w io.Writer, r io.Reader) error {
	api, err := swagger.Decode(r)
	if err != nil {
		return err
	}

	if err := Template.Execute(w, api); err != nil {
		return err
	}

	return nil
}
